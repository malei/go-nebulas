// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package pb

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

// Message is a sample proto buffer message for test purpose.
type Message struct {
	Name string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Time uint64            `protobuf:"fixed64,2,opt,name=time" json:"time,omitempty"`
	Tags []string          `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	Map  map[uint32]string `protobuf:"bytes,4,rep,name=map" json:"map,omitempty" protobuf_key:"fixed32,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body []*Message_Body   `protobuf:"bytes,5,rep,name=body" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Message) GetMap() map[uint32]string {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Message) GetBody() []*Message_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message_Body struct {
	Val []string `protobuf:"bytes,3,rep,name=val" json:"val,omitempty"`
}

func (m *Message_Body) Reset()                    { *m = Message_Body{} }
func (m *Message_Body) String() string            { return proto.CompactTextString(m) }
func (*Message_Body) ProtoMessage()               {}
func (*Message_Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Message_Body) GetVal() []string {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Message_Body)(nil), "pb.Message.Body")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0xc5, 0x30,
	0x10, 0x45, 0x49, 0x93, 0xf7, 0x6a, 0x47, 0x84, 0x12, 0xba, 0x08, 0x5d, 0x05, 0x11, 0xe9, 0x2a,
	0x0b, 0x05, 0x11, 0x97, 0x82, 0xcb, 0x6e, 0xf2, 0x07, 0x09, 0x0d, 0x45, 0x6c, 0x9a, 0xd0, 0xd6,
	0x42, 0x3e, 0xd8, 0xff, 0x90, 0x49, 0x2b, 0xb8, 0x3b, 0x73, 0xe7, 0x30, 0xb9, 0x81, 0x3b, 0xef,
	0xd6, 0xd5, 0x8c, 0x4e, 0xc5, 0x25, 0x6c, 0x81, 0x17, 0xd1, 0xde, 0xff, 0x10, 0x28, 0xfb, 0x23,
	0xe5, 0x1c, 0xd8, 0x6c, 0xbc, 0x13, 0x44, 0x92, 0xae, 0xd2, 0x99, 0x31, 0xdb, 0x3e, 0xbd, 0x13,
	0x85, 0x24, 0xdd, 0x55, 0x67, 0xce, 0x99, 0x19, 0x57, 0x41, 0x25, 0x45, 0x0f, 0x99, 0x3f, 0x02,
	0xf5, 0x26, 0x0a, 0x26, 0x69, 0x77, 0xfb, 0xd4, 0xa8, 0x68, 0xd5, 0x79, 0x55, 0xf5, 0x26, 0x7e,
	0xcc, 0xdb, 0x92, 0x34, 0x0a, 0xfc, 0x01, 0x98, 0x0d, 0x43, 0x12, 0x97, 0x2c, 0xd6, 0xff, 0xc5,
	0xf7, 0x30, 0x24, 0x9d, 0xb7, 0xad, 0x00, 0x86, 0x13, 0xaf, 0x81, 0xee, 0x66, 0x3a, 0x1f, 0x42,
	0x6c, 0x5f, 0xe0, 0xe6, 0xef, 0x20, 0x6e, 0xbf, 0x5c, 0xca, 0x75, 0x4b, 0x8d, 0xc8, 0x1b, 0xb8,
	0xec, 0x66, 0xfa, 0x3e, 0xea, 0x56, 0xfa, 0x18, 0xde, 0x8a, 0x57, 0x62, 0xaf, 0xf9, 0xcb, 0xcf,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x12, 0x56, 0x0f, 0x03, 0x01, 0x00, 0x00,
}
