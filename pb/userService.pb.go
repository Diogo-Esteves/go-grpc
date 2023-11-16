// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userService.proto

package _

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

type UserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_874b4b11a4ddc7c4, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
}

func init() {
	proto.RegisterFile("userService.proto", fileDescriptor_874b4b11a4ddc7c4)
}

var fileDescriptor_874b4b11a4ddc7c4 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01,
	0x09, 0x49, 0x71, 0x81, 0x48, 0x88, 0x88, 0x92, 0x1a, 0x17, 0x77, 0x68, 0x71, 0x6a, 0x51, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x38, 0x17, 0x3b, 0x48, 0x32, 0x3e, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x39, 0x88, 0x0d, 0xc4, 0xf5, 0x4c, 0x31, 0x32, 0x83, 0xa8, 0x83, 0x1a,
	0x27, 0xa4, 0xce, 0xc5, 0xed, 0x9e, 0x5a, 0x02, 0x12, 0x71, 0xaa, 0xf4, 0x4c, 0x11, 0x12, 0xd4,
	0x03, 0x1b, 0x89, 0x64, 0x92, 0x14, 0x2b, 0x98, 0xe7, 0xc4, 0x1c, 0xc5, 0xa8, 0x97, 0xc4, 0x06,
	0xb6, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x88, 0x72, 0xd2, 0x92, 0x00, 0x00, 0x00,
}
