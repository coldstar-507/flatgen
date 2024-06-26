// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MediaMetadataT struct {
	TimeId string `json:"time_id"`
	OwnerId string `json:"owner_id"`
	Timestamp int64 `json:"timestamp"`
	Mime string `json:"mime"`
	IsReversed bool `json:"is_reversed"`
	IsEncrypted bool `json:"is_encrypted"`
	IsPaidToView bool `json:"is_paid_to_view"`
	IsPaidToOwn bool `json:"is_paid_to_own"`
	IsLocked bool `json:"is_locked"`
	IsSaved bool `json:"is_saved"`
	Temp string `json:"temp"`
}

func (t *MediaMetadataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	timeIdOffset := builder.CreateString(t.TimeId)
	ownerIdOffset := builder.CreateString(t.OwnerId)
	mimeOffset := builder.CreateString(t.Mime)
	tempOffset := builder.CreateString(t.Temp)
	MediaMetadataStart(builder)
	MediaMetadataAddTimeId(builder, timeIdOffset)
	MediaMetadataAddOwnerId(builder, ownerIdOffset)
	MediaMetadataAddTimestamp(builder, t.Timestamp)
	MediaMetadataAddMime(builder, mimeOffset)
	MediaMetadataAddIsReversed(builder, t.IsReversed)
	MediaMetadataAddIsEncrypted(builder, t.IsEncrypted)
	MediaMetadataAddIsPaidToView(builder, t.IsPaidToView)
	MediaMetadataAddIsPaidToOwn(builder, t.IsPaidToOwn)
	MediaMetadataAddIsLocked(builder, t.IsLocked)
	MediaMetadataAddIsSaved(builder, t.IsSaved)
	MediaMetadataAddTemp(builder, tempOffset)
	return MediaMetadataEnd(builder)
}

func (rcv *MediaMetadata) UnPackTo(t *MediaMetadataT) {
	t.TimeId = string(rcv.TimeId())
	t.OwnerId = string(rcv.OwnerId())
	t.Timestamp = rcv.Timestamp()
	t.Mime = string(rcv.Mime())
	t.IsReversed = rcv.IsReversed()
	t.IsEncrypted = rcv.IsEncrypted()
	t.IsPaidToView = rcv.IsPaidToView()
	t.IsPaidToOwn = rcv.IsPaidToOwn()
	t.IsLocked = rcv.IsLocked()
	t.IsSaved = rcv.IsSaved()
	t.Temp = string(rcv.Temp())
}

func (rcv *MediaMetadata) UnPack() *MediaMetadataT {
	if rcv == nil { return nil }
	t := &MediaMetadataT{}
	rcv.UnPackTo(t)
	return t
}

type MediaMetadata struct {
	_tab flatbuffers.Table
}

func GetRootAsMediaMetadata(buf []byte, offset flatbuffers.UOffsetT) *MediaMetadata {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MediaMetadata{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMediaMetadata(buf []byte, offset flatbuffers.UOffsetT) *MediaMetadata {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MediaMetadata{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MediaMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MediaMetadata) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MediaMetadata) TimeId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MediaMetadata) OwnerId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MediaMetadata) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MediaMetadata) MutateTimestamp(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *MediaMetadata) Mime() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MediaMetadata) IsReversed() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsReversed(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *MediaMetadata) IsEncrypted() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsEncrypted(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func (rcv *MediaMetadata) IsPaidToView() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsPaidToView(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func (rcv *MediaMetadata) IsPaidToOwn() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsPaidToOwn(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *MediaMetadata) IsLocked() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsLocked(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func (rcv *MediaMetadata) IsSaved() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MediaMetadata) MutateIsSaved(n bool) bool {
	return rcv._tab.MutateBoolSlot(22, n)
}

func (rcv *MediaMetadata) Temp() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MediaMetadataStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func MediaMetadataAddTimeId(builder *flatbuffers.Builder, timeId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(timeId), 0)
}
func MediaMetadataAddOwnerId(builder *flatbuffers.Builder, ownerId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ownerId), 0)
}
func MediaMetadataAddTimestamp(builder *flatbuffers.Builder, timestamp int64) {
	builder.PrependInt64Slot(2, timestamp, 0)
}
func MediaMetadataAddMime(builder *flatbuffers.Builder, mime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(mime), 0)
}
func MediaMetadataAddIsReversed(builder *flatbuffers.Builder, isReversed bool) {
	builder.PrependBoolSlot(4, isReversed, false)
}
func MediaMetadataAddIsEncrypted(builder *flatbuffers.Builder, isEncrypted bool) {
	builder.PrependBoolSlot(5, isEncrypted, false)
}
func MediaMetadataAddIsPaidToView(builder *flatbuffers.Builder, isPaidToView bool) {
	builder.PrependBoolSlot(6, isPaidToView, false)
}
func MediaMetadataAddIsPaidToOwn(builder *flatbuffers.Builder, isPaidToOwn bool) {
	builder.PrependBoolSlot(7, isPaidToOwn, false)
}
func MediaMetadataAddIsLocked(builder *flatbuffers.Builder, isLocked bool) {
	builder.PrependBoolSlot(8, isLocked, false)
}
func MediaMetadataAddIsSaved(builder *flatbuffers.Builder, isSaved bool) {
	builder.PrependBoolSlot(9, isSaved, false)
}
func MediaMetadataAddTemp(builder *flatbuffers.Builder, temp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(temp), 0)
}
func MediaMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
