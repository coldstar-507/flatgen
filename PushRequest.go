// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PushRequestT struct {
	PushId *PushIdT `json:"push_id"`
	Payload []byte `json:"payload"`
}

func (t *PushRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	pushIdOffset := t.PushId.Pack(builder)
	payloadOffset := flatbuffers.UOffsetT(0)
	if t.Payload != nil {
		payloadOffset = builder.CreateByteString(t.Payload)
	}
	PushRequestStart(builder)
	PushRequestAddPushId(builder, pushIdOffset)
	PushRequestAddPayload(builder, payloadOffset)
	return PushRequestEnd(builder)
}

func (rcv *PushRequest) UnPackTo(t *PushRequestT) {
	t.PushId = rcv.PushId(nil).UnPack()
	t.Payload = rcv.PayloadBytes()
}

func (rcv *PushRequest) UnPack() *PushRequestT {
	if rcv == nil {
		return nil
	}
	t := &PushRequestT{}
	rcv.UnPackTo(t)
	return t
}

type PushRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsPushRequest(buf []byte, offset flatbuffers.UOffsetT) *PushRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PushRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishPushRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPushRequest(buf []byte, offset flatbuffers.UOffsetT) *PushRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PushRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPushRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PushRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PushRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PushRequest) PushId(obj *PushId) *PushId {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PushId)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PushRequest) Payload(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *PushRequest) PayloadLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PushRequest) PayloadBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PushRequest) MutatePayload(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func PushRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PushRequestAddPushId(builder *flatbuffers.Builder, pushId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pushId), 0)
}
func PushRequestAddPayload(builder *flatbuffers.Builder, payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(payload), 0)
}
func PushRequestStartPayloadVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PushRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
