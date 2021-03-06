// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protonum.proto

/*
Package protonum is a generated protocol buffer package.

It is generated from these files:
	protonum.proto

It has these top-level messages:
	ProtoNum
	KeyValue
	TopsyTurvy
	Topsy
	Turvy
	Knot
	BightKnot
*/
package protonum

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

// ProtoNum is used for testing of the protonum package.
type ProtoNum struct {
	KeyValue         []*KeyValue `protobuf:"bytes,2,rep,name=KeyValue" json:"KeyValue,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ProtoNum) Reset()                    { *m = ProtoNum{} }
func (m *ProtoNum) String() string            { return proto.CompactTextString(m) }
func (*ProtoNum) ProtoMessage()               {}
func (*ProtoNum) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{0} }

func (m *ProtoNum) GetKeyValue() []*KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

// KeyValue is used for testing of the protonum package.
type KeyValue struct {
	Key              *string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{1} }

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// TopsyTurvy is used for testing of the protonum package.
type TopsyTurvy struct {
	Topsy            *Topsy `protobuf:"bytes,1,opt,name=Topsy" json:"Topsy,omitempty"`
	Turvy            *Turvy `protobuf:"bytes,2,opt,name=Turvy" json:"Turvy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TopsyTurvy) Reset()                    { *m = TopsyTurvy{} }
func (m *TopsyTurvy) String() string            { return proto.CompactTextString(m) }
func (*TopsyTurvy) ProtoMessage()               {}
func (*TopsyTurvy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{2} }

func (m *TopsyTurvy) GetTopsy() *Topsy {
	if m != nil {
		return m.Topsy
	}
	return nil
}

func (m *TopsyTurvy) GetTurvy() *Turvy {
	if m != nil {
		return m.Turvy
	}
	return nil
}

// Topsy is used for testing of the protonum package.
type Topsy struct {
	A                *int64 `protobuf:"varint,1,opt,name=A" json:"A,omitempty"`
	B                *int64 `protobuf:"varint,2,opt,name=B" json:"B,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Topsy) Reset()                    { *m = Topsy{} }
func (m *Topsy) String() string            { return proto.CompactTextString(m) }
func (*Topsy) ProtoMessage()               {}
func (*Topsy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{3} }

func (m *Topsy) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *Topsy) GetB() int64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

// Turvy is used for testing of the protonum package.
type Turvy struct {
	B                *int64 `protobuf:"varint,1,opt,name=B" json:"B,omitempty"`
	A                *int64 `protobuf:"varint,2,opt,name=A" json:"A,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Turvy) Reset()                    { *m = Turvy{} }
func (m *Turvy) String() string            { return proto.CompactTextString(m) }
func (*Turvy) ProtoMessage()               {}
func (*Turvy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{4} }

func (m *Turvy) GetB() int64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *Turvy) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

// Knot is used for testing of the protonum package.
type Knot struct {
	Bight            []*BightKnot `protobuf:"bytes,1,rep,name=Bight" json:"Bight,omitempty"`
	Elbow            *bool        `protobuf:"varint,2,opt,name=Elbow" json:"Elbow,omitempty"`
	BitterEnd        *Knot        `protobuf:"bytes,3,opt,name=BitterEnd" json:"BitterEnd,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Knot) Reset()                    { *m = Knot{} }
func (m *Knot) String() string            { return proto.CompactTextString(m) }
func (*Knot) ProtoMessage()               {}
func (*Knot) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{5} }

func (m *Knot) GetBight() []*BightKnot {
	if m != nil {
		return m.Bight
	}
	return nil
}

func (m *Knot) GetElbow() bool {
	if m != nil && m.Elbow != nil {
		return *m.Elbow
	}
	return false
}

func (m *Knot) GetBitterEnd() *Knot {
	if m != nil {
		return m.BitterEnd
	}
	return nil
}

// BightKnot is used for testing of the protonum package.
type BightKnot struct {
	Loop             *Knot  `protobuf:"bytes,1,opt,name=Loop" json:"Loop,omitempty"`
	Turn             *bool  `protobuf:"varint,2,opt,name=Turn" json:"Turn,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BightKnot) Reset()                    { *m = BightKnot{} }
func (m *BightKnot) String() string            { return proto.CompactTextString(m) }
func (*BightKnot) ProtoMessage()               {}
func (*BightKnot) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{6} }

func (m *BightKnot) GetLoop() *Knot {
	if m != nil {
		return m.Loop
	}
	return nil
}

func (m *BightKnot) GetTurn() bool {
	if m != nil && m.Turn != nil {
		return *m.Turn
	}
	return false
}

func init() {
	proto.RegisterType((*ProtoNum)(nil), "protonum.ProtoNum")
	proto.RegisterType((*KeyValue)(nil), "protonum.KeyValue")
	proto.RegisterType((*TopsyTurvy)(nil), "protonum.TopsyTurvy")
	proto.RegisterType((*Topsy)(nil), "protonum.Topsy")
	proto.RegisterType((*Turvy)(nil), "protonum.Turvy")
	proto.RegisterType((*Knot)(nil), "protonum.Knot")
	proto.RegisterType((*BightKnot)(nil), "protonum.BightKnot")
}
func (this *ProtoNum) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *KeyValue) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *TopsyTurvy) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Topsy) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Turvy) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Knot) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *BightKnot) Description() (desc *descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func ProtonumDescription() (desc *descriptor.FileDescriptorSet) {
	d := &descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3869 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5b, 0x70, 0x1b, 0xd7,
		0x79, 0x36, 0x6e, 0x24, 0xf0, 0x03, 0x04, 0x97, 0x87, 0xb4, 0x04, 0xd1, 0xb1, 0x25, 0xc1, 0x76,
		0x44, 0xd9, 0x09, 0x95, 0x91, 0x2d, 0xd9, 0x5a, 0x35, 0x71, 0x01, 0x70, 0xc5, 0x40, 0x26, 0x09,
		0x64, 0x01, 0xfa, 0x96, 0xe9, 0xec, 0x2c, 0x17, 0x07, 0xe0, 0x4a, 0x8b, 0xdd, 0xcd, 0xee, 0x42,
		0x32, 0x35, 0x7d, 0x70, 0xc7, 0xbd, 0x4c, 0xa6, 0xd3, 0x7b, 0x67, 0x92, 0xb8, 0x8e, 0x7b, 0x99,
		0x49, 0xdd, 0xa6, 0xb7, 0xa4, 0x69, 0xd3, 0xcb, 0x53, 0x5f, 0xd2, 0xf6, 0xa9, 0x33, 0x79, 0xef,
		0x43, 0x2f, 0x9e, 0xe9, 0xcd, 0x6d, 0xd2, 0xd6, 0x0f, 0x9d, 0xd1, 0x4b, 0xe7, 0xdc, 0x16, 0xbb,
		0x00, 0xa8, 0x05, 0x33, 0x63, 0xe7, 0x89, 0x3c, 0xff, 0xf9, 0xbf, 0xef, 0x9c, 0xf3, 0x9f, 0xff,
		0xfc, 0xff, 0x7f, 0xce, 0x02, 0xbe, 0x77, 0x0d, 0xce, 0x0d, 0x1c, 0x67, 0x60, 0xe1, 0x4b, 0xae,
		0xe7, 0x04, 0xce, 0xc1, 0xa8, 0x7f, 0xa9, 0x87, 0x7d, 0xc3, 0x33, 0xdd, 0xc0, 0xf1, 0x36, 0xa9,
		0x0c, 0x2d, 0x33, 0x8d, 0x4d, 0xa1, 0x51, 0xdd, 0x85, 0x95, 0x1b, 0xa6, 0x85, 0xb7, 0x42, 0xc5,
		0x0e, 0x0e, 0xd0, 0xf3, 0x90, 0xed, 0x9b, 0x16, 0xae, 0xa4, 0xce, 0x65, 0x36, 0x8a, 0x97, 0x9f,
		0xd8, 0x9c, 0x00, 0x6d, 0xc6, 0x11, 0x6d, 0x22, 0x56, 0x29, 0xa2, 0xfa, 0x5e, 0x16, 0x56, 0x67,
		0xf4, 0x22, 0x04, 0x59, 0x5b, 0x1f, 0x12, 0xc6, 0xd4, 0x46, 0x41, 0xa5, 0xff, 0xa3, 0x0a, 0x2c,
		0xba, 0xba, 0x71, 0x5b, 0x1f, 0xe0, 0x4a, 0x9a, 0x8a, 0x45, 0x13, 0x3d, 0x06, 0xd0, 0xc3, 0x2e,
		0xb6, 0x7b, 0xd8, 0x36, 0x8e, 0x2a, 0x99, 0x73, 0x99, 0x8d, 0x82, 0x1a, 0x91, 0xa0, 0xa7, 0x61,
		0xc5, 0x1d, 0x1d, 0x58, 0xa6, 0xa1, 0x45, 0xd4, 0xe0, 0x5c, 0x66, 0x23, 0xa7, 0x4a, 0xac, 0x63,
		0x6b, 0xac, 0x7c, 0x01, 0x96, 0xef, 0x62, 0xfd, 0x76, 0x54, 0xb5, 0x48, 0x55, 0xcb, 0x44, 0x1c,
		0x51, 0x6c, 0x40, 0x69, 0x88, 0x7d, 0x5f, 0x1f, 0x60, 0x2d, 0x38, 0x72, 0x71, 0x25, 0x4b, 0x57,
		0x7f, 0x6e, 0x6a, 0xf5, 0x93, 0x2b, 0x2f, 0x72, 0x54, 0xf7, 0xc8, 0xc5, 0xa8, 0x06, 0x05, 0x6c,
		0x8f, 0x86, 0x8c, 0x21, 0x77, 0x8c, 0xfd, 0x14, 0x7b, 0x34, 0x9c, 0x64, 0xc9, 0x13, 0x18, 0xa7,
		0x58, 0xf4, 0xb1, 0x77, 0xc7, 0x34, 0x70, 0x65, 0x81, 0x12, 0x5c, 0x98, 0x22, 0xe8, 0xb0, 0xfe,
		0x49, 0x0e, 0x81, 0x43, 0x0d, 0x28, 0xe0, 0xd7, 0x03, 0x6c, 0xfb, 0xa6, 0x63, 0x57, 0x16, 0x29,
		0xc9, 0x93, 0x33, 0x76, 0x11, 0x5b, 0xbd, 0x49, 0x8a, 0x31, 0x0e, 0x5d, 0x85, 0x45, 0xc7, 0x0d,
		0x4c, 0xc7, 0xf6, 0x2b, 0xf9, 0x73, 0xa9, 0x8d, 0xe2, 0xe5, 0x8f, 0xcd, 0x74, 0x84, 0x16, 0xd3,
		0x51, 0x85, 0x32, 0x6a, 0x82, 0xe4, 0x3b, 0x23, 0xcf, 0xc0, 0x9a, 0xe1, 0xf4, 0xb0, 0x66, 0xda,
		0x7d, 0xa7, 0x52, 0xa0, 0x04, 0x67, 0xa7, 0x17, 0x42, 0x15, 0x1b, 0x4e, 0x0f, 0x37, 0xed, 0xbe,
		0xa3, 0x96, 0xfd, 0x58, 0x1b, 0x9d, 0x82, 0x05, 0xff, 0xc8, 0x0e, 0xf4, 0xd7, 0x2b, 0x25, 0xea,
		0x21, 0xbc, 0x55, 0xfd, 0x8b, 0x05, 0x58, 0x9e, 0xc7, 0xc5, 0xae, 0x43, 0xae, 0x4f, 0x56, 0x59,
		0x49, 0x9f, 0xc4, 0x06, 0x0c, 0x13, 0x37, 0xe2, 0xc2, 0x0f, 0x68, 0xc4, 0x1a, 0x14, 0x6d, 0xec,
		0x07, 0xb8, 0xc7, 0x3c, 0x22, 0x33, 0xa7, 0x4f, 0x01, 0x03, 0x4d, 0xbb, 0x54, 0xf6, 0x07, 0x72,
		0xa9, 0x57, 0x60, 0x39, 0x9c, 0x92, 0xe6, 0xe9, 0xf6, 0x40, 0xf8, 0xe6, 0xa5, 0xa4, 0x99, 0x6c,
		0x2a, 0x02, 0xa7, 0x12, 0x98, 0x5a, 0xc6, 0xb1, 0x36, 0xda, 0x02, 0x70, 0x6c, 0xec, 0xf4, 0xb5,
		0x1e, 0x36, 0xac, 0x4a, 0xfe, 0x18, 0x2b, 0xb5, 0x88, 0xca, 0x94, 0x95, 0x1c, 0x26, 0x35, 0x2c,
		0x74, 0x6d, 0xec, 0x6a, 0x8b, 0xc7, 0x78, 0xca, 0x2e, 0x3b, 0x64, 0x53, 0xde, 0xb6, 0x0f, 0x65,
		0x0f, 0x13, 0xbf, 0xc7, 0x3d, 0xbe, 0xb2, 0x02, 0x9d, 0xc4, 0x66, 0xe2, 0xca, 0x54, 0x0e, 0x63,
		0x0b, 0x5b, 0xf2, 0xa2, 0x4d, 0xf4, 0x38, 0x84, 0x02, 0x8d, 0xba, 0x15, 0xd0, 0x28, 0x54, 0x12,
		0xc2, 0x3d, 0x7d, 0x88, 0xd7, 0xef, 0x41, 0x39, 0x6e, 0x1e, 0xb4, 0x06, 0x39, 0x3f, 0xd0, 0xbd,
		0x80, 0x7a, 0x61, 0x4e, 0x65, 0x0d, 0x24, 0x41, 0x06, 0xdb, 0x3d, 0x1a, 0xe5, 0x72, 0x2a, 0xf9,
		0x17, 0xfd, 0xe8, 0x78, 0xc1, 0x19, 0xba, 0xe0, 0x8f, 0x4f, 0xef, 0x68, 0x8c, 0x79, 0x72, 0xdd,
		0xeb, 0xcf, 0xc1, 0x52, 0x6c, 0x01, 0xf3, 0x0e, 0x5d, 0xfd, 0x71, 0x78, 0x78, 0x26, 0x35, 0x7a,
		0x05, 0xd6, 0x46, 0xb6, 0x69, 0x07, 0xd8, 0x73, 0x3d, 0x4c, 0x3c, 0x96, 0x0d, 0x55, 0xf9, 0x97,
		0xc5, 0x63, 0x7c, 0x6e, 0x3f, 0xaa, 0xcd, 0x58, 0xd4, 0xd5, 0xd1, 0xb4, 0xf0, 0xa9, 0x42, 0xfe,
		0x5f, 0x17, 0xa5, 0x37, 0xde, 0x78, 0xe3, 0x8d, 0x74, 0xf5, 0xcb, 0x0b, 0xb0, 0x36, 0xeb, 0xcc,
		0xcc, 0x3c, 0xbe, 0xa7, 0x60, 0xc1, 0x1e, 0x0d, 0x0f, 0xb0, 0x47, 0x8d, 0x94, 0x53, 0x79, 0x0b,
		0xd5, 0x20, 0x67, 0xe9, 0x07, 0xd8, 0xaa, 0x64, 0xcf, 0xa5, 0x36, 0xca, 0x97, 0x9f, 0x9e, 0xeb,
		0x54, 0x6e, 0xee, 0x10, 0x88, 0xca, 0x90, 0xe8, 0x33, 0x90, 0xe5, 0x21, 0x9a, 0x30, 0x3c, 0x35,
		0x1f, 0x03, 0x39, 0x4b, 0x2a, 0xc5, 0xa1, 0x47, 0xa0, 0x40, 0xfe, 0x32, 0xdf, 0x58, 0xa0, 0x73,
		0xce, 0x13, 0x01, 0xf1, 0x0b, 0xb4, 0x0e, 0x79, 0x7a, 0x4c, 0x7a, 0x58, 0xa4, 0xb6, 0xb0, 0x4d,
		0x1c, 0xab, 0x87, 0xfb, 0xfa, 0xc8, 0x0a, 0xb4, 0x3b, 0xba, 0x35, 0xc2, 0xd4, 0xe1, 0x0b, 0x6a,
		0x89, 0x0b, 0x5f, 0x22, 0x32, 0x74, 0x16, 0x8a, 0xec, 0x54, 0x99, 0x76, 0x0f, 0xbf, 0x4e, 0xa3,
		0x67, 0x4e, 0x65, 0x07, 0xad, 0x49, 0x24, 0x64, 0xf8, 0x5b, 0xbe, 0x63, 0x0b, 0xd7, 0xa4, 0x43,
		0x10, 0x01, 0x1d, 0xfe, 0xb9, 0xc9, 0xc0, 0xfd, 0xe8, 0xec, 0xe5, 0x4d, 0xfa, 0x54, 0xf5, 0xdb,
		0x69, 0xc8, 0xd2, 0x78, 0xb1, 0x0c, 0xc5, 0xee, 0xab, 0x6d, 0x45, 0xdb, 0x6a, 0xed, 0xd7, 0x77,
		0x14, 0x29, 0x85, 0xca, 0x00, 0x54, 0x70, 0x63, 0xa7, 0x55, 0xeb, 0x4a, 0xe9, 0xb0, 0xdd, 0xdc,
		0xeb, 0x5e, 0x7d, 0x56, 0xca, 0x84, 0x80, 0x7d, 0x26, 0xc8, 0x46, 0x15, 0x9e, 0xb9, 0x2c, 0xe5,
		0x90, 0x04, 0x25, 0x46, 0xd0, 0x7c, 0x45, 0xd9, 0xba, 0xfa, 0xac, 0xb4, 0x10, 0x97, 0x3c, 0x73,
		0x59, 0x5a, 0x44, 0x4b, 0x50, 0xa0, 0x92, 0x7a, 0xab, 0xb5, 0x23, 0xe5, 0x43, 0xce, 0x4e, 0x57,
		0x6d, 0xee, 0x6d, 0x4b, 0x85, 0x90, 0x73, 0x5b, 0x6d, 0xed, 0xb7, 0x25, 0x08, 0x19, 0x76, 0x95,
		0x4e, 0xa7, 0xb6, 0xad, 0x48, 0xc5, 0x50, 0xa3, 0xfe, 0x6a, 0x57, 0xe9, 0x48, 0xa5, 0xd8, 0xb4,
		0x9e, 0xb9, 0x2c, 0x2d, 0x85, 0x43, 0x28, 0x7b, 0xfb, 0xbb, 0x52, 0x19, 0xad, 0xc0, 0x12, 0x1b,
		0x42, 0x4c, 0x62, 0x79, 0x42, 0x74, 0xf5, 0x59, 0x49, 0x1a, 0x4f, 0x84, 0xb1, 0xac, 0xc4, 0x04,
		0x57, 0x9f, 0x95, 0x50, 0xb5, 0x01, 0x39, 0xea, 0x5d, 0x08, 0x41, 0x79, 0xa7, 0x56, 0x57, 0x76,
		0xb4, 0x56, 0xbb, 0xdb, 0x6c, 0xed, 0xd5, 0x76, 0xa4, 0xd4, 0x58, 0xa6, 0x2a, 0x9f, 0xdb, 0x6f,
		0xaa, 0xca, 0x96, 0x94, 0x8e, 0xca, 0xda, 0x4a, 0xad, 0xab, 0x6c, 0x49, 0x99, 0xaa, 0x01, 0x6b,
		0xb3, 0xe2, 0xe4, 0xcc, 0x93, 0x11, 0xd9, 0xe2, 0xf4, 0x31, 0x5b, 0x4c, 0xb9, 0xa6, 0xb6, 0xf8,
		0x9f, 0xd3, 0xb0, 0x3a, 0x23, 0x57, 0xcc, 0x1c, 0xe4, 0x05, 0xc8, 0x31, 0x17, 0x65, 0xd9, 0xf3,
		0xe2, 0xcc, 0xa4, 0x43, 0x1d, 0x76, 0x2a, 0x83, 0x52, 0x5c, 0xb4, 0x82, 0xc8, 0x1c, 0x53, 0x41,
		0x10, 0x8a, 0xa9, 0x98, 0xfe, 0x63, 0x53, 0x31, 0x9d, 0xa5, 0xbd, 0xab, 0xf3, 0xa4, 0x3d, 0x2a,
		0x3b, 0x59, 0x6c, 0xcf, 0xcd, 0x88, 0xed, 0xd7, 0x61, 0x65, 0x8a, 0x68, 0xee, 0x18, 0xfb, 0x66,
		0x0a, 0x2a, 0xc7, 0x19, 0x27, 0x21, 0xd2, 0xa5, 0x63, 0x91, 0xee, 0xfa, 0xa4, 0x05, 0xcf, 0x1f,
		0xbf, 0x09, 0x53, 0x7b, 0xfd, 0x6e, 0x0a, 0x4e, 0xcd, 0xae, 0x14, 0x67, 0xce, 0xe1, 0x33, 0xb0,
		0x30, 0xc4, 0xc1, 0xa1, 0x23, 0xaa, 0xa5, 0x8f, 0xcf, 0xc8, 0xc1, 0xa4, 0x7b, 0x72, 0xb3, 0x39,
		0x2a, 0x9a, 0xc4, 0x33, 0xc7, 0x95, 0x7b, 0x6c, 0x36, 0x53, 0x33, 0xfd, 0x62, 0x1a, 0x1e, 0x9e,
		0x49, 0x3e, 0x73, 0xa2, 0x8f, 0x02, 0x98, 0xb6, 0x3b, 0x0a, 0x58, 0x45, 0xc4, 0x02, 0x6c, 0x81,
		0x4a, 0x68, 0xf0, 0x22, 0xc1, 0x73, 0x14, 0x84, 0xfd, 0x19, 0xda, 0x0f, 0x4c, 0x44, 0x15, 0x9e,
		0x1f, 0x4f, 0x34, 0x4b, 0x27, 0xfa, 0xd8, 0x31, 0x2b, 0x9d, 0x72, 0xcc, 0x4f, 0x81, 0x64, 0x58,
		0x26, 0xb6, 0x03, 0xcd, 0x0f, 0x3c, 0xac, 0x0f, 0x4d, 0x7b, 0x40, 0x33, 0x48, 0x5e, 0xce, 0xf5,
		0x75, 0xcb, 0xc7, 0xea, 0x32, 0xeb, 0xee, 0x88, 0x5e, 0x82, 0xa0, 0x0e, 0xe4, 0x45, 0x10, 0x0b,
		0x31, 0x04, 0xeb, 0x0e, 0x11, 0xd5, 0x6f, 0xe5, 0xa1, 0x18, 0xa9, 0xab, 0xd1, 0x79, 0x28, 0xdd,
		0xd2, 0xef, 0xe8, 0x9a, 0xb8, 0x2b, 0x31, 0x4b, 0x14, 0x89, 0xac, 0xcd, 0xef, 0x4b, 0x9f, 0x82,
		0x35, 0xaa, 0xe2, 0x8c, 0x02, 0xec, 0x69, 0x86, 0xa5, 0xfb, 0x3e, 0x35, 0x5a, 0x9e, 0xaa, 0x22,
		0xd2, 0xd7, 0x22, 0x5d, 0x0d, 0xd1, 0x83, 0xae, 0xc0, 0x2a, 0x45, 0x0c, 0x47, 0x56, 0x60, 0xba,
		0x16, 0xd6, 0xc8, 0xed, 0xcd, 0xa7, 0x99, 0x24, 0x9c, 0xd9, 0x0a, 0xd1, 0xd8, 0xe5, 0x0a, 0x64,
		0x46, 0x3e, 0xda, 0x82, 0x47, 0x29, 0x6c, 0x80, 0x6d, 0xec, 0xe9, 0x01, 0xd6, 0xf0, 0x17, 0x46,
		0xba, 0xe5, 0x6b, 0xba, 0xdd, 0xd3, 0x0e, 0x75, 0xff, 0xb0, 0xb2, 0x46, 0x08, 0xea, 0xe9, 0x4a,
		0x4a, 0x3d, 0x43, 0x14, 0xb7, 0xb9, 0x9e, 0x42, 0xd5, 0x6a, 0x76, 0xef, 0xb3, 0xba, 0x7f, 0x88,
		0x64, 0x38, 0x45, 0x59, 0xfc, 0xc0, 0x33, 0xed, 0x81, 0x66, 0x1c, 0x62, 0xe3, 0xb6, 0x36, 0x0a,
		0xfa, 0xcf, 0x57, 0x1e, 0x89, 0x8e, 0x4f, 0x67, 0xd8, 0xa1, 0x3a, 0x0d, 0xa2, 0xb2, 0x1f, 0xf4,
		0x9f, 0x47, 0x1d, 0x28, 0x91, 0xcd, 0x18, 0x9a, 0xf7, 0xb0, 0xd6, 0x77, 0x3c, 0x9a, 0x1a, 0xcb,
		0x33, 0x42, 0x53, 0xc4, 0x82, 0x9b, 0x2d, 0x0e, 0xd8, 0x75, 0x7a, 0x58, 0xce, 0x75, 0xda, 0x8a,
		0xb2, 0xa5, 0x16, 0x05, 0xcb, 0x0d, 0xc7, 0x23, 0x0e, 0x35, 0x70, 0x42, 0x03, 0x17, 0x99, 0x43,
		0x0d, 0x1c, 0x61, 0xde, 0x2b, 0xb0, 0x6a, 0x18, 0x6c, 0xcd, 0xa6, 0xa1, 0xf1, 0x3b, 0x96, 0x5f,
		0x91, 0x62, 0xc6, 0x32, 0x8c, 0x6d, 0xa6, 0xc0, 0x7d, 0xdc, 0x47, 0xd7, 0xe0, 0xe1, 0xb1, 0xb1,
		0xa2, 0xc0, 0x95, 0xa9, 0x55, 0x4e, 0x42, 0xaf, 0xc0, 0xaa, 0x7b, 0x34, 0x0d, 0x44, 0xb1, 0x11,
		0xdd, 0xa3, 0x49, 0xd8, 0x73, 0xb0, 0xe6, 0x1e, 0xba, 0xd3, 0xb8, 0xa7, 0xa2, 0x38, 0xe4, 0x1e,
		0xba, 0x93, 0xc0, 0x27, 0xe9, 0x85, 0xdb, 0xc3, 0x86, 0x1e, 0xe0, 0x5e, 0xe5, 0x74, 0x54, 0x3d,
		0xd2, 0x81, 0x2e, 0x81, 0x64, 0x18, 0x1a, 0xb6, 0xf5, 0x03, 0x0b, 0x6b, 0xba, 0x87, 0x6d, 0xdd,
		0xaf, 0x9c, 0x8d, 0x2a, 0x97, 0x0d, 0x43, 0xa1, 0xbd, 0x35, 0xda, 0x89, 0x9e, 0x82, 0x15, 0xe7,
		0xe0, 0x96, 0xc1, 0x5c, 0x52, 0x73, 0x3d, 0xdc, 0x37, 0x5f, 0xaf, 0x3c, 0x41, 0xed, 0xbb, 0x4c,
		0x3a, 0xa8, 0x43, 0xb6, 0xa9, 0x18, 0x5d, 0x04, 0xc9, 0xf0, 0x0f, 0x75, 0xcf, 0xa5, 0x31, 0xd9,
		0x77, 0x75, 0x03, 0x57, 0x9e, 0x64, 0xaa, 0x4c, 0xbe, 0x27, 0xc4, 0xe4, 0x48, 0xf8, 0x77, 0xcd,
		0x7e, 0x20, 0x18, 0x2f, 0xb0, 0x23, 0x41, 0x65, 0x9c, 0x6d, 0x03, 0x24, 0x62, 0x8a, 0xd8, 0xc0,
		0x1b, 0x54, 0xad, 0xec, 0x1e, 0xba, 0xd1, 0x71, 0x1f, 0x87, 0x25, 0xa2, 0x39, 0x1e, 0xf4, 0x22,
		0x2b, 0xc8, 0xdc, 0xc3, 0xc8, 0x88, 0x1f, 0x5a, 0x6d, 0x5c, 0x95, 0xa1, 0x14, 0xf5, 0x4f, 0x54,
		0x00, 0xe6, 0xa1, 0x52, 0x8a, 0x14, 0x2b, 0x8d, 0xd6, 0x16, 0x29, 0x33, 0x5e, 0x53, 0xa4, 0x34,
		0x29, 0x77, 0x76, 0x9a, 0x5d, 0x45, 0x53, 0xf7, 0xf7, 0xba, 0xcd, 0x5d, 0x45, 0xca, 0x44, 0xeb,
		0xea, 0xef, 0xa4, 0xa1, 0x1c, 0xbf, 0x22, 0xa1, 0x1f, 0x81, 0xd3, 0xe2, 0x3d, 0xc3, 0xc7, 0x81,
		0x76, 0xd7, 0xf4, 0xe8, 0x91, 0x19, 0xea, 0x2c, 0x7d, 0x85, 0x9b, 0xb6, 0xc6, 0xb5, 0x3a, 0x38,
		0x78, 0xd9, 0xf4, 0xc8, 0x81, 0x18, 0xea, 0x01, 0xda, 0x81, 0xb3, 0xb6, 0xa3, 0xf9, 0x81, 0x6e,
		0xf7, 0x74, 0xaf, 0xa7, 0x8d, 0x5f, 0x92, 0x34, 0xdd, 0x30, 0xb0, 0xef, 0x3b, 0x2c, 0x55, 0x85,
		0x2c, 0x1f, 0xb3, 0x9d, 0x0e, 0x57, 0x1e, 0xc7, 0xf0, 0x1a, 0x57, 0x9d, 0x70, 0xb0, 0xcc, 0x71,
		0x0e, 0xf6, 0x08, 0x14, 0x86, 0xba, 0xab, 0x61, 0x3b, 0xf0, 0x8e, 0x68, 0x61, 0x9c, 0x57, 0xf3,
		0x43, 0xdd, 0x55, 0x48, 0xfb, 0xa3, 0xb9, 0x9f, 0xfc, 0x7d, 0x06, 0x4a, 0xd1, 0xe2, 0x98, 0xdc,
		0x35, 0x0c, 0x9a, 0x47, 0x52, 0x34, 0xd2, 0x3c, 0xfe, 0xc0, 0x52, 0x7a, 0xb3, 0x41, 0x12, 0x8c,
		0xbc, 0xc0, 0x4a, 0x56, 0x95, 0x21, 0x49, 0x72, 0x27, 0xb1, 0x05, 0xb3, 0x12, 0x21, 0xaf, 0xf2,
		0x16, 0xda, 0x86, 0x85, 0x5b, 0x3e, 0xe5, 0x5e, 0xa0, 0xdc, 0x4f, 0x3c, 0x98, 0xfb, 0x66, 0x87,
		0x92, 0x17, 0x6e, 0x76, 0xb4, 0xbd, 0x96, 0xba, 0x5b, 0xdb, 0x51, 0x39, 0x1c, 0x9d, 0x81, 0xac,
		0xa5, 0xdf, 0x3b, 0x8a, 0xa7, 0x22, 0x2a, 0x9a, 0xd7, 0xf0, 0x67, 0x20, 0x7b, 0x17, 0xeb, 0xb7,
		0xe3, 0x09, 0x80, 0x8a, 0x3e, 0x44, 0xd7, 0xbf, 0x04, 0x39, 0x6a, 0x2f, 0x04, 0xc0, 0x2d, 0x26,
		0x3d, 0x84, 0xf2, 0x90, 0x6d, 0xb4, 0x54, 0xe2, 0xfe, 0x12, 0x94, 0x98, 0x54, 0x6b, 0x37, 0x95,
		0x86, 0x22, 0xa5, 0xab, 0x57, 0x60, 0x81, 0x19, 0x81, 0x1c, 0x8d, 0xd0, 0x0c, 0xd2, 0x43, 0xbc,
		0xc9, 0x39, 0x52, 0xa2, 0x77, 0x7f, 0xb7, 0xae, 0xa8, 0x52, 0x3a, 0xba, 0xbd, 0x3e, 0x94, 0xa2,
		0x75, 0xf1, 0x47, 0xe3, 0x53, 0x7f, 0x99, 0x82, 0x62, 0xa4, 0xce, 0x25, 0x05, 0x8a, 0x6e, 0x59,
		0xce, 0x5d, 0x4d, 0xb7, 0x4c, 0xdd, 0xe7, 0x4e, 0x01, 0x54, 0x54, 0x23, 0x92, 0x79, 0x37, 0xed,
		0x23, 0x99, 0xfc, 0x3b, 0x29, 0x90, 0x26, 0x4b, 0xcc, 0x89, 0x09, 0xa6, 0x7e, 0xa8, 0x13, 0x7c,
		0x3b, 0x05, 0xe5, 0x78, 0x5d, 0x39, 0x31, 0xbd, 0xf3, 0x3f, 0xd4, 0xe9, 0xfd, 0x43, 0x1a, 0x96,
		0x62, 0xd5, 0xe4, 0xbc, 0xb3, 0xfb, 0x02, 0xac, 0x98, 0x3d, 0x3c, 0x74, 0x9d, 0x00, 0xdb, 0xc6,
		0x91, 0x66, 0xe1, 0x3b, 0xd8, 0xaa, 0x54, 0x69, 0xa0, 0xb8, 0xf4, 0xe0, 0x7a, 0x75, 0xb3, 0x39,
		0xc6, 0xed, 0x10, 0x98, 0xbc, 0xda, 0xdc, 0x52, 0x76, 0xdb, 0xad, 0xae, 0xb2, 0xd7, 0x78, 0x55,
		0xdb, 0xdf, 0x7b, 0x71, 0xaf, 0xf5, 0xf2, 0x9e, 0x2a, 0x99, 0x13, 0x6a, 0x1f, 0xe2, 0x51, 0x6f,
		0x83, 0x34, 0x39, 0x29, 0x74, 0x1a, 0x66, 0x4d, 0x4b, 0x7a, 0x08, 0xad, 0xc2, 0xf2, 0x5e, 0x4b,
		0xeb, 0x34, 0xb7, 0x14, 0x4d, 0xb9, 0x71, 0x43, 0x69, 0x74, 0x3b, 0xec, 0x05, 0x22, 0xd4, 0xee,
		0xc6, 0x0f, 0xf5, 0x5b, 0x19, 0x58, 0x9d, 0x31, 0x13, 0x54, 0xe3, 0x77, 0x07, 0x76, 0x9d, 0xf9,
		0xe4, 0x3c, 0xb3, 0xdf, 0x24, 0x29, 0xbf, 0xad, 0x7b, 0x01, 0xbf, 0x6a, 0x5c, 0x04, 0x62, 0x25,
		0x3b, 0x30, 0xfb, 0x26, 0xf6, 0xf8, 0x83, 0x0d, 0xbb, 0x50, 0x2c, 0x8f, 0xe5, 0xec, 0xcd, 0xe6,
		0x13, 0x80, 0x5c, 0xc7, 0x37, 0x03, 0xf3, 0x0e, 0xd6, 0x4c, 0x5b, 0xbc, 0xee, 0x90, 0x0b, 0x46,
		0x56, 0x95, 0x44, 0x4f, 0xd3, 0x0e, 0x42, 0x6d, 0x1b, 0x0f, 0xf4, 0x09, 0x6d, 0x12, 0xc0, 0x33,
		0xaa, 0x24, 0x7a, 0x42, 0xed, 0xf3, 0x50, 0xea, 0x39, 0x23, 0x52, 0x75, 0x31, 0x3d, 0x92, 0x2f,
		0x52, 0x6a, 0x91, 0xc9, 0x42, 0x15, 0x5e, 0x4f, 0x8f, 0x9f, 0x95, 0x4a, 0x6a, 0x91, 0xc9, 0x98,
		0xca, 0x05, 0x58, 0xd6, 0x07, 0x03, 0x8f, 0x90, 0x0b, 0x22, 0x76, 0x43, 0x28, 0x87, 0x62, 0xaa,
		0xb8, 0x7e, 0x13, 0xf2, 0xc2, 0x0e, 0x24, 0x25, 0x13, 0x4b, 0x68, 0x2e, 0xbb, 0xf6, 0xa6, 0x37,
		0x0a, 0x6a, 0xde, 0x16, 0x9d, 0xe7, 0xa1, 0x64, 0xfa, 0xda, 0xf8, 0x95, 0x3c, 0x7d, 0x2e, 0xbd,
		0x91, 0x57, 0x8b, 0xa6, 0x1f, 0xbe, 0x30, 0x56, 0xdf, 0x4d, 0x43, 0x39, 0xfe, 0xca, 0x8f, 0xb6,
		0x20, 0x6f, 0x39, 0x86, 0x4e, 0x5d, 0x8b, 0x7d, 0x62, 0xda, 0x48, 0xf8, 0x30, 0xb0, 0xb9, 0xc3,
		0xf5, 0xd5, 0x10, 0xb9, 0xfe, 0x77, 0x29, 0xc8, 0x0b, 0x31, 0x3a, 0x05, 0x59, 0x57, 0x0f, 0x0e,
		0x29, 0x5d, 0xae, 0x9e, 0x96, 0x52, 0x2a, 0x6d, 0x13, 0xb9, 0xef, 0xea, 0x36, 0x75, 0x01, 0x2e,
		0x27, 0x6d, 0xb2, 0xaf, 0x16, 0xd6, 0x7b, 0xf4, 0xfa, 0xe1, 0x0c, 0x87, 0xd8, 0x0e, 0x7c, 0xb1,
		0xaf, 0x5c, 0xde, 0xe0, 0x62, 0xf4, 0x34, 0xac, 0x04, 0x9e, 0x6e, 0x5a, 0x31, 0xdd, 0x2c, 0xd5,
		0x95, 0x44, 0x47, 0xa8, 0x2c, 0xc3, 0x19, 0xc1, 0xdb, 0xc3, 0x81, 0x6e, 0x1c, 0xe2, 0xde, 0x18,
		0xb4, 0x40, 0x9f, 0x19, 0x4e, 0x73, 0x85, 0x2d, 0xde, 0x2f, 0xb0, 0xd5, 0xef, 0xa6, 0x60, 0x45,
		0x5c, 0x98, 0x7a, 0xa1, 0xb1, 0x76, 0x01, 0x74, 0xdb, 0x76, 0x82, 0xa8, 0xb9, 0xa6, 0x5d, 0x79,
		0x0a, 0xb7, 0x59, 0x0b, 0x41, 0x6a, 0x84, 0x60, 0x7d, 0x08, 0x30, 0xee, 0x39, 0xd6, 0x6c, 0x67,
		0xa1, 0xc8, 0x3f, 0xe1, 0xd0, 0xef, 0x80, 0xec, 0x8a, 0x0d, 0x4c, 0x44, 0x6e, 0x56, 0x68, 0x0d,
		0x72, 0x07, 0x78, 0x60, 0xda, 0xfc, 0x61, 0x96, 0x35, 0xc4, 0x43, 0x48, 0x36, 0x7c, 0x08, 0xa9,
		0x7f, 0x1e, 0x56, 0x0d, 0x67, 0x38, 0x39, 0xdd, 0xba, 0x34, 0x71, 0xcd, 0xf7, 0x3f, 0x9b, 0x7a,
		0x0d, 0xc6, 0x25, 0xe6, 0xff, 0xa5, 0x52, 0xbf, 0x95, 0xce, 0x6c, 0xb7, 0xeb, 0x5f, 0x4f, 0xaf,
		0x6f, 0x33, 0x68, 0x5b, 0xac, 0x54, 0xc5, 0x7d, 0x0b, 0x1b, 0x64, 0xf6, 0xf0, 0xa5, 0x0b, 0xf0,
		0xc9, 0x81, 0x19, 0x1c, 0x8e, 0x0e, 0x36, 0x0d, 0x67, 0x78, 0x69, 0xe0, 0x0c, 0x9c, 0xf1, 0xa7,
		0x4f, 0xd2, 0xa2, 0x0d, 0xfa, 0x1f, 0xff, 0xfc, 0x59, 0x08, 0xa5, 0xeb, 0x89, 0xdf, 0x4a, 0xe5,
		0x3d, 0x58, 0xe5, 0xca, 0x1a, 0xfd, 0xfe, 0xc2, 0x6e, 0x11, 0xe8, 0x81, 0x6f, 0x58, 0x95, 0x6f,
		0xbe, 0x47, 0xd3, 0xb5, 0xba, 0xc2, 0xa1, 0xa4, 0x8f, 0x5d, 0x34, 0x64, 0x15, 0x1e, 0x8e, 0xf1,
		0xb1, 0xa3, 0x89, 0xbd, 0x04, 0xc6, 0xef, 0x70, 0xc6, 0xd5, 0x08, 0x63, 0x87, 0x43, 0xe5, 0x06,
		0x2c, 0x9d, 0x84, 0xeb, 0xaf, 0x39, 0x57, 0x09, 0x47, 0x49, 0xb6, 0x61, 0x99, 0x92, 0x18, 0x23,
		0x3f, 0x70, 0x86, 0x34, 0xee, 0x3d, 0x98, 0xe6, 0x6f, 0xde, 0x63, 0x67, 0xa5, 0x4c, 0x60, 0x8d,
		0x10, 0x25, 0xcb, 0x40, 0x3f, 0x39, 0xf5, 0xb0, 0x61, 0x25, 0x30, 0xfc, 0x2d, 0x9f, 0x48, 0xa8,
		0x2f, 0xbf, 0x04, 0x6b, 0xe4, 0x7f, 0x1a, 0x96, 0xa2, 0x33, 0x49, 0x7e, 0xf0, 0xaa, 0x7c, 0xf7,
		0x4d, 0x76, 0x1c, 0x57, 0x43, 0x82, 0xc8, 0x9c, 0x22, 0xbb, 0x38, 0xc0, 0x41, 0x80, 0x3d, 0x5f,
		0xd3, 0xad, 0x59, 0xd3, 0x8b, 0xbc, 0x18, 0x54, 0xbe, 0xf2, 0x7e, 0x7c, 0x17, 0xb7, 0x19, 0xb2,
		0x66, 0x59, 0xf2, 0x3e, 0x9c, 0x9e, 0xe1, 0x15, 0x73, 0x70, 0xbe, 0xc5, 0x39, 0xd7, 0xa6, 0x3c,
		0x83, 0xd0, 0xb6, 0x41, 0xc8, 0xc3, 0xbd, 0x9c, 0x83, 0xf3, 0xd7, 0x38, 0x27, 0xe2, 0x58, 0xb1,
		0xa5, 0x84, 0xf1, 0x26, 0xac, 0xdc, 0xc1, 0xde, 0x81, 0xe3, 0xf3, 0x57, 0x9a, 0x39, 0xe8, 0xde,
		0xe6, 0x74, 0xcb, 0x1c, 0x48, 0x9f, 0x6d, 0x08, 0xd7, 0x35, 0xc8, 0xf7, 0x75, 0x03, 0xcf, 0x41,
		0xf1, 0x55, 0x4e, 0xb1, 0x48, 0xf4, 0x09, 0xb4, 0x06, 0xa5, 0x81, 0xc3, 0x33, 0x53, 0x32, 0xfc,
		0x1d, 0x0e, 0x2f, 0x0a, 0x0c, 0xa7, 0x70, 0x1d, 0x77, 0x64, 0x91, 0xb4, 0x95, 0x4c, 0xf1, 0xeb,
		0x82, 0x42, 0x60, 0x38, 0xc5, 0x09, 0xcc, 0xfa, 0x1b, 0x82, 0xc2, 0x8f, 0xd8, 0xf3, 0x05, 0x28,
		0x3a, 0xb6, 0x75, 0xe4, 0xd8, 0xf3, 0x4c, 0xe2, 0x37, 0x39, 0x03, 0x70, 0x08, 0x21, 0xb8, 0x0e,
		0x85, 0x79, 0x37, 0xe2, 0x6b, 0xef, 0x8b, 0xe3, 0x21, 0x76, 0x60, 0x1b, 0x96, 0x45, 0x80, 0x32,
		0x1d, 0x7b, 0x0e, 0x8a, 0xdf, 0xe6, 0x14, 0xe5, 0x08, 0x8c, 0x2f, 0x23, 0xc0, 0x7e, 0x30, 0xc0,
		0xf3, 0x90, 0xbc, 0x2b, 0x96, 0xc1, 0x21, 0xdc, 0x94, 0x07, 0xd8, 0x36, 0x0e, 0xe7, 0x63, 0xf8,
		0x1d, 0x61, 0x4a, 0x81, 0x21, 0x14, 0x0d, 0x58, 0x1a, 0xea, 0x9e, 0x7f, 0xa8, 0x5b, 0x73, 0x6d,
		0xc7, 0xef, 0x72, 0x8e, 0x52, 0x08, 0xe2, 0x16, 0x19, 0xd9, 0x27, 0xa1, 0xf9, 0xba, 0xb0, 0x48,
		0x04, 0xc6, 0x8f, 0x9e, 0x1f, 0xd0, 0x27, 0xad, 0x93, 0xb0, 0xfd, 0x9e, 0x38, 0x7a, 0x0c, 0xbb,
		0x1b, 0x65, 0xbc, 0x0e, 0x05, 0xdf, 0xbc, 0x37, 0x17, 0xcd, 0xef, 0x8b, 0x9d, 0xa6, 0x00, 0x02,
		0x7e, 0x15, 0xce, 0xcc, 0x4c, 0x13, 0x73, 0x90, 0xfd, 0x01, 0x27, 0x3b, 0x35, 0x23, 0x55, 0xf0,
		0x90, 0x70, 0x52, 0xca, 0x3f, 0x14, 0x21, 0x01, 0x4f, 0x70, 0xb5, 0xc9, 0x5d, 0xc1, 0xd7, 0xfb,
		0x27, 0xb3, 0xda, 0x1f, 0x09, 0xab, 0x31, 0x6c, 0xcc, 0x6a, 0x5d, 0x38, 0xc5, 0x19, 0x4f, 0xb6,
		0xaf, 0xdf, 0x10, 0x81, 0x95, 0xa1, 0xf7, 0xe3, 0xbb, 0xfb, 0x79, 0x58, 0x0f, 0xcd, 0x29, 0x8a,
		0x52, 0x5f, 0x1b, 0xea, 0xee, 0x1c, 0xcc, 0xdf, 0xe4, 0xcc, 0x22, 0xe2, 0x87, 0x55, 0xad, 0xbf,
		0xab, 0xbb, 0x84, 0xfc, 0x15, 0xa8, 0x08, 0xf2, 0x91, 0xed, 0x61, 0xc3, 0x19, 0xd8, 0xe6, 0x3d,
		0xdc, 0x9b, 0x83, 0xfa, 0x8f, 0x27, 0xb6, 0x6a, 0x3f, 0x02, 0x27, 0xcc, 0x4d, 0x90, 0xc2, 0x5a,
		0x45, 0x33, 0x87, 0xae, 0xe3, 0x05, 0x09, 0x8c, 0xdf, 0x12, 0x3b, 0x15, 0xe2, 0x9a, 0x14, 0x26,
		0x2b, 0x50, 0xa6, 0xcd, 0x79, 0x5d, 0xf2, 0x4f, 0x38, 0xd1, 0xd2, 0x18, 0xc5, 0x03, 0x87, 0xe1,
		0x0c, 0x5d, 0xdd, 0x9b, 0x27, 0xfe, 0xfd, 0xa9, 0x08, 0x1c, 0x1c, 0xc2, 0x03, 0x47, 0x70, 0xe4,
		0x62, 0x92, 0xed, 0xe7, 0x60, 0xf8, 0xb6, 0x08, 0x1c, 0x02, 0xc3, 0x29, 0x44, 0xc1, 0x30, 0x07,
		0xc5, 0x9f, 0x09, 0x0a, 0x81, 0x21, 0x14, 0x9f, 0x1b, 0x27, 0x5a, 0x0f, 0x0f, 0x4c, 0x3f, 0xf0,
		0x58, 0x29, 0xfc, 0x60, 0xaa, 0x3f, 0x7f, 0x3f, 0x5e, 0x84, 0xa9, 0x11, 0xa8, 0x7c, 0x13, 0x96,
		0x27, 0x4a, 0x0c, 0x94, 0xf4, 0xfb, 0x95, 0xca, 0x4f, 0x7c, 0xc0, 0x83, 0x51, 0xbc, 0xc2, 0x90,
		0x77, 0xc8, 0xbe, 0xc7, 0xeb, 0x80, 0x64, 0xb2, 0x37, 0x3f, 0x08, 0xb7, 0x3e, 0x56, 0x06, 0xc8,
		0x37, 0x60, 0x29, 0x56, 0x03, 0x24, 0x53, 0xfd, 0x24, 0xa7, 0x2a, 0x45, 0x4b, 0x00, 0xf9, 0x0a,
		0x64, 0x49, 0x3e, 0x4f, 0x86, 0xff, 0x14, 0x87, 0x53, 0x75, 0xf9, 0xd3, 0x90, 0x17, 0x79, 0x3c,
		0x19, 0xfa, 0xd3, 0x1c, 0x1a, 0x42, 0x08, 0x5c, 0xe4, 0xf0, 0x64, 0xf8, 0xcf, 0x08, 0xb8, 0x80,
		0x10, 0xf8, 0xfc, 0x26, 0xfc, 0xab, 0x9f, 0xcd, 0xf2, 0x38, 0x2c, 0x6c, 0x77, 0x1d, 0x16, 0x79,
		0xf2, 0x4e, 0x46, 0x7f, 0x91, 0x0f, 0x2e, 0x10, 0xf2, 0x73, 0x90, 0x9b, 0xd3, 0xe0, 0x3f, 0xc7,
		0xa1, 0x4c, 0x5f, 0x6e, 0x40, 0x31, 0x92, 0xb0, 0x93, 0xe1, 0x3f, 0xcf, 0xe1, 0x51, 0x14, 0x99,
		0x3a, 0x4f, 0xd8, 0xc9, 0x04, 0xbf, 0x20, 0xa6, 0xce, 0x11, 0xc4, 0x6c, 0x22, 0x57, 0x27, 0xa3,
		0x7f, 0x51, 0x58, 0x5d, 0x40, 0xe4, 0x17, 0xa0, 0x10, 0xc6, 0xdf, 0x64, 0xfc, 0x2f, 0x71, 0xfc,
		0x18, 0x43, 0x2c, 0x10, 0x89, 0xff, 0xc9, 0x14, 0xbf, 0x2c, 0x2c, 0x10, 0x41, 0x91, 0x63, 0x34,
		0x99, 0xd3, 0x93, 0x99, 0x7e, 0x45, 0x1c, 0xa3, 0x89, 0x94, 0x4e, 0x76, 0x93, 0x86, 0xc1, 0x64,
		0x8a, 0x5f, 0x15, 0xbb, 0x49, 0xf5, 0xc9, 0x34, 0x26, 0x93, 0x64, 0x32, 0xc7, 0x97, 0xc4, 0x34,
		0x26, 0x72, 0xa4, 0xdc, 0x06, 0x34, 0x9d, 0x20, 0x93, 0xf9, 0xbe, 0xcc, 0xf9, 0x56, 0xa6, 0xf2,
		0xa3, 0xfc, 0x32, 0x9c, 0x9a, 0x9d, 0x1c, 0x93, 0x59, 0xbf, 0xf2, 0xc1, 0xc4, 0x75, 0x26, 0x9a,
		0x1b, 0xe5, 0xee, 0x38, 0xca, 0x46, 0x13, 0x63, 0x32, 0xed, 0x5b, 0x1f, 0xc4, 0x03, 0x6d, 0x34,
		0x2f, 0xca, 0x35, 0x80, 0x71, 0x4e, 0x4a, 0xe6, 0x7a, 0x9b, 0x73, 0x45, 0x40, 0xe4, 0x68, 0xf0,
		0x94, 0x94, 0x8c, 0xff, 0xaa, 0x38, 0x1a, 0x1c, 0x41, 0x8e, 0x86, 0xc8, 0x46, 0xc9, 0xe8, 0x77,
		0xc4, 0xd1, 0x10, 0x10, 0xf9, 0x3a, 0xe4, 0xed, 0x91, 0x65, 0x11, 0xdf, 0x42, 0x0f, 0xfe, 0x49,
		0x56, 0xe5, 0xdf, 0xee, 0x73, 0xb0, 0x00, 0xc8, 0x57, 0x20, 0x87, 0x87, 0x07, 0xb8, 0x97, 0x84,
		0xfc, 0xf7, 0xfb, 0x22, 0x9e, 0x10, 0x6d, 0xf9, 0x05, 0x00, 0x76, 0x99, 0xa6, 0x1f, 0x8a, 0x12,
		0xb0, 0xff, 0x71, 0x9f, 0xff, 0x58, 0x62, 0x0c, 0x19, 0x13, 0xb0, 0x9f, 0x5e, 0x3c, 0x98, 0xe0,
		0xfd, 0x38, 0x01, 0xbd, 0x80, 0x5f, 0x83, 0xc5, 0x5b, 0xbe, 0x63, 0x07, 0xfa, 0x20, 0x09, 0xfd,
		0x9f, 0x1c, 0x2d, 0xf4, 0x89, 0xc1, 0x86, 0x8e, 0x87, 0x03, 0x7d, 0xe0, 0x27, 0x61, 0xff, 0x8b,
		0x63, 0x43, 0x00, 0x01, 0x1b, 0xba, 0x1f, 0xcc, 0xb3, 0xee, 0xef, 0x09, 0xb0, 0x00, 0x90, 0x49,
		0x93, 0xff, 0x6f, 0xe3, 0xa3, 0x24, 0xec, 0xf7, 0xc5, 0xa4, 0xb9, 0xbe, 0xfc, 0x69, 0x28, 0x90,
		0x7f, 0xd9, 0x2f, 0xa0, 0x12, 0xc0, 0xff, 0xcd, 0xc1, 0x63, 0x04, 0x19, 0xd9, 0x0f, 0x7a, 0x81,
		0x99, 0x6c, 0xec, 0xff, 0xe1, 0x3b, 0x2d, 0xf4, 0xe5, 0x1a, 0x14, 0xfd, 0xa0, 0xd7, 0x1b, 0xf1,
		0x8a, 0x26, 0x01, 0xfe, 0xbf, 0xf7, 0xc3, 0x4b, 0x6e, 0x88, 0xa9, 0x2b, 0xb3, 0xdf, 0xeb, 0x60,
		0xdb, 0xd9, 0x76, 0xd8, 0x4b, 0xdd, 0x6b, 0xd5, 0xe4, 0x27, 0x37, 0xf8, 0x5a, 0x96, 0xd7, 0x9c,
		0xf6, 0x68, 0xc8, 0x9f, 0xde, 0xf2, 0xa2, 0xbd, 0x7e, 0xb2, 0x37, 0xbb, 0xaa, 0x0c, 0x79, 0x3a,
		0xf2, 0xde, 0x68, 0x88, 0x36, 0x21, 0xff, 0x22, 0x3e, 0x7a, 0x29, 0xf2, 0xa3, 0x34, 0xb4, 0x19,
		0x8e, 0x23, 0x7a, 0xd4, 0x50, 0xa7, 0x7a, 0x79, 0xac, 0x8f, 0x24, 0xc8, 0xbc, 0x88, 0x8f, 0xf8,
		0xcf, 0x67, 0xc8, 0xbf, 0x68, 0x0d, 0x72, 0x82, 0x8a, 0xc8, 0x58, 0xa3, 0xfa, 0x1a, 0x40, 0xd7,
		0x71, 0xfd, 0xa3, 0xee, 0xc8, 0xbb, 0x73, 0x84, 0x9e, 0x84, 0x1c, 0x6d, 0x51, 0x5c, 0xf1, 0xf2,
		0xf2, 0x78, 0x38, 0x2a, 0x56, 0x59, 0x2f, 0x55, 0x23, 0xfa, 0xfc, 0xd7, 0x78, 0x51, 0x35, 0x22,
		0x56, 0x59, 0x6f, 0xf5, 0x71, 0xce, 0x86, 0x4a, 0x90, 0xaa, 0x51, 0xca, 0x8c, 0x9a, 0xaa, 0x91,
		0x56, 0x9d, 0x22, 0x33, 0x6a, 0xaa, 0x4e, 0x95, 0xe8, 0xd8, 0x54, 0xcc, 0x95, 0xea, 0x0c, 0xc2,
		0x95, 0x6a, 0xd5, 0x11, 0x64, 0x5f, 0xb4, 0x9d, 0x00, 0x5d, 0x84, 0x5c, 0xdd, 0x1c, 0x1c, 0x06,
		0xfc, 0x65, 0x78, 0x75, 0x3c, 0x30, 0x15, 0x13, 0x1d, 0x95, 0x69, 0x90, 0xe5, 0x2a, 0xd6, 0x81,
		0x73, 0x97, 0x7f, 0x70, 0x64, 0x0d, 0xf4, 0x09, 0x28, 0xd4, 0x4d, 0x52, 0x7b, 0x2a, 0x76, 0x8f,
		0xff, 0x6e, 0xab, 0x1c, 0xb1, 0x29, 0xc1, 0x8f, 0x15, 0xaa, 0x0d, 0xa2, 0xcd, 0x79, 0x51, 0x15,
		0xb2, 0x3b, 0x8e, 0xe3, 0x72, 0xd3, 0x4c, 0xa2, 0x68, 0x1f, 0x42, 0x90, 0xed, 0x8e, 0x3c, 0x9b,
		0x8f, 0x49, 0xff, 0xaf, 0xe7, 0xbf, 0xff, 0x8f, 0x8f, 0xa5, 0xbe, 0xf1, 0x4f, 0x8f, 0xa5, 0xfe,
		0x3f, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x40, 0xc2, 0xe3, 0xb4, 0x32, 0x00, 0x00,
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
func (this *ProtoNum) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protonum.ProtoNum{")
	if this.KeyValue != nil {
		s = append(s, "KeyValue: "+fmt.Sprintf("%#v", this.KeyValue)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *KeyValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.KeyValue{")
	if this.Key != nil {
		s = append(s, "Key: "+valueToGoStringProtonum(this.Key, "string")+",\n")
	}
	if this.Value != nil {
		s = append(s, "Value: "+valueToGoStringProtonum(this.Value, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TopsyTurvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.TopsyTurvy{")
	if this.Topsy != nil {
		s = append(s, "Topsy: "+fmt.Sprintf("%#v", this.Topsy)+",\n")
	}
	if this.Turvy != nil {
		s = append(s, "Turvy: "+fmt.Sprintf("%#v", this.Turvy)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Topsy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.Topsy{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringProtonum(this.A, "int64")+",\n")
	}
	if this.B != nil {
		s = append(s, "B: "+valueToGoStringProtonum(this.B, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Turvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.Turvy{")
	if this.B != nil {
		s = append(s, "B: "+valueToGoStringProtonum(this.B, "int64")+",\n")
	}
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringProtonum(this.A, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Knot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&protonum.Knot{")
	if this.Bight != nil {
		s = append(s, "Bight: "+fmt.Sprintf("%#v", this.Bight)+",\n")
	}
	if this.Elbow != nil {
		s = append(s, "Elbow: "+valueToGoStringProtonum(this.Elbow, "bool")+",\n")
	}
	if this.BitterEnd != nil {
		s = append(s, "BitterEnd: "+fmt.Sprintf("%#v", this.BitterEnd)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BightKnot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.BightKnot{")
	if this.Loop != nil {
		s = append(s, "Loop: "+fmt.Sprintf("%#v", this.Loop)+",\n")
	}
	if this.Turn != nil {
		s = append(s, "Turn: "+valueToGoStringProtonum(this.Turn, "bool")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtonum(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("protonum.proto", fileDescriptorProtonum) }

var fileDescriptorProtonum = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xd3, 0x42, 0x3a, 0x95, 0x2a, 0xa3, 0x87, 0xe0, 0xa1, 0x94, 0x15, 0xa1, 0x82,
	0xa6, 0xd0, 0xa3, 0xb7, 0xae, 0xf4, 0x54, 0x11, 0x59, 0x8a, 0x07, 0x6f, 0x46, 0x63, 0x1a, 0x68,
	0xb2, 0x21, 0xee, 0x2a, 0x79, 0x2b, 0x5f, 0x49, 0x9f, 0xc0, 0x47, 0x90, 0x9d, 0xdd, 0x34, 0xea,
	0x6d, 0xfe, 0x99, 0xef, 0xff, 0xff, 0x64, 0x61, 0x54, 0xd5, 0x4a, 0xab, 0xd2, 0x14, 0x31, 0x0d,
	0x18, 0xb6, 0xfa, 0xe4, 0x32, 0xcb, 0xf5, 0xc6, 0x24, 0xf1, 0x93, 0x2a, 0x66, 0x99, 0xca, 0xd4,
	0x8c, 0x2e, 0x89, 0x79, 0x21, 0x45, 0x82, 0x26, 0x67, 0xe4, 0x57, 0x10, 0xde, 0xd9, 0xe1, 0xd6,
	0x14, 0x18, 0x43, 0xb8, 0x4a, 0x9b, 0xfb, 0xc7, 0xad, 0x49, 0xa3, 0xbd, 0x49, 0x30, 0x1d, 0xce,
	0x31, 0xde, 0xf5, 0xb4, 0x17, 0xb9, 0x63, 0xf8, 0xbc, 0xe3, 0xf1, 0x10, 0x82, 0x55, 0xda, 0x44,
	0x6c, 0xc2, 0xa6, 0x03, 0x69, 0x47, 0x3c, 0x86, 0x7e, 0x1b, 0x65, 0x77, 0x4e, 0xf0, 0x07, 0x80,
	0xb5, 0xaa, 0x5e, 0x9b, 0xb5, 0xa9, 0xdf, 0x1a, 0x3c, 0x83, 0x3e, 0x29, 0xf2, 0x0d, 0xe7, 0x07,
	0x5d, 0x1d, 0xad, 0xa5, 0xbb, 0x12, 0x66, 0x79, 0x8a, 0xfa, 0x8b, 0xd9, 0xb5, 0x74, 0x57, 0x7e,
	0xea, 0xd3, 0x70, 0x1f, 0xd8, 0x82, 0x22, 0x03, 0xc9, 0x16, 0x56, 0x09, 0x72, 0x06, 0x92, 0x09,
	0x82, 0xa8, 0x9b, 0xd6, 0x1e, 0x12, 0xce, 0xe2, 0xa1, 0x05, 0x37, 0xd0, 0x5b, 0x95, 0x4a, 0xe3,
	0x39, 0xf4, 0x45, 0x9e, 0x6d, 0x74, 0xc4, 0xe8, 0x39, 0x8e, 0xba, 0x62, 0x5a, 0x5b, 0x46, 0x3a,
	0xc2, 0xfe, 0xee, 0x72, 0x9b, 0xa8, 0x77, 0x0a, 0x09, 0xa5, 0x13, 0x78, 0x01, 0x03, 0x91, 0x6b,
	0x9d, 0xd6, 0xcb, 0xf2, 0x39, 0x0a, 0xe8, 0xeb, 0x47, 0xbf, 0xde, 0xd4, 0xfa, 0x3b, 0x80, 0x5f,
	0x5b, 0xda, 0xe7, 0x22, 0x87, 0xde, 0x8d, 0x52, 0x95, 0x7f, 0x9a, 0xff, 0x2e, 0xba, 0x21, 0x42,
	0x6f, 0x6d, 0xea, 0xd2, 0x77, 0xd2, 0x2c, 0xc2, 0xef, 0xcf, 0x31, 0xfb, 0xf8, 0x1a, 0xb3, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xec, 0x74, 0x43, 0x25, 0x02, 0x00, 0x00,
}
