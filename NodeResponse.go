// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatgen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NodeResponseT struct {
	Nodes []*NodeT `json:"nodes"`
}

func (t *NodeResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nodesOffset := flatbuffers.UOffsetT(0)
	if t.Nodes != nil {
		nodesLength := len(t.Nodes)
		nodesOffsets := make([]flatbuffers.UOffsetT, nodesLength)
		for j := 0; j < nodesLength; j++ {
			nodesOffsets[j] = t.Nodes[j].Pack(builder)
		}
		NodeResponseStartNodesVector(builder, nodesLength)
		for j := nodesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(nodesOffsets[j])
		}
		nodesOffset = builder.EndVector(nodesLength)
	}
	NodeResponseStart(builder)
	NodeResponseAddNodes(builder, nodesOffset)
	return NodeResponseEnd(builder)
}

func (rcv *NodeResponse) UnPackTo(t *NodeResponseT) {
	nodesLength := rcv.NodesLength()
	t.Nodes = make([]*NodeT, nodesLength)
	for j := 0; j < nodesLength; j++ {
		x := Node{}
		rcv.Nodes(&x, j)
		t.Nodes[j] = x.UnPack()
	}
}

func (rcv *NodeResponse) UnPack() *NodeResponseT {
	if rcv == nil { return nil }
	t := &NodeResponseT{}
	rcv.UnPackTo(t)
	return t
}

type NodeResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsNodeResponse(buf []byte, offset flatbuffers.UOffsetT) *NodeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NodeResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNodeResponse(buf []byte, offset flatbuffers.UOffsetT) *NodeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NodeResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NodeResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NodeResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NodeResponse) Nodes(obj *Node, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *NodeResponse) NodesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NodeResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NodeResponseAddNodes(builder *flatbuffers.Builder, nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(nodes), 0)
}
func NodeResponseStartNodesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NodeResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
