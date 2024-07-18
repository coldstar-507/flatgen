// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FullMediaT struct {
	Metadata *MediaMetadataT `json:"metadata"`
	Data []byte `json:"data"`
}

func (t *FullMediaT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	metadataOffset := t.Metadata.Pack(builder)
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	FullMediaStart(builder)
	FullMediaAddMetadata(builder, metadataOffset)
	FullMediaAddData(builder, dataOffset)
	return FullMediaEnd(builder)
}

func (rcv *FullMedia) UnPackTo(t *FullMediaT) {
	t.Metadata = rcv.Metadata(nil).UnPack()
	t.Data = rcv.DataBytes()
}

func (rcv *FullMedia) UnPack() *FullMediaT {
	if rcv == nil {
		return nil
	}
	t := &FullMediaT{}
	rcv.UnPackTo(t)
	return t
}

type FullMedia struct {
	_tab flatbuffers.Table
}

func GetRootAsFullMedia(buf []byte, offset flatbuffers.UOffsetT) *FullMedia {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FullMedia{}
	x.Init(buf, n+offset)
	return x
}

func FinishFullMediaBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsFullMedia(buf []byte, offset flatbuffers.UOffsetT) *FullMedia {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FullMedia{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedFullMediaBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *FullMedia) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FullMedia) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FullMedia) Metadata(obj *MediaMetadata) *MediaMetadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MediaMetadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *FullMedia) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FullMedia) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FullMedia) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FullMedia) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func FullMediaStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FullMediaAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metadata), 0)
}
func FullMediaAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(data), 0)
}
func FullMediaStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FullMediaEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
