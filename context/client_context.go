package context

import (
	"encoding/binary"
	"fmt"
	"golang_server/Packet"
	. "golang_server/util"
	"net"
	"runtime"
	"sync"
	"time"
)

const client_send_channel_size int = 10000

type ClientContext struct {
	User_index       int
	Client_state     ClientStateEnum
	conn             *net.TCPConn
	recv_timeout_sec time.Duration

	wg sync.WaitGroup

	last_recv_seq_index int
	next_send_seq_index int

	send_channel                    chan *[]byte
	Last_confirm_send_seq_index     int
	Confirm_waiting_send_packet_map map[int]*[]byte

	server_context *ServerContext

	//closed                  bool
	//closed_recv_packet_list [][]byte
	//close_mutex *sync.Mutex
}

func (client_context *ClientContext) Init(conn *net.TCPConn, server_context *ServerContext) {
	client_context.conn = conn
	//todo 데이터 다운로드, 매칭 대기 등 정상적으로 타임아웃이 길어지는 경우를 고려해야한다
	client_context.recv_timeout_sec = 1000

	client_context.last_recv_seq_index = 0
	client_context.next_send_seq_index = 1
	client_context.send_channel = make(chan *[]byte, client_send_channel_size)
	client_context.Last_confirm_send_seq_index = 0
	client_context.Confirm_waiting_send_packet_map = make(map[int]*[]byte)
	client_context.server_context = server_context

	//client_context.closed = false
	//client_context.closed_recv_packet_list = [][]byte{}
	//client_context.close_mutex = &sync.Mutex{}

	conn.SetNoDelay(true)
	conn.SetKeepAlive(true)
	//conn.SetLinger(1)
	conn.SetReadBuffer(1000 * 100)
	conn.SetWriteBuffer(1000 * 100)

	go client_context.recvExec()
	go client_context.sendExec()
}

func (client_context *ClientContext) Close() {
	client_context.conn.Close()
}

func (client_context *ClientContext) AddWg() {
	client_context.wg.Add(1)
}

func (client_context *ClientContext) DoneWg() {
	client_context.wg.Done()
}

func (client_context *ClientContext) recvExec() {
	defer func() {
		r := recover()
		if r != nil { //io.EOF {
			fmt.Println("recvExec recover : ", r)
			client_context.Close()
		}
	}()
	const prefixSizeLength = 4
	const prefixSeqLength = 4

	conn := (*client_context.conn)
	//clientReader := bufio.NewReader(client_context.conn)

	RecvFunc := func(need_len int) ([]byte, bool) {
		buffer := make([]byte, 0)
		for {
			temp_buffer := make([]byte, need_len-len(buffer))

			conn.SetReadDeadline(time.Now().Add(client_context.recv_timeout_sec * time.Second))
			n, recv_err := conn.Read(temp_buffer) // 클라이언트에서 받은 데이터를 읽음
			//n, recv_err := clientReader.Read(temp_header_buffer)

			if recv_err != nil {
				fmt.Println("Recv Error: ", recv_err)
				return buffer, false
			}
			if n == 0 {
				return buffer, false
			}

			if n == need_len {
				buffer = temp_buffer
				break
			} else {
				buffer = append(buffer, temp_buffer[0:n]...)
				if need_len <= len(buffer) {
					break
				}
			}
		}
		return buffer, true
	}

	for {
		client_context.wg.Wait()
		client_context.AddWg()
		packet_size_buffer, ok := RecvFunc(prefixSizeLength)
		if !ok {
			client_context.Close()
			return
		}
		packet_size := int(binary.LittleEndian.Uint32(packet_size_buffer))
		///////////////////////////////////////////////////////////////////////////
		recv_seq_buffer, ok := RecvFunc(prefixSeqLength)
		if !ok {
			client_context.Close()
			return
		}

		recv_seq_index := int(binary.LittleEndian.Uint32(recv_seq_buffer))
		if client_context.last_recv_seq_index+1 != recv_seq_index {
			fmt.Println("client_context.Last_recv_seq_index + 1 != recv_seq_index",
				client_context.last_recv_seq_index+1, recv_seq_index)
			client_context.Close()
			return
		}
		client_context.last_recv_seq_index = recv_seq_index
		///////////////////////////////////////////////////////////////////////////
		last_recv_seq_buffer, ok := RecvFunc(prefixSeqLength)
		if !ok {
			client_context.Close()
			return
		}

		last_seq_index := int(binary.LittleEndian.Uint32(last_recv_seq_buffer)) // 클라가 마지막 처리한 서버 패킷
		if last_seq_index < client_context.Last_confirm_send_seq_index {
			client_context.Close()
			return
		}
		client_context.Last_confirm_send_seq_index = last_seq_index

		///////////////////////////////////////////////////////////////////////////
		packet_buffer, ok := RecvFunc(packet_size)
		if !ok {
			client_context.Close()
			return
		}

		packet_base := Packet.GetRootAsPacketBase(packet_buffer, 0)
		packet_code := packet_base.BodyType()

		if !Contains([]Packet.PacketCode{
			Packet.PacketCodeUserHeartBeatClient,
		}, packet_code) {
			fmt.Println("PacketCode: ", packet_code.String(), " user_index: ", client_context.User_index)
		}
		client_context.server_context.Recv_channel <- &PacketRecvChannelStruct{client_context, packet_code, &packet_buffer}
		runtime.Gosched()
	}
}

