// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetSession

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
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	//返回用户名
	UserName             string   `protobuf:"bytes,3,opt,name=User_name,json=UserName,proto3" json:"User_name,omitempty"`
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

func (m *Response) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetSession.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSession.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc6, 0x30,
	0x10, 0xc4, 0xfd, 0xfc, 0xb0, 0x7f, 0xf6, 0x24, 0x8b, 0x68, 0xb1, 0x1e, 0x6a, 0x2e, 0x7a, 0x8a,
	0xa0, 0xcf, 0x50, 0x3c, 0xd5, 0x43, 0x4b, 0xaf, 0x4a, 0xd4, 0xa5, 0x14, 0x4c, 0x52, 0x77, 0xab,
	0xe8, 0xdb, 0x4b, 0xd3, 0x48, 0x4f, 0x7a, 0x4a, 0xe6, 0xb7, 0xb3, 0x30, 0xb3, 0x50, 0x4e, 0xec,
	0x67, 0x7f, 0x43, 0x5f, 0xc6, 0x4e, 0x6f, 0xf4, 0xfb, 0xea, 0x40, 0xf1, 0x6c, 0xf0, 0xda, 0x8e,
	0x2f, 0xec, 0xb5, 0xf0, 0xa7, 0xbe, 0xa7, 0xb9, 0x23, 0x91, 0xd1, 0x3b, 0x55, 0x42, 0xda, 0x90,
	0x88, 0x19, 0x08, 0x8f, 0x61, 0x2f, 0xe6, 0xbb, 0xd8, 0x55, 0xbb, 0xeb, 0xbc, 0x5d, 0xbe, 0xea,
	0x0a, 0xd2, 0x96, 0xde, 0x3f, 0x48, 0x66, 0xbc, 0x80, 0x3c, 0xae, 0x8c, 0xaf, 0xd1, 0xb2, 0x01,
	0xd5, 0x43, 0xd6, 0x92, 0x4c, 0xde, 0x09, 0xe1, 0x09, 0x1c, 0xd5, 0xcc, 0xce, 0x47, 0xd7, 0x2a,
	0xf0, 0x14, 0x92, 0x9a, 0xb9, 0x91, 0xa1, 0x38, 0x0c, 0x38, 0x2a, 0x2c, 0x21, 0xef, 0x85, 0xf8,
	0xc9, 0x19, 0x4b, 0xc5, 0x3e, 0x8c, 0xb2, 0x05, 0x3c, 0x18, 0x4b, 0xb7, 0x8f, 0x90, 0xd6, 0x6b,
	0x0d, 0xec, 0x00, 0xb6, 0xd4, 0x58, 0xe9, 0x3f, 0xfa, 0xe8, 0x98, 0xf7, 0xfc, 0xf2, 0x1f, 0xc7,
	0x1a, 0x54, 0x1d, 0x3c, 0x27, 0xe1, 0x38, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0xea,
	0xe5, 0x68, 0x3b, 0x01, 0x00, 0x00,
}
