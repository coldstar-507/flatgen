// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MessageTokenT struct {
	DeviceId string `json:"device_id"`
	Token string `json:"token"`
}

func (t *MessageTokenT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	deviceIdOffset := builder.CreateString(t.DeviceId)
	tokenOffset := builder.CreateString(t.Token)
	MessageTokenStart(builder)
	MessageTokenAddDeviceId(builder, deviceIdOffset)
	MessageTokenAddToken(builder, tokenOffset)
	return MessageTokenEnd(builder)
}

func (rcv *MessageToken) UnPackTo(t *MessageTokenT) {
	t.DeviceId = string(rcv.DeviceId())
	t.Token = string(rcv.Token())
}

func (rcv *MessageToken) UnPack() *MessageTokenT {
	if rcv == nil { return nil }
	t := &MessageTokenT{}
	rcv.UnPackTo(t)
	return t
}

type MessageToken struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageToken(buf []byte, offset flatbuffers.UOffsetT) *MessageToken {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageToken{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMessageToken(buf []byte, offset flatbuffers.UOffsetT) *MessageToken {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MessageToken{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MessageToken) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageToken) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageToken) DeviceId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MessageToken) Token() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MessageTokenStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageTokenAddDeviceId(builder *flatbuffers.Builder, deviceId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(deviceId), 0)
}
func MessageTokenAddToken(builder *flatbuffers.Builder, token flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(token), 0)
}
func MessageTokenEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
