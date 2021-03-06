// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

/*
Package prototests is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	BigMsg
	SmallMsg
	Packed
*/
package prototests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import gzip "compress/gzip"
import bytes "bytes"
import ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BigMsg contains a field and a message field.
type BigMsg struct {
	Field            *int64    `protobuf:"varint,1,opt,name=Field" json:"Field,omitempty"`
	Msg              *SmallMsg `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BigMsg) Reset()                    { *m = BigMsg{} }
func (m *BigMsg) String() string            { return proto.CompactTextString(m) }
func (*BigMsg) ProtoMessage()               {}
func (*BigMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

func (m *BigMsg) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BigMsg) GetMsg() *SmallMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SmallMsg only contains some native fields.
type SmallMsg struct {
	ScarBusStop      *string  `protobuf:"bytes,1,opt,name=ScarBusStop" json:"ScarBusStop,omitempty"`
	FlightParachute  []uint32 `protobuf:"fixed32,12,rep,name=FlightParachute" json:"FlightParachute,omitempty"`
	MapShark         *string  `protobuf:"bytes,18,opt,name=MapShark" json:"MapShark,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMsg) Reset()                    { *m = SmallMsg{} }
func (m *SmallMsg) String() string            { return proto.CompactTextString(m) }
func (*SmallMsg) ProtoMessage()               {}
func (*SmallMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

func (m *SmallMsg) GetScarBusStop() string {
	if m != nil && m.ScarBusStop != nil {
		return *m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func (m *SmallMsg) GetMapShark() string {
	if m != nil && m.MapShark != nil {
		return *m.MapShark
	}
	return ""
}

// Packed contains some repeated packed fields.
type Packed struct {
	Ints             []int64   `protobuf:"varint,4,rep,packed,name=Ints" json:"Ints,omitempty"`
	Floats           []float64 `protobuf:"fixed64,5,rep,packed,name=Floats" json:"Floats,omitempty"`
	Uints            []uint32  `protobuf:"varint,6,rep,packed,name=Uints" json:"Uints,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Packed) Reset()                    { *m = Packed{} }
func (m *Packed) String() string            { return proto.CompactTextString(m) }
func (*Packed) ProtoMessage()               {}
func (*Packed) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{2} }

func (m *Packed) GetInts() []int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *Packed) GetFloats() []float64 {
	if m != nil {
		return m.Floats
	}
	return nil
}

func (m *Packed) GetUints() []uint32 {
	if m != nil {
		return m.Uints
	}
	return nil
}

func init() {
	proto.RegisterType((*BigMsg)(nil), "prototests.BigMsg")
	proto.RegisterType((*SmallMsg)(nil), "prototests.SmallMsg")
	proto.RegisterType((*Packed)(nil), "prototests.Packed")
}
func (this *BigMsg) Description() (desc *descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *SmallMsg) Description() (desc *descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *Packed) Description() (desc *descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func MsgDescription() (desc *descriptor.FileDescriptorSet) {
	d := &descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3824 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5b, 0x6c, 0x23, 0xd7,
		0x79, 0x36, 0xaf, 0x22, 0x7f, 0x52, 0xd4, 0xe8, 0x48, 0x96, 0xb9, 0x72, 0xbc, 0xab, 0xa5, 0x2f,
		0x2b, 0xdb, 0x8d, 0x36, 0x58, 0x7b, 0xd7, 0x5e, 0x6e, 0x13, 0x97, 0x92, 0x28, 0x85, 0x5b, 0x49,
		0x64, 0x86, 0x92, 0x6f, 0x41, 0x31, 0x38, 0x1a, 0x1e, 0x92, 0xb3, 0x3b, 0x9c, 0x99, 0xcc, 0x0c,
		0x77, 0xad, 0x45, 0x1f, 0x5c, 0xb8, 0x17, 0x04, 0x45, 0xef, 0x05, 0x92, 0xb8, 0x8e, 0x7b, 0x01,
		0x5a, 0xb7, 0xe9, 0x2d, 0x69, 0xda, 0xf4, 0xf2, 0xd4, 0x97, 0xb4, 0x7d, 0x2a, 0x90, 0xf7, 0x3e,
		0x34, 0xad, 0x81, 0xa6, 0xad, 0xdb, 0xa4, 0xad, 0x1f, 0x02, 0xf8, 0xa5, 0x38, 0xb7, 0xe1, 0x0c,
		0x49, 0xed, 0x50, 0x01, 0xec, 0x3c, 0x49, 0xe7, 0x3f, 0xff, 0xf7, 0xcd, 0x3f, 0xff, 0xf9, 0xcf,
		0xff, 0xff, 0xe7, 0x0c, 0xe1, 0xbb, 0xd7, 0x61, 0xad, 0x67, 0xdb, 0x3d, 0x93, 0x5c, 0x76, 0x5c,
		0xdb, 0xb7, 0x8f, 0x87, 0xdd, 0xcb, 0x1d, 0xe2, 0xe9, 0xae, 0xe1, 0xf8, 0xb6, 0xbb, 0xc1, 0x64,
		0x68, 0x81, 0x6b, 0x6c, 0x48, 0x8d, 0xca, 0x3e, 0x2c, 0xee, 0x18, 0x26, 0xd9, 0x0e, 0x14, 0xdb,
		0xc4, 0x47, 0xcf, 0x43, 0xba, 0x6b, 0x98, 0xa4, 0x9c, 0x58, 0x4b, 0xad, 0x17, 0xae, 0x3c, 0xb6,
		0x31, 0x06, 0xda, 0x88, 0x22, 0x5a, 0x54, 0xac, 0x32, 0x44, 0xe5, 0xdd, 0x34, 0x2c, 0x4d, 0x99,
		0x45, 0x08, 0xd2, 0x16, 0x1e, 0x50, 0xc6, 0xc4, 0x7a, 0x5e, 0x65, 0xff, 0xa3, 0x32, 0xcc, 0x39,
		0x58, 0xbf, 0x8d, 0x7b, 0xa4, 0x9c, 0x64, 0x62, 0x39, 0x44, 0xe7, 0x01, 0x3a, 0xc4, 0x21, 0x56,
		0x87, 0x58, 0xfa, 0x49, 0x39, 0xb5, 0x96, 0x5a, 0xcf, 0xab, 0x21, 0x09, 0x7a, 0x1a, 0x16, 0x9d,
		0xe1, 0xb1, 0x69, 0xe8, 0x5a, 0x48, 0x0d, 0xd6, 0x52, 0xeb, 0x19, 0x55, 0xe1, 0x13, 0xdb, 0x23,
		0xe5, 0x4b, 0xb0, 0x70, 0x97, 0xe0, 0xdb, 0x61, 0xd5, 0x02, 0x53, 0x2d, 0x51, 0x71, 0x48, 0x71,
		0x0b, 0x8a, 0x03, 0xe2, 0x79, 0xb8, 0x47, 0x34, 0xff, 0xc4, 0x21, 0xe5, 0x34, 0x7b, 0xfb, 0xb5,
		0x89, 0xb7, 0x1f, 0x7f, 0xf3, 0x82, 0x40, 0x1d, 0x9e, 0x38, 0x04, 0xd5, 0x20, 0x4f, 0xac, 0xe1,
		0x80, 0x33, 0x64, 0x4e, 0xf1, 0x5f, 0xdd, 0x1a, 0x0e, 0xc6, 0x59, 0x72, 0x14, 0x26, 0x28, 0xe6,
		0x3c, 0xe2, 0xde, 0x31, 0x74, 0x52, 0xce, 0x32, 0x82, 0x4b, 0x13, 0x04, 0x6d, 0x3e, 0x3f, 0xce,
		0x21, 0x71, 0x68, 0x0b, 0xf2, 0xe4, 0x35, 0x9f, 0x58, 0x9e, 0x61, 0x5b, 0xe5, 0x39, 0x46, 0xf2,
		0xf8, 0x94, 0x55, 0x24, 0x66, 0x67, 0x9c, 0x62, 0x84, 0x43, 0xd7, 0x60, 0xce, 0x76, 0x7c, 0xc3,
		0xb6, 0xbc, 0x72, 0x6e, 0x2d, 0xb1, 0x5e, 0xb8, 0xf2, 0xb1, 0xa9, 0x81, 0xd0, 0xe4, 0x3a, 0xaa,
		0x54, 0x46, 0x0d, 0x50, 0x3c, 0x7b, 0xe8, 0xea, 0x44, 0xd3, 0xed, 0x0e, 0xd1, 0x0c, 0xab, 0x6b,
		0x97, 0xf3, 0x8c, 0xe0, 0xc2, 0xe4, 0x8b, 0x30, 0xc5, 0x2d, 0xbb, 0x43, 0x1a, 0x56, 0xd7, 0x56,
		0x4b, 0x5e, 0x64, 0x8c, 0x56, 0x20, 0xeb, 0x9d, 0x58, 0x3e, 0x7e, 0xad, 0x5c, 0x64, 0x11, 0x22,
		0x46, 0x95, 0xbf, 0xce, 0xc2, 0xc2, 0x2c, 0x21, 0x76, 0x03, 0x32, 0x5d, 0xfa, 0x96, 0xe5, 0xe4,
		0x59, 0x7c, 0xc0, 0x31, 0x51, 0x27, 0x66, 0x7f, 0x40, 0x27, 0xd6, 0xa0, 0x60, 0x11, 0xcf, 0x27,
		0x1d, 0x1e, 0x11, 0xa9, 0x19, 0x63, 0x0a, 0x38, 0x68, 0x32, 0xa4, 0xd2, 0x3f, 0x50, 0x48, 0xbd,
		0x0c, 0x0b, 0x81, 0x49, 0x9a, 0x8b, 0xad, 0x9e, 0x8c, 0xcd, 0xcb, 0x71, 0x96, 0x6c, 0xd4, 0x25,
		0x4e, 0xa5, 0x30, 0xb5, 0x44, 0x22, 0x63, 0xb4, 0x0d, 0x60, 0x5b, 0xc4, 0xee, 0x6a, 0x1d, 0xa2,
		0x9b, 0xe5, 0xdc, 0x29, 0x5e, 0x6a, 0x52, 0x95, 0x09, 0x2f, 0xd9, 0x5c, 0xaa, 0x9b, 0xe8, 0xfa,
		0x28, 0xd4, 0xe6, 0x4e, 0x89, 0x94, 0x7d, 0xbe, 0xc9, 0x26, 0xa2, 0xed, 0x08, 0x4a, 0x2e, 0xa1,
		0x71, 0x4f, 0x3a, 0xe2, 0xcd, 0xf2, 0xcc, 0x88, 0x8d, 0xd8, 0x37, 0x53, 0x05, 0x8c, 0xbf, 0xd8,
		0xbc, 0x1b, 0x1e, 0xa2, 0x47, 0x21, 0x10, 0x68, 0x2c, 0xac, 0x80, 0x65, 0xa1, 0xa2, 0x14, 0x1e,
		0xe0, 0x01, 0x59, 0xbd, 0x07, 0xa5, 0xa8, 0x7b, 0xd0, 0x32, 0x64, 0x3c, 0x1f, 0xbb, 0x3e, 0x8b,
		0xc2, 0x8c, 0xca, 0x07, 0x48, 0x81, 0x14, 0xb1, 0x3a, 0x2c, 0xcb, 0x65, 0x54, 0xfa, 0x2f, 0xfa,
		0xb1, 0xd1, 0x0b, 0xa7, 0xd8, 0x0b, 0x3f, 0x31, 0xb9, 0xa2, 0x11, 0xe6, 0xf1, 0xf7, 0x5e, 0x7d,
		0x0e, 0xe6, 0x23, 0x2f, 0x30, 0xeb, 0xa3, 0x2b, 0x3f, 0x09, 0x0f, 0x4e, 0xa5, 0x46, 0x2f, 0xc3,
		0xf2, 0xd0, 0x32, 0x2c, 0x9f, 0xb8, 0x8e, 0x4b, 0x68, 0xc4, 0xf2, 0x47, 0x95, 0xff, 0x6d, 0xee,
		0x94, 0x98, 0x3b, 0x0a, 0x6b, 0x73, 0x16, 0x75, 0x69, 0x38, 0x29, 0x7c, 0x2a, 0x9f, 0xfb, 0xce,
		0x9c, 0xf2, 0xfa, 0xeb, 0xaf, 0xbf, 0x9e, 0xac, 0x7c, 0x31, 0x0b, 0xcb, 0xd3, 0xf6, 0xcc, 0xd4,
		0xed, 0xbb, 0x02, 0x59, 0x6b, 0x38, 0x38, 0x26, 0x2e, 0x73, 0x52, 0x46, 0x15, 0x23, 0x54, 0x83,
		0x8c, 0x89, 0x8f, 0x89, 0x59, 0x4e, 0xaf, 0x25, 0xd6, 0x4b, 0x57, 0x9e, 0x9e, 0x69, 0x57, 0x6e,
		0xec, 0x51, 0x88, 0xca, 0x91, 0xe8, 0x53, 0x90, 0x16, 0x29, 0x9a, 0x32, 0x3c, 0x35, 0x1b, 0x03,
		0xdd, 0x4b, 0x2a, 0xc3, 0xa1, 0x87, 0x21, 0x4f, 0xff, 0xf2, 0xd8, 0xc8, 0x32, 0x9b, 0x73, 0x54,
		0x40, 0xe3, 0x02, 0xad, 0x42, 0x8e, 0x6d, 0x93, 0x0e, 0x91, 0xa5, 0x2d, 0x18, 0xd3, 0xc0, 0xea,
		0x90, 0x2e, 0x1e, 0x9a, 0xbe, 0x76, 0x07, 0x9b, 0x43, 0xc2, 0x02, 0x3e, 0xaf, 0x16, 0x85, 0xf0,
		0x45, 0x2a, 0x43, 0x17, 0xa0, 0xc0, 0x77, 0x95, 0x61, 0x75, 0xc8, 0x6b, 0x2c, 0x7b, 0x66, 0x54,
		0xbe, 0xd1, 0x1a, 0x54, 0x42, 0x1f, 0x7f, 0xcb, 0xb3, 0x2d, 0x19, 0x9a, 0xec, 0x11, 0x54, 0xc0,
		0x1e, 0xff, 0xdc, 0x78, 0xe2, 0x7e, 0x64, 0xfa, 0xeb, 0x8d, 0xc7, 0x54, 0xe5, 0x1b, 0x49, 0x48,
		0xb3, 0x7c, 0xb1, 0x00, 0x85, 0xc3, 0x57, 0x5a, 0x75, 0x6d, 0xbb, 0x79, 0xb4, 0xb9, 0x57, 0x57,
		0x12, 0xa8, 0x04, 0xc0, 0x04, 0x3b, 0x7b, 0xcd, 0xda, 0xa1, 0x92, 0x0c, 0xc6, 0x8d, 0x83, 0xc3,
		0x6b, 0xcf, 0x2a, 0xa9, 0x00, 0x70, 0xc4, 0x05, 0xe9, 0xb0, 0xc2, 0x33, 0x57, 0x94, 0x0c, 0x52,
		0xa0, 0xc8, 0x09, 0x1a, 0x2f, 0xd7, 0xb7, 0xaf, 0x3d, 0xab, 0x64, 0xa3, 0x92, 0x67, 0xae, 0x28,
		0x73, 0x68, 0x1e, 0xf2, 0x4c, 0xb2, 0xd9, 0x6c, 0xee, 0x29, 0xb9, 0x80, 0xb3, 0x7d, 0xa8, 0x36,
		0x0e, 0x76, 0x95, 0x7c, 0xc0, 0xb9, 0xab, 0x36, 0x8f, 0x5a, 0x0a, 0x04, 0x0c, 0xfb, 0xf5, 0x76,
		0xbb, 0xb6, 0x5b, 0x57, 0x0a, 0x81, 0xc6, 0xe6, 0x2b, 0x87, 0xf5, 0xb6, 0x52, 0x8c, 0x98, 0xf5,
		0xcc, 0x15, 0x65, 0x3e, 0x78, 0x44, 0xfd, 0xe0, 0x68, 0x5f, 0x29, 0xa1, 0x45, 0x98, 0xe7, 0x8f,
		0x90, 0x46, 0x2c, 0x8c, 0x89, 0xae, 0x3d, 0xab, 0x28, 0x23, 0x43, 0x38, 0xcb, 0x62, 0x44, 0x70,
		0xed, 0x59, 0x05, 0x55, 0xb6, 0x20, 0xc3, 0xa2, 0x0b, 0x21, 0x28, 0xed, 0xd5, 0x36, 0xeb, 0x7b,
		0x5a, 0xb3, 0x75, 0xd8, 0x68, 0x1e, 0xd4, 0xf6, 0x94, 0xc4, 0x48, 0xa6, 0xd6, 0x3f, 0x73, 0xd4,
		0x50, 0xeb, 0xdb, 0x4a, 0x32, 0x2c, 0x6b, 0xd5, 0x6b, 0x87, 0xf5, 0x6d, 0x25, 0x55, 0xd1, 0x61,
		0x79, 0x5a, 0x9e, 0x9c, 0xba, 0x33, 0x42, 0x4b, 0x9c, 0x3c, 0x65, 0x89, 0x19, 0xd7, 0xc4, 0x12,
		0xff, 0x6b, 0x12, 0x96, 0xa6, 0xd4, 0x8a, 0xa9, 0x0f, 0x79, 0x01, 0x32, 0x3c, 0x44, 0x79, 0xf5,
		0x7c, 0x72, 0x6a, 0xd1, 0x61, 0x01, 0x3b, 0x51, 0x41, 0x19, 0x2e, 0xdc, 0x41, 0xa4, 0x4e, 0xe9,
		0x20, 0x28, 0xc5, 0x44, 0x4e, 0xff, 0x89, 0x89, 0x9c, 0xce, 0xcb, 0xde, 0xb5, 0x59, 0xca, 0x1e,
		0x93, 0x9d, 0x2d, 0xb7, 0x67, 0xa6, 0xe4, 0xf6, 0x1b, 0xb0, 0x38, 0x41, 0x34, 0x73, 0x8e, 0x7d,
		0x23, 0x01, 0xe5, 0xd3, 0x9c, 0x13, 0x93, 0xe9, 0x92, 0x91, 0x4c, 0x77, 0x63, 0xdc, 0x83, 0x17,
		0x4f, 0x5f, 0x84, 0x89, 0xb5, 0x7e, 0x27, 0x01, 0x2b, 0xd3, 0x3b, 0xc5, 0xa9, 0x36, 0x7c, 0x0a,
		0xb2, 0x03, 0xe2, 0xf7, 0x6d, 0xd9, 0x2d, 0x3d, 0x31, 0xa5, 0x06, 0xd3, 0xe9, 0xf1, 0xc5, 0x16,
		0xa8, 0x70, 0x11, 0x4f, 0x9d, 0xd6, 0xee, 0x71, 0x6b, 0x26, 0x2c, 0xfd, 0x7c, 0x12, 0x1e, 0x9c,
		0x4a, 0x3e, 0xd5, 0xd0, 0x47, 0x00, 0x0c, 0xcb, 0x19, 0xfa, 0xbc, 0x23, 0xe2, 0x09, 0x36, 0xcf,
		0x24, 0x2c, 0x79, 0xd1, 0xe4, 0x39, 0xf4, 0x83, 0xf9, 0x14, 0x9b, 0x07, 0x2e, 0x62, 0x0a, 0xcf,
		0x8f, 0x0c, 0x4d, 0x33, 0x43, 0xcf, 0x9f, 0xf2, 0xa6, 0x13, 0x81, 0xf9, 0x09, 0x50, 0x74, 0xd3,
		0x20, 0x96, 0xaf, 0x79, 0xbe, 0x4b, 0xf0, 0xc0, 0xb0, 0x7a, 0xac, 0x82, 0xe4, 0xaa, 0x99, 0x2e,
		0x36, 0x3d, 0xa2, 0x2e, 0xf0, 0xe9, 0xb6, 0x9c, 0xa5, 0x08, 0x16, 0x40, 0x6e, 0x08, 0x91, 0x8d,
		0x20, 0xf8, 0x74, 0x80, 0xa8, 0x7c, 0x3d, 0x07, 0x85, 0x50, 0x5f, 0x8d, 0x2e, 0x42, 0xf1, 0x16,
		0xbe, 0x83, 0x35, 0x79, 0x56, 0xe2, 0x9e, 0x28, 0x50, 0x59, 0x4b, 0x9c, 0x97, 0x3e, 0x01, 0xcb,
		0x4c, 0xc5, 0x1e, 0xfa, 0xc4, 0xd5, 0x74, 0x13, 0x7b, 0x1e, 0x73, 0x5a, 0x8e, 0xa9, 0x22, 0x3a,
		0xd7, 0xa4, 0x53, 0x5b, 0x72, 0x06, 0x5d, 0x85, 0x25, 0x86, 0x18, 0x0c, 0x4d, 0xdf, 0x70, 0x4c,
		0xa2, 0xd1, 0xd3, 0x9b, 0xc7, 0x2a, 0x49, 0x60, 0xd9, 0x22, 0xd5, 0xd8, 0x17, 0x0a, 0xd4, 0x22,
		0x0f, 0x6d, 0xc3, 0x23, 0x0c, 0xd6, 0x23, 0x16, 0x71, 0xb1, 0x4f, 0x34, 0xf2, 0xb9, 0x21, 0x36,
		0x3d, 0x0d, 0x5b, 0x1d, 0xad, 0x8f, 0xbd, 0x7e, 0x79, 0x99, 0x12, 0x6c, 0x26, 0xcb, 0x09, 0xf5,
		0x1c, 0x55, 0xdc, 0x15, 0x7a, 0x75, 0xa6, 0x56, 0xb3, 0x3a, 0x9f, 0xc6, 0x5e, 0x1f, 0x55, 0x61,
		0x85, 0xb1, 0x78, 0xbe, 0x6b, 0x58, 0x3d, 0x4d, 0xef, 0x13, 0xfd, 0xb6, 0x36, 0xf4, 0xbb, 0xcf,
		0x97, 0x1f, 0x0e, 0x3f, 0x9f, 0x59, 0xd8, 0x66, 0x3a, 0x5b, 0x54, 0xe5, 0xc8, 0xef, 0x3e, 0x8f,
		0xda, 0x50, 0xa4, 0x8b, 0x31, 0x30, 0xee, 0x11, 0xad, 0x6b, 0xbb, 0xac, 0x34, 0x96, 0xa6, 0xa4,
		0xa6, 0x90, 0x07, 0x37, 0x9a, 0x02, 0xb0, 0x6f, 0x77, 0x48, 0x35, 0xd3, 0x6e, 0xd5, 0xeb, 0xdb,
		0x6a, 0x41, 0xb2, 0xec, 0xd8, 0x2e, 0x0d, 0xa8, 0x9e, 0x1d, 0x38, 0xb8, 0xc0, 0x03, 0xaa, 0x67,
		0x4b, 0xf7, 0x5e, 0x85, 0x25, 0x5d, 0xe7, 0xef, 0x6c, 0xe8, 0x9a, 0x38, 0x63, 0x79, 0x65, 0x25,
		0xe2, 0x2c, 0x5d, 0xdf, 0xe5, 0x0a, 0x22, 0xc6, 0x3d, 0x74, 0x1d, 0x1e, 0x1c, 0x39, 0x2b, 0x0c,
		0x5c, 0x9c, 0x78, 0xcb, 0x71, 0xe8, 0x55, 0x58, 0x72, 0x4e, 0x26, 0x81, 0x28, 0xf2, 0x44, 0xe7,
		0x64, 0x1c, 0xf6, 0x1c, 0x2c, 0x3b, 0x7d, 0x67, 0x12, 0xf7, 0x54, 0x18, 0x87, 0x9c, 0xbe, 0x33,
		0x0e, 0x7c, 0x9c, 0x1d, 0xb8, 0x5d, 0xa2, 0x63, 0x9f, 0x74, 0xca, 0x0f, 0x85, 0xd5, 0x43, 0x13,
		0xe8, 0x32, 0x28, 0xba, 0xae, 0x11, 0x0b, 0x1f, 0x9b, 0x44, 0xc3, 0x2e, 0xb1, 0xb0, 0x57, 0xbe,
		0x10, 0x56, 0x2e, 0xe9, 0x7a, 0x9d, 0xcd, 0xd6, 0xd8, 0x24, 0x7a, 0x0a, 0x16, 0xed, 0xe3, 0x5b,
		0x3a, 0x0f, 0x49, 0xcd, 0x71, 0x49, 0xd7, 0x78, 0xad, 0xfc, 0x18, 0xf3, 0xef, 0x02, 0x9d, 0x60,
		0x01, 0xd9, 0x62, 0x62, 0xf4, 0x24, 0x28, 0xba, 0xd7, 0xc7, 0xae, 0xc3, 0x72, 0xb2, 0xe7, 0x60,
		0x9d, 0x94, 0x1f, 0xe7, 0xaa, 0x5c, 0x7e, 0x20, 0xc5, 0x74, 0x4b, 0x78, 0x77, 0x8d, 0xae, 0x2f,
		0x19, 0x2f, 0xf1, 0x2d, 0xc1, 0x64, 0x82, 0x6d, 0x1d, 0x14, 0xea, 0x8a, 0xc8, 0x83, 0xd7, 0x99,
		0x5a, 0xc9, 0xe9, 0x3b, 0xe1, 0xe7, 0x3e, 0x0a, 0xf3, 0x54, 0x73, 0xf4, 0xd0, 0x27, 0x79, 0x43,
		0xe6, 0xf4, 0x43, 0x4f, 0xfc, 0xd0, 0x7a, 0xe3, 0x4a, 0x15, 0x8a, 0xe1, 0xf8, 0x44, 0x79, 0xe0,
		0x11, 0xaa, 0x24, 0x68, 0xb3, 0xb2, 0xd5, 0xdc, 0xa6, 0x6d, 0xc6, 0xab, 0x75, 0x25, 0x49, 0xdb,
		0x9d, 0xbd, 0xc6, 0x61, 0x5d, 0x53, 0x8f, 0x0e, 0x0e, 0x1b, 0xfb, 0x75, 0x25, 0x15, 0xee, 0xab,
		0xbf, 0x99, 0x84, 0x52, 0xf4, 0x88, 0x84, 0x7e, 0x14, 0x1e, 0x92, 0xf7, 0x19, 0x1e, 0xf1, 0xb5,
		0xbb, 0x86, 0xcb, 0xb6, 0xcc, 0x00, 0xf3, 0xf2, 0x15, 0x2c, 0xda, 0xb2, 0xd0, 0x6a, 0x13, 0xff,
		0x25, 0xc3, 0xa5, 0x1b, 0x62, 0x80, 0x7d, 0xb4, 0x07, 0x17, 0x2c, 0x5b, 0xf3, 0x7c, 0x6c, 0x75,
		0xb0, 0xdb, 0xd1, 0x46, 0x37, 0x49, 0x1a, 0xd6, 0x75, 0xe2, 0x79, 0x36, 0x2f, 0x55, 0x01, 0xcb,
		0xc7, 0x2c, 0xbb, 0x2d, 0x94, 0x47, 0x39, 0xbc, 0x26, 0x54, 0xc7, 0x02, 0x2c, 0x75, 0x5a, 0x80,
		0x3d, 0x0c, 0xf9, 0x01, 0x76, 0x34, 0x62, 0xf9, 0xee, 0x09, 0x6b, 0x8c, 0x73, 0x6a, 0x6e, 0x80,
		0x9d, 0x3a, 0x1d, 0x7f, 0x34, 0xe7, 0x93, 0x7f, 0x4a, 0x41, 0x31, 0xdc, 0x1c, 0xd3, 0xb3, 0x86,
		0xce, 0xea, 0x48, 0x82, 0x65, 0x9a, 0x47, 0xef, 0xdb, 0x4a, 0x6f, 0x6c, 0xd1, 0x02, 0x53, 0xcd,
		0xf2, 0x96, 0x55, 0xe5, 0x48, 0x5a, 0xdc, 0x69, 0x6e, 0x21, 0xbc, 0x45, 0xc8, 0xa9, 0x62, 0x84,
		0x76, 0x21, 0x7b, 0xcb, 0x63, 0xdc, 0x59, 0xc6, 0xfd, 0xd8, 0xfd, 0xb9, 0x6f, 0xb6, 0x19, 0x79,
		0xfe, 0x66, 0x5b, 0x3b, 0x68, 0xaa, 0xfb, 0xb5, 0x3d, 0x55, 0xc0, 0xd1, 0x39, 0x48, 0x9b, 0xf8,
		0xde, 0x49, 0xb4, 0x14, 0x31, 0xd1, 0xac, 0x8e, 0x3f, 0x07, 0xe9, 0xbb, 0x04, 0xdf, 0x8e, 0x16,
		0x00, 0x26, 0xfa, 0x10, 0x43, 0xff, 0x32, 0x64, 0x98, 0xbf, 0x10, 0x80, 0xf0, 0x98, 0xf2, 0x00,
		0xca, 0x41, 0x7a, 0xab, 0xa9, 0xd2, 0xf0, 0x57, 0xa0, 0xc8, 0xa5, 0x5a, 0xab, 0x51, 0xdf, 0xaa,
		0x2b, 0xc9, 0xca, 0x55, 0xc8, 0x72, 0x27, 0xd0, 0xad, 0x11, 0xb8, 0x41, 0x79, 0x40, 0x0c, 0x05,
		0x47, 0x42, 0xce, 0x1e, 0xed, 0x6f, 0xd6, 0x55, 0x25, 0x19, 0x5e, 0x5e, 0x0f, 0x8a, 0xe1, 0xbe,
		0xf8, 0xa3, 0x89, 0xa9, 0xbf, 0x49, 0x40, 0x21, 0xd4, 0xe7, 0xd2, 0x06, 0x05, 0x9b, 0xa6, 0x7d,
		0x57, 0xc3, 0xa6, 0x81, 0x3d, 0x11, 0x14, 0xc0, 0x44, 0x35, 0x2a, 0x99, 0x75, 0xd1, 0x3e, 0x12,
		0xe3, 0xdf, 0x4e, 0x80, 0x32, 0xde, 0x62, 0x8e, 0x19, 0x98, 0xf8, 0xa1, 0x1a, 0xf8, 0x56, 0x02,
		0x4a, 0xd1, 0xbe, 0x72, 0xcc, 0xbc, 0x8b, 0x3f, 0x54, 0xf3, 0xfe, 0x39, 0x09, 0xf3, 0x91, 0x6e,
		0x72, 0x56, 0xeb, 0x3e, 0x07, 0x8b, 0x46, 0x87, 0x0c, 0x1c, 0xdb, 0x27, 0x96, 0x7e, 0xa2, 0x99,
		0xe4, 0x0e, 0x31, 0xcb, 0x15, 0x96, 0x28, 0x2e, 0xdf, 0xbf, 0x5f, 0xdd, 0x68, 0x8c, 0x70, 0x7b,
		0x14, 0x56, 0x5d, 0x6a, 0x6c, 0xd7, 0xf7, 0x5b, 0xcd, 0xc3, 0xfa, 0xc1, 0xd6, 0x2b, 0xda, 0xd1,
		0xc1, 0x8f, 0x1f, 0x34, 0x5f, 0x3a, 0x50, 0x15, 0x63, 0x4c, 0xed, 0x43, 0xdc, 0xea, 0x2d, 0x50,
		0xc6, 0x8d, 0x42, 0x0f, 0xc1, 0x34, 0xb3, 0x94, 0x07, 0xd0, 0x12, 0x2c, 0x1c, 0x34, 0xb5, 0x76,
		0x63, 0xbb, 0xae, 0xd5, 0x77, 0x76, 0xea, 0x5b, 0x87, 0x6d, 0x7e, 0x03, 0x11, 0x68, 0x1f, 0x46,
		0x37, 0xf5, 0x9b, 0x29, 0x58, 0x9a, 0x62, 0x09, 0xaa, 0x89, 0xb3, 0x03, 0x3f, 0xce, 0x7c, 0x7c,
		0x16, 0xeb, 0x37, 0x68, 0xc9, 0x6f, 0x61, 0xd7, 0x17, 0x47, 0x8d, 0x27, 0x81, 0x7a, 0xc9, 0xf2,
		0x8d, 0xae, 0x41, 0x5c, 0x71, 0x61, 0xc3, 0x0f, 0x14, 0x0b, 0x23, 0x39, 0xbf, 0xb3, 0xf9, 0x11,
		0x40, 0x8e, 0xed, 0x19, 0xbe, 0x71, 0x87, 0x68, 0x86, 0x25, 0x6f, 0x77, 0xe8, 0x01, 0x23, 0xad,
		0x2a, 0x72, 0xa6, 0x61, 0xf9, 0x81, 0xb6, 0x45, 0x7a, 0x78, 0x4c, 0x9b, 0x26, 0xf0, 0x94, 0xaa,
		0xc8, 0x99, 0x40, 0xfb, 0x22, 0x14, 0x3b, 0xf6, 0x90, 0x76, 0x5d, 0x5c, 0x8f, 0xd6, 0x8b, 0x84,
		0x5a, 0xe0, 0xb2, 0x40, 0x45, 0xf4, 0xd3, 0xa3, 0x6b, 0xa5, 0xa2, 0x5a, 0xe0, 0x32, 0xae, 0x72,
		0x09, 0x16, 0x70, 0xaf, 0xe7, 0x52, 0x72, 0x49, 0xc4, 0x4f, 0x08, 0xa5, 0x40, 0xcc, 0x14, 0x57,
		0x6f, 0x42, 0x4e, 0xfa, 0x81, 0x96, 0x64, 0xea, 0x09, 0xcd, 0xe1, 0xc7, 0xde, 0xe4, 0x7a, 0x5e,
		0xcd, 0x59, 0x72, 0xf2, 0x22, 0x14, 0x0d, 0x4f, 0x1b, 0xdd, 0x92, 0x27, 0xd7, 0x92, 0xeb, 0x39,
		0xb5, 0x60, 0x78, 0xc1, 0x0d, 0x63, 0xe5, 0x9d, 0x24, 0x94, 0xa2, 0xb7, 0xfc, 0x68, 0x1b, 0x72,
		0xa6, 0xad, 0x63, 0x16, 0x5a, 0xfc, 0x13, 0xd3, 0x7a, 0xcc, 0x87, 0x81, 0x8d, 0x3d, 0xa1, 0xaf,
		0x06, 0xc8, 0xd5, 0x7f, 0x4c, 0x40, 0x4e, 0x8a, 0xd1, 0x0a, 0xa4, 0x1d, 0xec, 0xf7, 0x19, 0x5d,
		0x66, 0x33, 0xa9, 0x24, 0x54, 0x36, 0xa6, 0x72, 0xcf, 0xc1, 0x16, 0x0b, 0x01, 0x21, 0xa7, 0x63,
		0xba, 0xae, 0x26, 0xc1, 0x1d, 0x76, 0xfc, 0xb0, 0x07, 0x03, 0x62, 0xf9, 0x9e, 0x5c, 0x57, 0x21,
		0xdf, 0x12, 0x62, 0xf4, 0x34, 0x2c, 0xfa, 0x2e, 0x36, 0xcc, 0x88, 0x6e, 0x9a, 0xe9, 0x2a, 0x72,
		0x22, 0x50, 0xae, 0xc2, 0x39, 0xc9, 0xdb, 0x21, 0x3e, 0xd6, 0xfb, 0xa4, 0x33, 0x02, 0x65, 0xd9,
		0x35, 0xc3, 0x43, 0x42, 0x61, 0x5b, 0xcc, 0x4b, 0x6c, 0xe5, 0x5b, 0x09, 0x58, 0x94, 0x07, 0xa6,
		0x4e, 0xe0, 0xac, 0x7d, 0x00, 0x6c, 0x59, 0xb6, 0x1f, 0x76, 0xd7, 0x64, 0x28, 0x4f, 0xe0, 0x36,
		0x6a, 0x01, 0x48, 0x0d, 0x11, 0xac, 0x0e, 0x00, 0x46, 0x33, 0xa7, 0xba, 0xed, 0x02, 0x14, 0xc4,
		0x27, 0x1c, 0xf6, 0x1d, 0x90, 0x1f, 0xb1, 0x81, 0x8b, 0xe8, 0xc9, 0x0a, 0x2d, 0x43, 0xe6, 0x98,
		0xf4, 0x0c, 0x4b, 0x5c, 0xcc, 0xf2, 0x81, 0xbc, 0x08, 0x49, 0x07, 0x17, 0x21, 0x9b, 0x9f, 0x85,
		0x25, 0xdd, 0x1e, 0x8c, 0x9b, 0xbb, 0xa9, 0x8c, 0x1d, 0xf3, 0xbd, 0x4f, 0x27, 0x5e, 0x85, 0x51,
		0x8b, 0xf9, 0xfd, 0x44, 0xe2, 0x77, 0x92, 0xa9, 0xdd, 0xd6, 0xe6, 0x57, 0x92, 0xab, 0xbb, 0x1c,
		0xda, 0x92, 0x6f, 0xaa, 0x92, 0xae, 0x49, 0x74, 0x6a, 0x3d, 0x7c, 0xe1, 0x12, 0x7c, 0xbc, 0x67,
		0xf8, 0xfd, 0xe1, 0xf1, 0x86, 0x6e, 0x0f, 0x2e, 0xf7, 0xec, 0x9e, 0x3d, 0xfa, 0xf4, 0x49, 0x47,
		0x6c, 0xc0, 0xfe, 0x13, 0x9f, 0x3f, 0xf3, 0x81, 0x74, 0x35, 0xf6, 0x5b, 0x69, 0xf5, 0x00, 0x96,
		0x84, 0xb2, 0xc6, 0xbe, 0xbf, 0xf0, 0x53, 0x04, 0xba, 0xef, 0x1d, 0x56, 0xf9, 0x6b, 0xef, 0xb2,
		0x72, 0xad, 0x2e, 0x0a, 0x28, 0x9d, 0xe3, 0x07, 0x8d, 0xaa, 0x0a, 0x0f, 0x46, 0xf8, 0xf8, 0xd6,
		0x24, 0x6e, 0x0c, 0xe3, 0x37, 0x05, 0xe3, 0x52, 0x88, 0xb1, 0x2d, 0xa0, 0xd5, 0x2d, 0x98, 0x3f,
		0x0b, 0xd7, 0xdf, 0x09, 0xae, 0x22, 0x09, 0x93, 0xec, 0xc2, 0x02, 0x23, 0xd1, 0x87, 0x9e, 0x6f,
		0x0f, 0x58, 0xde, 0xbb, 0x3f, 0xcd, 0xdf, 0xbf, 0xcb, 0xf7, 0x4a, 0x89, 0xc2, 0xb6, 0x02, 0x54,
		0xb5, 0x0a, 0xec, 0x93, 0x53, 0x87, 0xe8, 0x66, 0x0c, 0xc3, 0x3f, 0x08, 0x43, 0x02, 0xfd, 0xea,
		0x8b, 0xb0, 0x4c, 0xff, 0x67, 0x69, 0x29, 0x6c, 0x49, 0xfc, 0x85, 0x57, 0xf9, 0x5b, 0x6f, 0xf0,
		0xed, 0xb8, 0x14, 0x10, 0x84, 0x6c, 0x0a, 0xad, 0x62, 0x8f, 0xf8, 0x3e, 0x71, 0x3d, 0x0d, 0x9b,
		0xd3, 0xcc, 0x0b, 0xdd, 0x18, 0x94, 0xbf, 0xf4, 0x5e, 0x74, 0x15, 0x77, 0x39, 0xb2, 0x66, 0x9a,
		0xd5, 0x23, 0x78, 0x68, 0x4a, 0x54, 0xcc, 0xc0, 0xf9, 0xa6, 0xe0, 0x5c, 0x9e, 0x88, 0x0c, 0x4a,
		0xdb, 0x02, 0x29, 0x0f, 0xd6, 0x72, 0x06, 0xce, 0xdf, 0x10, 0x9c, 0x48, 0x60, 0xe5, 0x92, 0x52,
		0xc6, 0x9b, 0xb0, 0x78, 0x87, 0xb8, 0xc7, 0xb6, 0x27, 0x6e, 0x69, 0x66, 0xa0, 0x7b, 0x4b, 0xd0,
		0x2d, 0x08, 0x20, 0xbb, 0xb6, 0xa1, 0x5c, 0xd7, 0x21, 0xd7, 0xc5, 0x3a, 0x99, 0x81, 0xe2, 0xcb,
		0x82, 0x62, 0x8e, 0xea, 0x53, 0x68, 0x0d, 0x8a, 0x3d, 0x5b, 0x54, 0xa6, 0x78, 0xf8, 0xdb, 0x02,
		0x5e, 0x90, 0x18, 0x41, 0xe1, 0xd8, 0xce, 0xd0, 0xa4, 0x65, 0x2b, 0x9e, 0xe2, 0x37, 0x25, 0x85,
		0xc4, 0x08, 0x8a, 0x33, 0xb8, 0xf5, 0xb7, 0x24, 0x85, 0x17, 0xf2, 0xe7, 0x0b, 0x50, 0xb0, 0x2d,
		0xf3, 0xc4, 0xb6, 0x66, 0x31, 0xe2, 0xb7, 0x05, 0x03, 0x08, 0x08, 0x25, 0xb8, 0x01, 0xf9, 0x59,
		0x17, 0xe2, 0x77, 0xdf, 0x93, 0xdb, 0x43, 0xae, 0xc0, 0x2e, 0x2c, 0xc8, 0x04, 0x65, 0xd8, 0xd6,
		0x0c, 0x14, 0xbf, 0x27, 0x28, 0x4a, 0x21, 0x98, 0x78, 0x0d, 0x9f, 0x78, 0x7e, 0x8f, 0xcc, 0x42,
		0xf2, 0x8e, 0x7c, 0x0d, 0x01, 0x11, 0xae, 0x3c, 0x26, 0x96, 0xde, 0x9f, 0x8d, 0xe1, 0xf7, 0xa5,
		0x2b, 0x25, 0x86, 0x52, 0x6c, 0xc1, 0xfc, 0x00, 0xbb, 0x5e, 0x1f, 0x9b, 0x33, 0x2d, 0xc7, 0x1f,
		0x08, 0x8e, 0x62, 0x00, 0x12, 0x1e, 0x19, 0x5a, 0x67, 0xa1, 0xf9, 0x8a, 0xf4, 0x48, 0x08, 0x26,
		0xb6, 0x9e, 0xe7, 0xb3, 0x2b, 0xad, 0xb3, 0xb0, 0xfd, 0xa1, 0xdc, 0x7a, 0x1c, 0xbb, 0x1f, 0x66,
		0xbc, 0x01, 0x79, 0xcf, 0xb8, 0x37, 0x13, 0xcd, 0x1f, 0xc9, 0x95, 0x66, 0x00, 0x0a, 0x7e, 0x05,
		0xce, 0x4d, 0x2d, 0x13, 0x33, 0x90, 0xfd, 0xb1, 0x20, 0x5b, 0x99, 0x52, 0x2a, 0x44, 0x4a, 0x38,
		0x2b, 0xe5, 0x9f, 0xc8, 0x94, 0x40, 0xc6, 0xb8, 0x5a, 0xf4, 0xac, 0xe0, 0xe1, 0xee, 0xd9, 0xbc,
		0xf6, 0xa7, 0xd2, 0x6b, 0x1c, 0x1b, 0xf1, 0xda, 0x21, 0xac, 0x08, 0xc6, 0xb3, 0xad, 0xeb, 0x57,
		0x65, 0x62, 0xe5, 0xe8, 0xa3, 0xe8, 0xea, 0x7e, 0x16, 0x56, 0x03, 0x77, 0xca, 0xa6, 0xd4, 0xd3,
		0x06, 0xd8, 0x99, 0x81, 0xf9, 0x6b, 0x82, 0x59, 0x66, 0xfc, 0xa0, 0xab, 0xf5, 0xf6, 0xb1, 0x43,
		0xc9, 0x5f, 0x86, 0xb2, 0x24, 0x1f, 0x5a, 0x2e, 0xd1, 0xed, 0x9e, 0x65, 0xdc, 0x23, 0x9d, 0x19,
		0xa8, 0xff, 0x6c, 0x6c, 0xa9, 0x8e, 0x42, 0x70, 0xca, 0xdc, 0x00, 0x25, 0xe8, 0x55, 0x34, 0x63,
		0xe0, 0xd8, 0xae, 0x1f, 0xc3, 0xf8, 0x75, 0xb9, 0x52, 0x01, 0xae, 0xc1, 0x60, 0xd5, 0x3a, 0x94,
		0xd8, 0x70, 0xd6, 0x90, 0xfc, 0x73, 0x41, 0x34, 0x3f, 0x42, 0x89, 0xc4, 0xa1, 0xdb, 0x03, 0x07,
		0xbb, 0xb3, 0xe4, 0xbf, 0xbf, 0x90, 0x89, 0x43, 0x40, 0x44, 0xe2, 0xf0, 0x4f, 0x1c, 0x42, 0xab,
		0xfd, 0x0c, 0x0c, 0xdf, 0x90, 0x89, 0x43, 0x62, 0x04, 0x85, 0x6c, 0x18, 0x66, 0xa0, 0xf8, 0x4b,
		0x49, 0x21, 0x31, 0x94, 0xe2, 0x33, 0xa3, 0x42, 0xeb, 0x92, 0x9e, 0xe1, 0xf9, 0x2e, 0x6f, 0x85,
		0xef, 0x4f, 0xf5, 0x57, 0xef, 0x45, 0x9b, 0x30, 0x35, 0x04, 0xad, 0xde, 0x84, 0x85, 0xb1, 0x16,
		0x03, 0xc5, 0xfd, 0x7e, 0xa5, 0xfc, 0x53, 0xef, 0x8b, 0x64, 0x14, 0xed, 0x30, 0xaa, 0x7b, 0x74,
		0xdd, 0xa3, 0x7d, 0x40, 0x3c, 0xd9, 0x1b, 0xef, 0x07, 0x4b, 0x1f, 0x69, 0x03, 0xaa, 0x3b, 0x30,
		0x1f, 0xe9, 0x01, 0xe2, 0xa9, 0x7e, 0x5a, 0x50, 0x15, 0xc3, 0x2d, 0x40, 0xf5, 0x2a, 0xa4, 0x69,
		0x3d, 0x8f, 0x87, 0xff, 0x8c, 0x80, 0x33, 0xf5, 0xea, 0x27, 0x21, 0x27, 0xeb, 0x78, 0x3c, 0xf4,
		0x67, 0x05, 0x34, 0x80, 0x50, 0xb8, 0xac, 0xe1, 0xf1, 0xf0, 0x9f, 0x93, 0x70, 0x09, 0xa1, 0xf0,
		0xd9, 0x5d, 0xf8, 0xb7, 0x3f, 0x9f, 0x16, 0x79, 0x58, 0xfa, 0xee, 0x06, 0xcc, 0x89, 0xe2, 0x1d,
		0x8f, 0xfe, 0xbc, 0x78, 0xb8, 0x44, 0x54, 0x9f, 0x83, 0xcc, 0x8c, 0x0e, 0xff, 0x05, 0x01, 0xe5,
		0xfa, 0xd5, 0x2d, 0x28, 0x84, 0x0a, 0x76, 0x3c, 0xfc, 0x17, 0x05, 0x3c, 0x8c, 0xa2, 0xa6, 0x8b,
		0x82, 0x1d, 0x4f, 0xf0, 0x4b, 0xd2, 0x74, 0x81, 0xa0, 0x6e, 0x93, 0xb5, 0x3a, 0x1e, 0xfd, 0xcb,
		0xd2, 0xeb, 0x12, 0x52, 0x7d, 0x01, 0xf2, 0x41, 0xfe, 0x8d, 0xc7, 0xff, 0x8a, 0xc0, 0x8f, 0x30,
		0xd4, 0x03, 0xa1, 0xfc, 0x1f, 0x4f, 0xf1, 0xab, 0xd2, 0x03, 0x21, 0x14, 0xdd, 0x46, 0xe3, 0x35,
		0x3d, 0x9e, 0xe9, 0xd7, 0xe4, 0x36, 0x1a, 0x2b, 0xe9, 0x74, 0x35, 0x59, 0x1a, 0x8c, 0xa7, 0xf8,
		0x75, 0xb9, 0x9a, 0x4c, 0x9f, 0x9a, 0x31, 0x5e, 0x24, 0xe3, 0x39, 0xbe, 0x20, 0xcd, 0x18, 0xab,
		0x91, 0xd5, 0x16, 0xa0, 0xc9, 0x02, 0x19, 0xcf, 0xf7, 0x45, 0xc1, 0xb7, 0x38, 0x51, 0x1f, 0xab,
		0x2f, 0xc1, 0xca, 0xf4, 0xe2, 0x18, 0xcf, 0xfa, 0xa5, 0xf7, 0xc7, 0x8e, 0x33, 0xe1, 0xda, 0x58,
		0x3d, 0x1c, 0x65, 0xd9, 0x70, 0x61, 0x8c, 0xa7, 0x7d, 0xf3, 0xfd, 0x68, 0xa2, 0x0d, 0xd7, 0xc5,
		0x6a, 0x0d, 0x60, 0x54, 0x93, 0xe2, 0xb9, 0xde, 0x12, 0x5c, 0x21, 0x10, 0xdd, 0x1a, 0xa2, 0x24,
		0xc5, 0xe3, 0xbf, 0x2c, 0xb7, 0x86, 0x40, 0xd0, 0xad, 0x21, 0xab, 0x51, 0x3c, 0xfa, 0x6d, 0xb9,
		0x35, 0x24, 0xa4, 0x7a, 0x03, 0x72, 0xd6, 0xd0, 0x34, 0x69, 0x6c, 0xa1, 0xfb, 0xff, 0x24, 0xab,
		0xfc, 0xef, 0x1f, 0x08, 0xb0, 0x04, 0x54, 0xaf, 0x42, 0x86, 0x0c, 0x8e, 0x49, 0x27, 0x0e, 0xf9,
		0x1f, 0x1f, 0xc8, 0x7c, 0x42, 0xb5, 0xab, 0x2f, 0x00, 0xf0, 0xc3, 0x34, 0xfb, 0x50, 0x14, 0x83,
		0xfd, 0xcf, 0x0f, 0xc4, 0x8f, 0x25, 0x46, 0x90, 0x11, 0x01, 0xff, 0xe9, 0xc5, 0xfd, 0x09, 0xde,
		0x8b, 0x12, 0xb0, 0x03, 0xf8, 0x75, 0x98, 0xbb, 0xe5, 0xd9, 0x96, 0x8f, 0x7b, 0x71, 0xe8, 0xff,
		0x12, 0x68, 0xa9, 0x4f, 0x1d, 0x36, 0xb0, 0x5d, 0xe2, 0xe3, 0x9e, 0x17, 0x87, 0xfd, 0x6f, 0x81,
		0x0d, 0x00, 0x14, 0xac, 0x63, 0xcf, 0x9f, 0xe5, 0xbd, 0xbf, 0x2b, 0xc1, 0x12, 0x40, 0x8d, 0xa6,
		0xff, 0xdf, 0x26, 0x27, 0x71, 0xd8, 0xef, 0x49, 0xa3, 0x85, 0x7e, 0xf5, 0x93, 0x90, 0xa7, 0xff,
		0xf2, 0x5f, 0x40, 0xc5, 0x80, 0xff, 0x47, 0x80, 0x47, 0x08, 0xfa, 0x64, 0xcf, 0xef, 0xf8, 0x46,
		0xbc, 0xb3, 0xff, 0x57, 0xac, 0xb4, 0xd4, 0xaf, 0xd6, 0xa0, 0xe0, 0xf9, 0x9d, 0xce, 0x50, 0x74,
		0x34, 0x31, 0xf0, 0xff, 0xfb, 0x20, 0x38, 0xe4, 0x06, 0x98, 0xcd, 0xfa, 0xf4, 0xfb, 0x3a, 0xd8,
		0xb5, 0x77, 0x6d, 0x7e, 0x53, 0xf7, 0x6a, 0x25, 0xfe, 0xca, 0x0d, 0xbe, 0x93, 0x84, 0xfc, 0xc0,
		0xeb, 0x89, 0x5b, 0x37, 0xbe, 0xff, 0x68, 0x7d, 0xf1, 0x56, 0xcf, 0x76, 0x61, 0x57, 0xd9, 0x81,
		0xec, 0xa6, 0xd1, 0xdb, 0xf7, 0x7a, 0x68, 0x19, 0x32, 0xcc, 0x7c, 0xf6, 0xb5, 0x29, 0xa5, 0xf2,
		0x01, 0x7a, 0x02, 0x52, 0xfb, 0x5e, 0x4f, 0xfc, 0xde, 0x68, 0x79, 0x63, 0xf4, 0xa0, 0x8d, 0xf6,
		0x00, 0x9b, 0xe6, 0xbe, 0xd7, 0x53, 0xa9, 0x42, 0xc5, 0x85, 0x9c, 0x14, 0xa0, 0x35, 0x28, 0xb4,
		0x75, 0xec, 0x6e, 0x0e, 0xbd, 0xb6, 0x6f, 0x3b, 0xf2, 0xf7, 0x34, 0x21, 0x11, 0x5a, 0x87, 0x85,
		0x1d, 0xd3, 0xe8, 0xf5, 0xfd, 0x16, 0x76, 0xb1, 0xde, 0x1f, 0xfa, 0xa4, 0x5c, 0x5c, 0x4b, 0xad,
		0xcf, 0xa9, 0xe3, 0x62, 0xb4, 0x0a, 0xb9, 0x7d, 0xec, 0xb4, 0xfb, 0xd8, 0xbd, 0xcd, 0x7e, 0x9d,
		0x91, 0x57, 0x83, 0x71, 0xe5, 0x45, 0xc8, 0xb6, 0xf8, 0x87, 0xde, 0x15, 0x48, 0x37, 0xf8, 0xad,
		0x71, 0x6a, 0x3d, 0xc5, 0xaf, 0x59, 0xe9, 0x18, 0xad, 0x42, 0x76, 0xc7, 0xb4, 0xb1, 0xef, 0xb1,
		0x5f, 0xa0, 0x25, 0xd8, 0x8c, 0x90, 0xa0, 0x32, 0x64, 0x8e, 0x0c, 0x79, 0x6b, 0x3c, 0xcf, 0xa6,
		0xb8, 0x60, 0xb3, 0xf8, 0xbd, 0x6f, 0x9f, 0x4f, 0x7c, 0xff, 0xdb, 0xe7, 0x13, 0x5f, 0xfd, 0x97,
		0xf3, 0x89, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x97, 0xfa, 0x77, 0xf7, 0x31, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *BigMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.BigMsg{")
	if this.Field != nil {
		s = append(s, "Field: "+valueToGoStringMsg(this.Field, "int64")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.SmallMsg{")
	if this.ScarBusStop != nil {
		s = append(s, "ScarBusStop: "+valueToGoStringMsg(this.ScarBusStop, "string")+",\n")
	}
	if this.FlightParachute != nil {
		s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	}
	if this.MapShark != nil {
		s = append(s, "MapShark: "+valueToGoStringMsg(this.MapShark, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Packed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.Packed{")
	if this.Ints != nil {
		s = append(s, "Ints: "+fmt.Sprintf("%#v", this.Ints)+",\n")
	}
	if this.Floats != nil {
		s = append(s, "Floats: "+fmt.Sprintf("%#v", this.Floats)+",\n")
	}
	if this.Uints != nil {
		s = append(s, "Uints: "+fmt.Sprintf("%#v", this.Uints)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedBigMsg(r randyMsg, easy bool) *BigMsg {
	this := &BigMsg{}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field = &v1
	}
	if r.Intn(10) != 0 {
		this.Msg = NewPopulatedSmallMsg(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 4)
	}
	return this
}

func NewPopulatedSmallMsg(r randyMsg, easy bool) *SmallMsg {
	this := &SmallMsg{}
	if r.Intn(10) != 0 {
		v2 := string(randStringMsg(r))
		this.ScarBusStop = &v2
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.FlightParachute = make([]uint32, v3)
		for i := 0; i < v3; i++ {
			this.FlightParachute[i] = uint32(r.Uint32())
		}
	}
	if r.Intn(10) != 0 {
		v4 := string(randStringMsg(r))
		this.MapShark = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 19)
	}
	return this
}

func NewPopulatedPacked(r randyMsg, easy bool) *Packed {
	this := &Packed{}
	if r.Intn(10) != 0 {
		v5 := r.Intn(10)
		this.Ints = make([]int64, v5)
		for i := 0; i < v5; i++ {
			this.Ints[i] = int64(r.Int63())
			if r.Intn(2) == 0 {
				this.Ints[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(10)
		this.Floats = make([]float64, v6)
		for i := 0; i < v6; i++ {
			this.Floats[i] = float64(r.Float64())
			if r.Intn(2) == 0 {
				this.Floats[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v7 := r.Intn(10)
		this.Uints = make([]uint32, v7)
		for i := 0; i < v7; i++ {
			this.Uints[i] = uint32(r.Uint32())
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 7)
	}
	return this
}

type randyMsg interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMsg(r randyMsg) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMsg(r randyMsg) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneMsg(r)
	}
	return string(tmps)
}
func randUnrecognizedMsg(r randyMsg, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMsg(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMsg(dAtA []byte, r randyMsg, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMsg(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("msg.proto", fileDescriptorMsg) }

var fileDescriptorMsg = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8d, 0x4f, 0x6b, 0x83, 0x30,
	0x18, 0xc6, 0xc9, 0x52, 0x5d, 0xfb, 0xb6, 0x63, 0x23, 0x94, 0x11, 0x3c, 0x48, 0xf0, 0x30, 0x72,
	0x99, 0x85, 0x7d, 0x04, 0x0f, 0xc2, 0x0e, 0x42, 0x51, 0xb6, 0x7b, 0x6a, 0x5d, 0x94, 0xea, 0x22,
	0x26, 0x7e, 0xaf, 0x7d, 0xa5, 0xee, 0x13, 0xec, 0xb8, 0xe3, 0xe8, 0xeb, 0xfe, 0xd1, 0x53, 0xf2,
	0xfb, 0x3d, 0x79, 0x9e, 0xc0, 0xa2, 0xb3, 0x3a, 0xee, 0x07, 0xe3, 0x0c, 0x03, 0x3c, 0x5c, 0x65,
	0x9d, 0x0d, 0xee, 0x75, 0xe3, 0xea, 0x71, 0x17, 0x97, 0xa6, 0xdb, 0x68, 0xa3, 0xcd, 0x06, 0xb3,
	0xdd, 0xf8, 0x82, 0x84, 0x80, 0xb7, 0xa9, 0x1a, 0xa5, 0xe0, 0x27, 0x8d, 0xce, 0xac, 0x66, 0x6b,
	0xf0, 0xd2, 0xa6, 0x6a, 0xf7, 0x9c, 0x08, 0x22, 0x69, 0x3e, 0x01, 0xbb, 0x03, 0x9a, 0x59, 0xcd,
	0xa9, 0x20, 0x72, 0xf9, 0xb0, 0x8e, 0xff, 0x3e, 0x8a, 0x8b, 0x4e, 0xb5, 0x6d, 0x66, 0x75, 0x7e,
	0x7a, 0x10, 0x0d, 0x30, 0xff, 0x11, 0x4c, 0xc0, 0xb2, 0x28, 0xd5, 0x90, 0x8c, 0xb6, 0x70, 0xa6,
	0xc7, 0xbd, 0x45, 0xfe, 0x5f, 0x31, 0x09, 0xd7, 0x69, 0xdb, 0xe8, 0xda, 0x6d, 0xd5, 0xa0, 0xca,
	0x7a, 0x74, 0x15, 0x5f, 0x09, 0x2a, 0x2f, 0xf3, 0x73, 0xcd, 0x02, 0x98, 0x67, 0xaa, 0x2f, 0x6a,
	0x35, 0x1c, 0x38, 0xc3, 0xa1, 0x5f, 0x8e, 0x9e, 0xc1, 0xdf, 0xaa, 0xf2, 0x50, 0xed, 0xd9, 0x2d,
	0xcc, 0x1e, 0x5f, 0x9d, 0xe5, 0x33, 0x41, 0x25, 0x4d, 0x2e, 0x6e, 0x48, 0x8e, 0xcc, 0x02, 0xf0,
	0xd3, 0xd6, 0x28, 0x67, 0xb9, 0x27, 0xa8, 0x24, 0x98, 0x7c, 0x1b, 0xc6, 0xc1, 0x7b, 0x6a, 0x4e,
	0x25, 0x5f, 0x50, 0x79, 0x85, 0xd1, 0x24, 0x92, 0xd5, 0xc7, 0x31, 0x24, 0x9f, 0xc7, 0x90, 0xbc,
	0xbd, 0x87, 0xe4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xcd, 0x1a, 0x2c, 0x68, 0x01, 0x00, 0x00,
}