func (client_context *ClientContext) sendExec() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화

		fmt.Println("sendExec recover : ", r)
		client_context.Close()
	}()

	const PacketSendOnceMaxCount = 300

	//clientWriter := bufio.NewWriterSize(client_context.conn, PacketSendOnceMaxCount*2000)
	SendPacketMakeFunc := func(packet_buffer *[]byte) []byte {
		_packet := packet_buffer
		pre_buff := []byte{}
		pre_buff = binary.LittleEndian.AppendUint32(pre_buff, uint32(len(*_packet)))
		pre_buff = binary.LittleEndian.AppendUint32(pre_buff, uint32(client_context.next_send_seq_index))
		pre_buff = binary.LittleEndian.AppendUint32(pre_buff, uint32(client_context.last_recv_seq_index))
		result := append(pre_buff, *_packet...)

		client_context.Confirm_waiting_send_packet_map[client_context.next_send_seq_index] = &result
		client_context.next_send_seq_index++
		return result
	}

	for {
		packet_buffer := <-(client_context.send_channel)

		for i := client_context.Last_confirm_send_seq_index; 0 <= i; i-- {
			if _, ok := client_context.Confirm_waiting_send_packet_map[i]; ok {
				delete(client_context.Confirm_waiting_send_packet_map, i)
			} else {
				break
			}
		}

		packet_all := SendPacketMakeFunc(packet_buffer)
		count := Min(PacketSendOnceMaxCount, len(client_context.send_channel))
		if count == PacketSendOnceMaxCount {
			fmt.Println("count == PacketSendOnceMaxCount", client_context.User_index)
		}
		for i := 0; i < count; i++ {
			add_packet_struct := <-(client_context.send_channel)

			result := SendPacketMakeFunc(add_packet_struct)

			packet_all = append(packet_all, result...)
		}
		_, err := (*client_context.conn).Write(packet_all)
		//_, err := clientWriter.Write(packet_all)

		if err != nil {
			fmt.Println("if err != nil {", err)
			client_context.Close()
			return
		}
		//clientWriter.Flush()
		runtime.Gosched()
	}
}

func (client_context *ClientContext) SendPacket(packet_buffer []byte) {
	if client_send_channel_size <= len(client_context.send_channel) {
		fmt.Println("send_channel_count<= len(client_context.send_channel)")
		//client_context.Client_state = ClientStateErrorKickOut
		client_context.Close()
		return
	}

	client_context.send_channel <- &packet_buffer
}
