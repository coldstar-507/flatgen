// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ChatScrollRequestT struct {
	ChatId string `json:"chat_id"`
	Before bool `json:"before"`
}

func (t *ChatScrollRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	chatIdOffset := builder.CreateString(t.ChatId)
	ChatScrollRequestStart(builder)
	ChatScrollRequestAddChatId(builder, chatIdOffset)
	ChatScrollRequestAddBefore(builder, t.Before)
	return ChatScrollRequestEnd(builder)
}

func (rcv *ChatScrollRequest) UnPackTo(t *ChatScrollRequestT) {
	t.ChatId = string(rcv.ChatId())
	t.Before = rcv.Before()
}

func (rcv *ChatScrollRequest) UnPack() *ChatScrollRequestT {
	if rcv == nil { return nil }
	t := &ChatScrollRequestT{}
	rcv.UnPackTo(t)
	return t
}

type ChatScrollRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsChatScrollRequest(buf []byte, offset flatbuffers.UOffsetT) *ChatScrollRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ChatScrollRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsChatScrollRequest(buf []byte, offset flatbuffers.UOffsetT) *ChatScrollRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ChatScrollRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ChatScrollRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ChatScrollRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ChatScrollRequest) ChatId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChatScrollRequest) Before() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ChatScrollRequest) MutateBefore(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func ChatScrollRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ChatScrollRequestAddChatId(builder *flatbuffers.Builder, chatId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(chatId), 0)
}
func ChatScrollRequestAddBefore(builder *flatbuffers.Builder, before bool) {
	builder.PrependBoolSlot(1, before, false)
}
func ChatScrollRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
