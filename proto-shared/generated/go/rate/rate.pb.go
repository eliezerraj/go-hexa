// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rate/rate.proto

package rate

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

type Rate struct {
	Rate                 int32    `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}
func (*Rate) Descriptor() ([]byte, []int) {
	return fileDescriptor_466953fd5f8ee6df, []int{0}
}

func (m *Rate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rate.Unmarshal(m, b)
}
func (m *Rate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rate.Marshal(b, m, deterministic)
}
func (m *Rate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rate.Merge(m, src)
}
func (m *Rate) XXX_Size() int {
	return xxx_messageInfo_Rate.Size(m)
}
func (m *Rate) XXX_DiscardUnknown() {
	xxx_messageInfo_Rate.DiscardUnknown(m)
}

var xxx_messageInfo_Rate proto.InternalMessageInfo

func (m *Rate) GetRate() int32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Rate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RateRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_466953fd5f8ee6df, []int{1}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type RateResponse struct {
	Rate                 *Rate    `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_466953fd5f8ee6df, []int{2}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetRate() *Rate {
	if m != nil {
		return m.Rate
	}
	return nil
}

func init() {
	proto.RegisterType((*Rate)(nil), "rate.Rate")
	proto.RegisterType((*RateRequest)(nil), "rate.RateRequest")
	proto.RegisterType((*RateResponse)(nil), "rate.RateResponse")
}

func init() { proto.RegisterFile("rate/rate.proto", fileDescriptor_466953fd5f8ee6df) }

var fileDescriptor_466953fd5f8ee6df = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc4, 0x30,
	0x0c, 0x85, 0x29, 0x3a, 0x40, 0xe7, 0x22, 0x10, 0x9e, 0x2a, 0x06, 0x54, 0x75, 0xe1, 0x10, 0xa2,
	0x95, 0x8e, 0x09, 0x89, 0x05, 0x16, 0x98, 0xc3, 0xc6, 0x96, 0x4b, 0xad, 0x36, 0x03, 0x49, 0x49,
	0x5c, 0x7e, 0x3f, 0x8a, 0xa3, 0x93, 0xb2, 0x44, 0xf6, 0x17, 0xbf, 0xf7, 0x12, 0xc3, 0x75, 0xd0,
	0x4c, 0x43, 0x3a, 0xfa, 0x25, 0x78, 0xf6, 0xb8, 0x49, 0x75, 0xf7, 0x0a, 0x1b, 0xa5, 0x99, 0x10,
	0x41, 0xfa, 0xa6, 0x6a, 0xab, 0xdd, 0x99, 0x92, 0x1a, 0x5b, 0xa8, 0x47, 0x8a, 0x26, 0xd8, 0x85,
	0xad, 0x77, 0xcd, 0x69, 0x5b, 0xed, 0xb6, 0xaa, 0x44, 0xdd, 0x3d, 0xd4, 0x49, 0xad, 0xe8, 0x77,
	0xa5, 0xc8, 0xd8, 0xc0, 0x85, 0x36, 0xc6, 0xaf, 0x8e, 0xc5, 0x67, 0xab, 0x8e, 0x6d, 0xd7, 0xc3,
	0x65, 0x1e, 0x8c, 0x8b, 0x77, 0x91, 0xf0, 0xae, 0x88, 0xab, 0xf7, 0xd0, 0xcb, 0xbb, 0x64, 0x42,
	0xf8, 0xfe, 0x33, 0x1b, 0x7f, 0x51, 0xf8, 0xb3, 0x86, 0xf0, 0x05, 0xae, 0x3e, 0x88, 0x13, 0x79,
	0xcb, 0x86, 0x78, 0x53, 0x48, 0x72, 0xfa, 0x2d, 0x96, 0x28, 0xe7, 0x74, 0x27, 0xef, 0x8f, 0xdf,
	0x0f, 0x93, 0xe5, 0x79, 0x3d, 0xf4, 0xc6, 0xff, 0x0c, 0xf2, 0xf5, 0xa7, 0x38, 0xeb, 0x40, 0xe3,
	0x30, 0x91, 0xa3, 0x24, 0x19, 0x87, 0xc9, 0xcb, 0x66, 0x0e, 0xe7, 0x72, 0xff, 0xfc, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x1f, 0x23, 0x3c, 0x2d, 0x01, 0x00, 0x00,
}