// Code generated by protoc-gen-go.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Dial
*/
package internal

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

type Dial struct {
	ID               *int64   `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UserID           *int64   `protobuf:"varint,2,opt,name=UserID" json:"UserID,omitempty"`
	Name             *string  `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Level            *float64 `protobuf:"fixed64,4,opt,name=Level" json:"Level,omitempty"`
	ModTime          *int64   `protobuf:"varint,5,opt,name=ModTime" json:"ModTime,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Dial) Reset()                    { *m = Dial{} }
func (m *Dial) String() string            { return proto.CompactTextString(m) }
func (*Dial) ProtoMessage()               {}
func (*Dial) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Dial) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Dial) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Dial) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Dial) GetLevel() float64 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Dial) GetModTime() int64 {
	if m != nil && m.ModTime != nil {
		return *m.ModTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Dial)(nil), "internal.Dial")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x8a,
	0xb8, 0x58, 0x5c, 0x32, 0x13, 0x73, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0x80, 0x2c, 0x21, 0x31, 0x2e, 0xb6, 0xd0, 0xe2, 0xd4, 0x22, 0xa0, 0x18, 0x13,
	0x58, 0x0c, 0xca, 0x13, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x06, 0x8a, 0x72,
	0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x3e, 0xa9, 0x65, 0xa9, 0x39, 0x12, 0x2c, 0x40, 0x41,
	0xc6, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0xdd, 0x37, 0x3f, 0x25, 0x24, 0x13, 0xa8, 0x98, 0x15,
	0x6c, 0x04, 0x8c, 0x0b, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x0f, 0xf9, 0xb1, 0x8e, 0x00, 0x00,
	0x00,
}
