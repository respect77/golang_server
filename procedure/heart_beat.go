package procedure

import (
	"golang_server/Packet"
	. "golang_server/context"
	. "golang_server/util"

	flatbuffers "github.com/google/flatbuffers/go"
)

func HeartBeatFunc(client_context *ClientContext, unionTable *flatbuffers.Table) {
	packet := new(Packet.UserHeartBeatClient)
	packet.Init(unionTable.Bytes, unionTable.Pos)

	builder := flatbuffers.NewBuilder(0)
	Packet.UserHeartBeatServerStart(builder)
	bodyOffset := Packet.UserHeartBeatServerEnd(builder)

	packet_buffer := MakePacketBuffer(builder, Packet.PacketCodeUserHeartBeatServer, bodyOffset)

	client_context.SendPacket(packet_buffer)
}
