package context

import (
	"fmt"
	"golang_server/Packet"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ClientProcedureFunc func(client_context *ClientContext, unionTable *flatbuffers.Table)

type RecvChannelInterface interface {
	Exec(server_context *ServerContext)
}

type PacketRecvChannelStruct struct {
	Client_context *ClientContext
	Packet_code    Packet.PacketCode
	Packet_buffer  *[]byte
}

func (data *PacketRecvChannelStruct) Exec(server_context *ServerContext) {
	packet_buffer := *data.Packet_buffer
	packet_code := data.Packet_code
	client_context := data.Client_context

	packet_base := Packet.GetRootAsPacketBase(packet_buffer, 0)

	unionTable := new(flatbuffers.Table)
	result := packet_base.Body(unionTable)
	if !result {
		fmt.Println("if !result", len(packet_buffer), packet_buffer)
		return
	}

	if function, ok := server_context.Client_proc_map[packet_code]; ok {
		function(client_context, unionTable)
	} else {
		fmt.Println("no procedure" + packet_code.String())
	}

	client_context.DoneWg()
}

type FuncRecvChannelStruct struct {
	Client_context *ClientContext
	Function       func(Client_context *ClientContext, args ...interface{})
	Param          []any
}

func (data *FuncRecvChannelStruct) Exec(server_context *ServerContext) {
	data.Function(data.Client_context, data.Param)

	if data.Client_context != nil {
		data.Client_context.DoneWg()
	}
}

type ServerContext struct {
	port int

	//Recv_channel    chan PacketRecvChannelInterface
	Recv_channel chan RecvChannelInterface

	Client_proc_map map[Packet.PacketCode]ClientProcedureFunc

	sigs chan os.Signal
}

func (server_context *ServerContext) Init(port int) {
	//server_context.Recv_channel = make(chan PacketRecvChannelInterface, 10000)
	server_context.Recv_channel = make(chan RecvChannelInterface, 100000)

	server_context.Client_proc_map = make(map[Packet.PacketCode]ClientProcedureFunc)

	server_context.port = port

	go server_context.exec()

	server_context.sigs = make(chan os.Signal, 1)
	signal.Notify(server_context.sigs, syscall.Signal(0xa))
	go func() {
		sig := <-server_context.sigs
		fmt.Println(sig)
		//server_context.Recv_channel <- &InnerRecvChannelStruct{nil, ProcessInnerCodeDataReload, nil}
	}()
}

func (server_context *ServerContext) Listen() {

	listener, listen_error := net.Listen("tcp", ":"+strconv.Itoa(server_context.port))
	if listen_error != nil {
		fmt.Println(listen_error)
		return
	}
	fmt.Println("Listen")

	defer listener.Close()

	for {
		conn, accept_error := listener.Accept()
		fmt.Println("Connect")
		if accept_error != nil {
			fmt.Println(accept_error)
			continue
		}

		server_context.Accepted(conn)
	}
}

func (server_context *ServerContext) Accepted(conn net.Conn) {
	tcp_conn := conn.(*net.TCPConn)
	client_context := new(ClientContext)
	client_context.Init(tcp_conn, server_context)
	client_context.AddWg()

	Connected := func(_client_context *ClientContext, args ...interface{}) {

	}

	server_context.Recv_channel <- &FuncRecvChannelStruct{client_context, Connected, nil}
}

func (server_context *ServerContext) exec() {
	for {
		process_data := <-(server_context.Recv_channel)
		if 99900 < len(server_context.Recv_channel) {
			fmt.Println("99900 < len(server_context.Recv_channel)")
		}
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			process_data.Exec(server_context)

		}()
	}
}

var once sync.Once

var server_context *ServerContext

func GetServerContext() *ServerContext {
	once.Do(func() {
		server_context = &ServerContext{}
	})
	return server_context
}
