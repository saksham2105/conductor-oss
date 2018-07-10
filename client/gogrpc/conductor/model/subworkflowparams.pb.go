// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/subworkflowparams.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubWorkflowParams struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version              *_struct.Value `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SubWorkflowParams) Reset()         { *m = SubWorkflowParams{} }
func (m *SubWorkflowParams) String() string { return proto.CompactTextString(m) }
func (*SubWorkflowParams) ProtoMessage()    {}
func (*SubWorkflowParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_subworkflowparams_182a77e44709d20f, []int{0}
}
func (m *SubWorkflowParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubWorkflowParams.Unmarshal(m, b)
}
func (m *SubWorkflowParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubWorkflowParams.Marshal(b, m, deterministic)
}
func (dst *SubWorkflowParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubWorkflowParams.Merge(dst, src)
}
func (m *SubWorkflowParams) XXX_Size() int {
	return xxx_messageInfo_SubWorkflowParams.Size(m)
}
func (m *SubWorkflowParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SubWorkflowParams.DiscardUnknown(m)
}

var xxx_messageInfo_SubWorkflowParams proto.InternalMessageInfo

func (m *SubWorkflowParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubWorkflowParams) GetVersion() *_struct.Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*SubWorkflowParams)(nil), "conductor.proto.SubWorkflowParams")
}

func init() {
	proto.RegisterFile("model/subworkflowparams.proto", fileDescriptor_subworkflowparams_182a77e44709d20f)
}

var fileDescriptor_subworkflowparams_182a77e44709d20f = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0x88, 0x62, 0x3c, 0x88, 0x11, 0xa4, 0xf8, 0x07, 0x8a, 0xa7, 0x9e, 0x12, 0xd1,
	0x9b, 0xc7, 0x7e, 0x82, 0x52, 0x41, 0xd1, 0x5b, 0x92, 0xa6, 0x31, 0x98, 0x64, 0x4a, 0xfe, 0x6c,
	0xf7, 0xe3, 0x2f, 0xa4, 0xed, 0xb2, 0xec, 0xde, 0x66, 0xde, 0xcc, 0xfb, 0xcd, 0xf0, 0xd0, 0xb3,
	0x85, 0x5e, 0x1a, 0x1a, 0x12, 0x9f, 0xc0, 0xff, 0x0f, 0x06, 0xa6, 0x91, 0x79, 0x66, 0x03, 0x19,
	0x3d, 0x44, 0xc0, 0x37, 0x02, 0x5c, 0x9f, 0x44, 0x04, 0x3f, 0x0b, 0x0f, 0x4f, 0x0a, 0x40, 0x19,
	0x49, 0x73, 0xc7, 0xd3, 0x40, 0x43, 0xf4, 0x49, 0xc4, 0x79, 0xfa, 0xf2, 0x83, 0x6e, 0x3f, 0x13,
	0xff, 0x5e, 0x48, 0x6d, 0x26, 0x61, 0x8c, 0xce, 0x1d, 0xb3, 0xb2, 0x2c, 0xaa, 0xa2, 0xbe, 0xea,
	0x72, 0x8d, 0x5f, 0xd1, 0xe5, 0x46, 0xfa, 0xa0, 0xc1, 0x95, 0x67, 0x55, 0x51, 0x5f, 0xbf, 0xdd,
	0x93, 0x19, 0x4c, 0x56, 0x30, 0xf9, 0x62, 0x26, 0xc9, 0x6e, 0x5d, 0x6b, 0x1c, 0x7a, 0x14, 0x60,
	0x89, 0x93, 0x71, 0x30, 0x7a, 0x4b, 0x8e, 0xfe, 0x6a, 0xee, 0x4e, 0xee, 0xb6, 0xfc, 0xf7, 0x43,
	0xe9, 0xf8, 0x97, 0x38, 0x11, 0x60, 0xe9, 0x62, 0xa4, 0x7b, 0x23, 0x15, 0x46, 0x4b, 0x17, 0xa9,
	0x02, 0xe5, 0x47, 0x71, 0xa0, 0xe7, 0x44, 0xf8, 0x45, 0xe6, 0xbe, 0xef, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x91, 0x1f, 0xb2, 0x21, 0x01, 0x00, 0x00,
}
