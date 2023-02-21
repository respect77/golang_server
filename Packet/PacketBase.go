// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Packet

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PacketBase struct {
	_tab flatbuffers.Table
}

func GetRootAsPacketBase(buf []byte, offset flatbuffers.UOffsetT) *PacketBase {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PacketBase{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPacketBase(buf []byte, offset flatbuffers.UOffsetT) *PacketBase {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PacketBase{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PacketBase) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PacketBase) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PacketBase) BodyType() PacketCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return PacketCode(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *PacketBase) MutateBodyType(n PacketCode) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *PacketBase) Body(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func PacketBaseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PacketBaseAddBodyType(builder *flatbuffers.Builder, bodyType PacketCode) {
	builder.PrependByteSlot(0, byte(bodyType), 0)
}
func PacketBaseAddBody(builder *flatbuffers.Builder, body flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(body), 0)
}
func PacketBaseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
