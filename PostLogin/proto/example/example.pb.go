// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostLogin

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
	Mobile               string   `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Sessionid            string   `protobuf:"bytes,3,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
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

func (m *Response) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *Response) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostLogin.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostLogin.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostLogin.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xad, 0xc5, 0x6e, 0x77, 0x4e, 0x32, 0x48, 0x59, 0x5a, 0xc1, 0x92, 0x93, 0xa7, 0x08,
	0x7a, 0xf6, 0xb8, 0x37, 0x17, 0x96, 0x15, 0xbc, 0x78, 0x4a, 0xed, 0x10, 0x02, 0x6d, 0x66, 0x9d,
	0x59, 0x7f, 0xfd, 0xf7, 0xd2, 0x6c, 0xac, 0x17, 0x3d, 0x25, 0xdf, 0xbc, 0xbc, 0x90, 0x2f, 0xb0,
	0xea, 0x85, 0x07, 0xbe, 0xa1, 0x4f, 0xb7, 0xef, 0x77, 0xf4, 0xb3, 0xda, 0x34, 0xc5, 0x85, 0x67,
	0xbb, 0x0f, 0x2f, 0xc2, 0x56, 0xe5, 0xdd, 0xb6, 0xac, 0xc3, 0x03, 0xfb, 0x10, 0xcd, 0x0a, 0x8a,
	0x86, 0x54, 0x9d, 0x27, 0x3c, 0x87, 0xa9, 0xba, 0xaf, 0x6a, 0xb2, 0x9e, 0x5c, 0x97, 0xdd, 0x61,
	0x6b, 0xee, 0xa1, 0xe8, 0xe8, 0xf5, 0x8d, 0x74, 0xc0, 0x05, 0xcc, 0x1a, 0xde, 0x84, 0x1d, 0xe5,
	0x3c, 0x13, 0x2e, 0x61, 0xde, 0x3a, 0xd5, 0x0f, 0x96, 0x6d, 0x75, 0x9a, 0x92, 0x23, 0x9b, 0x27,
	0x98, 0x77, 0xa4, 0x3d, 0x47, 0x25, 0xbc, 0x80, 0xb3, 0x5a, 0x24, 0x72, 0xae, 0x8f, 0x70, 0xb8,
	0xb5, 0x16, 0x69, 0xd4, 0xe7, 0x6e, 0x26, 0xbc, 0x84, 0xf2, 0x91, 0x54, 0x03, 0xc7, 0xb0, 0xad,
	0xa6, 0x29, 0xfa, 0x1d, 0xdc, 0x3e, 0x43, 0x51, 0x8f, 0x72, 0xd8, 0x42, 0x79, 0x74, 0xc1, 0x2b,
	0xfb, 0xb7, 0xa4, 0xcd, 0x12, 0xcb, 0xf5, 0xff, 0x07, 0xc6, 0x67, 0x9a, 0x93, 0xcd, 0x2c, 0xfd,
	0xd7, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xab, 0xeb, 0xfa, 0x4e, 0x01, 0x00, 0x00,
}