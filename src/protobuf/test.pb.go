// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protobuf/test.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4be219910fb9593, []int{0}
}

type Phone struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Model                string   `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4be219910fb9593, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Phone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Phone) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func init() {
	proto.RegisterEnum("example.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "example.Phone")
}

func init() { proto.RegisterFile("src/protobuf/test.proto", fileDescriptor_b4be219910fb9593) }

var fileDescriptor_b4be219910fb9593 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0xf3,
	0x84, 0xd8, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x95, 0x1c, 0xb9, 0x58, 0x03, 0x32, 0xf2,
	0xf3, 0x52, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98,
	0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x09, 0x66, 0xb0, 0x20,
	0x84, 0xa3, 0x25, 0xcf, 0xc5, 0x09, 0x36, 0x22, 0xa4, 0xb2, 0x20, 0x55, 0x88, 0x83, 0x8b, 0xc5,
	0xc3, 0xdf, 0xd7, 0x55, 0x80, 0x01, 0xc4, 0x0a, 0xf7, 0x0f, 0xf2, 0x16, 0x60, 0x4c, 0x62, 0x03,
	0xdb, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x2d, 0xfd, 0x33, 0x8e, 0x00, 0x00, 0x00,
}
