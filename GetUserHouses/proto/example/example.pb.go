// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetUserHouses

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Mix                  []byte   `protobuf:"bytes,3,opt,name=Mix,proto3" json:"Mix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetMix() []byte {
	if m != nil {
		return m.Mix
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetUserHouses.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetUserHouses.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetUserHouses.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0x86, 0x40,
	0x10, 0xc5, 0xb3, 0x8f, 0x3e, 0x73, 0x28, 0x88, 0x25, 0x42, 0xb4, 0x83, 0x6c, 0x41, 0x9e, 0x36,
	0xa8, 0xbf, 0x41, 0x8a, 0xc0, 0xcb, 0x46, 0xb7, 0x2e, 0x56, 0xc3, 0xb2, 0x90, 0x8e, 0xcd, 0x68,
	0xd8, 0x7f, 0x1f, 0xea, 0x46, 0x74, 0xe8, 0x3b, 0xcd, 0x7b, 0x8f, 0x07, 0xf3, 0x9b, 0x81, 0xbc,
	0x67, 0x1a, 0xe8, 0x1a, 0xa7, 0xa6, 0xed, 0xdf, 0xf1, 0x67, 0x9a, 0x25, 0x55, 0x99, 0x23, 0xd3,
	0xfa, 0x57, 0x26, 0x23, 0xfc, 0x69, 0xee, 0x70, 0x78, 0x12, 0xe4, 0x7b, 0x1a, 0x05, 0x45, 0xe7,
	0x10, 0xd7, 0x28, 0xd2, 0x38, 0x54, 0x27, 0xb0, 0x91, 0xe6, 0x2b, 0x8d, 0x8a, 0xa8, 0x4c, 0xec,
	0x2c, 0xf5, 0x15, 0xc4, 0x16, 0x3f, 0x46, 0x94, 0x41, 0x9d, 0x43, 0xf2, 0x88, 0x22, 0x9e, 0x3a,
	0xff, 0x16, 0x2a, 0xbf, 0x81, 0x7e, 0x80, 0x43, 0x8b, 0xd2, 0x53, 0x27, 0xa8, 0x4e, 0xe1, 0xa0,
	0x62, 0xee, 0x28, 0xb4, 0x56, 0xa3, 0xce, 0x60, 0x5b, 0x31, 0xb7, 0xe2, 0xd2, 0xfd, 0x25, 0x0e,
	0x6e, 0x5e, 0x5a, 0xfb, 0x29, 0xdd, 0x14, 0x51, 0x79, 0x64, 0x67, 0x79, 0xe3, 0x20, 0xae, 0x56,
	0x7c, 0xf5, 0x0c, 0xc7, 0x7f, 0x68, 0xd5, 0x85, 0xf9, 0xff, 0x14, 0x13, 0x50, 0xb3, 0xcb, 0xdd,
	0xa5, 0x15, 0x53, 0xef, 0xbd, 0x6c, 0x97, 0xef, 0xdc, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x15,
	0x25, 0x40, 0x4b, 0x3c, 0x01, 0x00, 0x00,
}
