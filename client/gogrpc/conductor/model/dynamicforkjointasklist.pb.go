// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/dynamicforkjointasklist.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DynamicForkJoinTaskList struct {
	DynamicTasks         []*DynamicForkJoinTask `protobuf:"bytes,1,rep,name=dynamic_tasks,json=dynamicTasks" json:"dynamic_tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DynamicForkJoinTaskList) Reset()         { *m = DynamicForkJoinTaskList{} }
func (m *DynamicForkJoinTaskList) String() string { return proto.CompactTextString(m) }
func (*DynamicForkJoinTaskList) ProtoMessage()    {}
func (*DynamicForkJoinTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamicforkjointasklist_5dc7aa3e0011d25e, []int{0}
}
func (m *DynamicForkJoinTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicForkJoinTaskList.Unmarshal(m, b)
}
func (m *DynamicForkJoinTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicForkJoinTaskList.Marshal(b, m, deterministic)
}
func (dst *DynamicForkJoinTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicForkJoinTaskList.Merge(dst, src)
}
func (m *DynamicForkJoinTaskList) XXX_Size() int {
	return xxx_messageInfo_DynamicForkJoinTaskList.Size(m)
}
func (m *DynamicForkJoinTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicForkJoinTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicForkJoinTaskList proto.InternalMessageInfo

func (m *DynamicForkJoinTaskList) GetDynamicTasks() []*DynamicForkJoinTask {
	if m != nil {
		return m.DynamicTasks
	}
	return nil
}

func init() {
	proto.RegisterType((*DynamicForkJoinTaskList)(nil), "conductor.proto.DynamicForkJoinTaskList")
}

func init() {
	proto.RegisterFile("model/dynamicforkjointasklist.proto", fileDescriptor_dynamicforkjointasklist_5dc7aa3e0011d25e)
}

var fileDescriptor_dynamicforkjointasklist_5dc7aa3e0011d25e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x4e, 0xcb, 0x2f, 0xca, 0xce, 0xca, 0xcf,
	0xcc, 0x2b, 0x49, 0x2c, 0xce, 0xce, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4f, 0xce, 0xcf, 0x4b, 0x29, 0x4d, 0x2e, 0xc9, 0x2f, 0x82, 0x08, 0x48, 0xc9, 0xe3, 0xd4,
	0x05, 0x51, 0xa0, 0x94, 0xc2, 0x25, 0xee, 0x02, 0x91, 0x74, 0xcb, 0x2f, 0xca, 0xf6, 0xca, 0xcf,
	0xcc, 0x0b, 0x49, 0x2c, 0xce, 0xf6, 0xc9, 0x2c, 0x2e, 0x11, 0xf2, 0xe4, 0xe2, 0x85, 0xea, 0x8b,
	0x07, 0x69, 0x28, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd1, 0x43, 0xb3, 0x44, 0x0f,
	0x8b, 0x01, 0x41, 0x3c, 0x50, 0xad, 0x20, 0x4e, 0xb1, 0x53, 0x09, 0x97, 0x74, 0x72, 0x7e, 0xae,
	0x5e, 0x5e, 0x6a, 0x49, 0x5a, 0x4e, 0x66, 0x05, 0xba, 0x01, 0x4e, 0x92, 0x38, 0x9c, 0x10, 0x90,
	0x14, 0x65, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xd5, 0xae,
	0x0f, 0xd7, 0xae, 0x9f, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0xa2, 0x9f, 0x9e, 0x9f, 0x5e, 0x54, 0x90,
	0x8c, 0x24, 0x0e, 0xf6, 0x75, 0x12, 0x1b, 0xd8, 0x74, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x76, 0xa8, 0x2e, 0xed, 0x3b, 0x01, 0x00, 0x00,
}
