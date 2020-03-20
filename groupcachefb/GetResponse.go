// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package groupcachefb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GetResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsGetResponse(buf []byte, offset flatbuffers.UOffsetT) *GetResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GetResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetResponse) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *GetResponse) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GetResponse) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *GetResponse) MinuteQps() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetResponse) MutateMinuteQps(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func GetResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GetResponseAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func GetResponseStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func GetResponseAddMinuteQps(builder *flatbuffers.Builder, minuteQps float64) {
	builder.PrependFloat64Slot(1, minuteQps, 0.0)
}
func GetResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
