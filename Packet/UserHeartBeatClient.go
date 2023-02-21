// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Packet

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserHeartBeatClient struct {
	_tab flatbuffers.Table
}

func GetRootAsUserHeartBeatClient(buf []byte, offset flatbuffers.UOffsetT) *UserHeartBeatClient {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UserHeartBeatClient{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUserHeartBeatClient(buf []byte, offset flatbuffers.UOffsetT) *UserHeartBeatClient {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UserHeartBeatClient{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UserHeartBeatClient) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UserHeartBeatClient) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UserHeartBeatClient) LastSeqIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UserHeartBeatClient) MutateLastSeqIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func UserHeartBeatClientStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UserHeartBeatClientAddLastSeqIndex(builder *flatbuffers.Builder, lastSeqIndex int32) {
	builder.PrependInt32Slot(0, lastSeqIndex, 0)
}
func UserHeartBeatClientEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}