// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostAvatar

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
	//二进制图片流
	Avatar []byte `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	//文件大小
	Filesize int64 `protobuf:"varint,2,opt,name=Filesize,proto3" json:"Filesize,omitempty"`
	//文件后缀
	Fileext              string   `protobuf:"bytes,3,opt,name=Fileext,proto3" json:"Fileext,omitempty"`
	SessionId            string   `protobuf:"bytes,4,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
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

func (m *Request) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *Request) GetFilesize() int64 {
	if m != nil {
		return m.Filesize
	}
	return 0
}

func (m *Request) GetFileext() string {
	if m != nil {
		return m.Fileext
	}
	return ""
}

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type Response struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//不完整的头像地址
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=Avatar_url,json=AvatarUrl,proto3" json:"Avatar_url,omitempty"`
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

func (m *Response) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostAvatar.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostAvatar.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostAvatar.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xad, 0xd1, 0xa6, 0x19, 0x3c, 0xc8, 0x20, 0x1a, 0x5a, 0x85, 0x98, 0x53, 0x4f, 0x2b,
	0xe8, 0x2f, 0xf0, 0x10, 0xc1, 0x83, 0x20, 0x5b, 0xc4, 0x9b, 0x12, 0x75, 0x08, 0x81, 0x24, 0x1b,
	0x67, 0x36, 0xa5, 0xfa, 0xeb, 0x25, 0xbb, 0x5b, 0x7b, 0xd2, 0xd3, 0xce, 0xf7, 0x66, 0x77, 0xe7,
	0xcd, 0x83, 0x45, 0xcf, 0xc6, 0x9a, 0x2b, 0xda, 0x94, 0x6d, 0xdf, 0xd0, 0xf6, 0x54, 0x4e, 0xc5,
	0xb3, 0xca, 0xa8, 0xb6, 0x7e, 0x67, 0xa3, 0x84, 0xd7, 0xea, 0xd1, 0x88, 0xbd, 0x5d, 0x97, 0xb6,
	0xe4, 0x7c, 0x01, 0xf1, 0x03, 0x89, 0x94, 0x15, 0xe1, 0x31, 0x44, 0x52, 0x7e, 0xa5, 0x93, 0x6c,
	0xb2, 0x4c, 0xf4, 0x58, 0xe6, 0x03, 0xc4, 0x9a, 0x3e, 0x07, 0x12, 0x8b, 0xa7, 0x30, 0xf5, 0x2f,
	0x5c, 0xff, 0x48, 0x07, 0xc2, 0x39, 0xcc, 0xee, 0xea, 0x86, 0xa4, 0xfe, 0xa6, 0x74, 0x3f, 0x9b,
	0x2c, 0x23, 0xfd, 0xcb, 0x98, 0x42, 0x3c, 0xd6, 0xb4, 0xb1, 0x69, 0xe4, 0x3e, 0xdd, 0x22, 0x9e,
	0x43, 0xb2, 0x22, 0x91, 0xda, 0x74, 0xf7, 0x1f, 0xe9, 0x81, 0xeb, 0xed, 0x84, 0xfc, 0x19, 0x66,
	0x9a, 0xa4, 0x37, 0x9d, 0x10, 0x9e, 0xc0, 0x61, 0xc1, 0xdc, 0x99, 0x60, 0xcb, 0xc3, 0xe8, 0xa6,
	0x60, 0x6e, 0xa5, 0x72, 0x33, 0x13, 0x1d, 0x08, 0x2f, 0x00, 0xbc, 0xaf, 0xd7, 0x81, 0x9b, 0x30,
	0x34, 0xf1, 0xca, 0x13, 0x37, 0xd7, 0x2f, 0x10, 0x17, 0x3e, 0x16, 0x5c, 0x01, 0xec, 0x52, 0xc0,
	0x4c, 0xfd, 0x91, 0x8f, 0x0a, 0xfb, 0xcf, 0x2f, 0xff, 0xb9, 0xe1, 0xad, 0xe6, 0x7b, 0x6f, 0x53,
	0x17, 0xf6, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x11, 0x6b, 0x27, 0x8b, 0x01, 0x00,
	0x00,
}
