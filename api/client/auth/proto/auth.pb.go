// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/client/auth/proto/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	api/client/auth/proto/auth.proto

It has these top-level messages:
	AuthReq
	AuthResp
	TopicReq
	TopicResp
*/
package auth

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

type METHOD int32

const (
	METHOD_Auth    METHOD = 0
	METHOD_SubAuth METHOD = 1
	METHOD_PubAuth METHOD = 2
)

var METHOD_name = map[int32]string{
	0: "Auth",
	1: "SubAuth",
	2: "PubAuth",
}
var METHOD_value = map[string]int32{
	"Auth":    0,
	"SubAuth": 1,
	"PubAuth": 2,
}

func (x METHOD) String() string {
	return proto.EnumName(METHOD_name, int32(x))
}
func (METHOD) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AuthReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pwd  string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *AuthReq) Reset()                    { *m = AuthReq{} }
func (m *AuthReq) String() string            { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()               {}
func (*AuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type AuthResp struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Verified bool   `protobuf:"varint,2,opt,name=verified" json:"verified,omitempty"`
}

func (m *AuthResp) Reset()                    { *m = AuthResp{} }
func (m *AuthResp) String() string            { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()               {}
func (*AuthResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResp) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

type TopicReq struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
	Topic    string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
}

func (m *TopicReq) Reset()                    { *m = TopicReq{} }
func (m *TopicReq) String() string            { return proto.CompactTextString(m) }
func (*TopicReq) ProtoMessage()               {}
func (*TopicReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TopicReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *TopicReq) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type TopicResp struct {
	Allow bool `protobuf:"varint,1,opt,name=allow" json:"allow,omitempty"`
}

func (m *TopicResp) Reset()                    { *m = TopicResp{} }
func (m *TopicResp) String() string            { return proto.CompactTextString(m) }
func (*TopicResp) ProtoMessage()               {}
func (*TopicResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TopicResp) GetAllow() bool {
	if m != nil {
		return m.Allow
	}
	return false
}

func init() {
	proto.RegisterType((*AuthReq)(nil), "auth.AuthReq")
	proto.RegisterType((*AuthResp)(nil), "auth.AuthResp")
	proto.RegisterType((*TopicReq)(nil), "auth.TopicReq")
	proto.RegisterType((*TopicResp)(nil), "auth.TopicResp")
	proto.RegisterEnum("auth.METHOD", METHOD_name, METHOD_value)
}

func init() { proto.RegisterFile("api/client/auth/proto/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x31, 0x4f, 0x86, 0x30,
	0x10, 0x40, 0xe5, 0x13, 0x3f, 0xca, 0xb9, 0x90, 0x8b, 0x03, 0x61, 0x42, 0x26, 0x63, 0x8c, 0x1d,
	0x5c, 0x5d, 0x4c, 0x34, 0xd1, 0xc1, 0x68, 0x2a, 0x7f, 0xa0, 0x40, 0x0d, 0x8d, 0x48, 0x2b, 0x16,
	0xf9, 0xfb, 0xa6, 0xbd, 0x86, 0xed, 0xbd, 0x26, 0xef, 0xee, 0x52, 0xa8, 0xa5, 0xd5, 0xbc, 0x9f,
	0xb4, 0x9a, 0x1d, 0x97, 0xab, 0x1b, 0xb9, 0x5d, 0x8c, 0x33, 0x01, 0x6f, 0x03, 0x62, 0xea, 0xb9,
	0xe1, 0x90, 0x3d, 0xac, 0x6e, 0x14, 0xea, 0x07, 0x11, 0xd2, 0x59, 0x7e, 0xab, 0x32, 0xa9, 0x93,
	0xab, 0x5c, 0x04, 0xc6, 0x02, 0x4e, 0xed, 0x36, 0x94, 0x87, 0xf0, 0xe4, 0xb1, 0xb9, 0x07, 0x46,
	0xc1, 0xaf, 0xc5, 0x0b, 0x38, 0x73, 0xe6, 0x4b, 0xcd, 0x31, 0x21, 0xc1, 0x0a, 0xd8, 0x9f, 0x5a,
	0xf4, 0xa7, 0x56, 0x14, 0x32, 0xb1, 0xbb, 0xaf, 0x5b, 0x63, 0x75, 0xef, 0xf7, 0x55, 0xc0, 0xe8,
	0xc0, 0x97, 0x21, 0x0e, 0xd8, 0x9d, 0x26, 0x5b, 0xdd, 0xc7, 0xcd, 0x24, 0xcd, 0x25, 0xe4, 0xb1,
	0xa6, 0xe5, 0x72, 0x9a, 0xcc, 0x16, 0x5a, 0x26, 0x48, 0xae, 0x6f, 0xe0, 0xf8, 0xfa, 0xd4, 0x3e,
	0xbf, 0x3d, 0x22, 0x83, 0xd4, 0x1f, 0x5a, 0x9c, 0xe0, 0x39, 0x64, 0x1f, 0x6b, 0x17, 0x24, 0xf1,
	0xf2, 0x1e, 0xe5, 0xd0, 0x1d, 0xc3, 0x57, 0xdc, 0xfd, 0x07, 0x00, 0x00, 0xff, 0xff, 0xae, 0x97,
	0x3b, 0x3a, 0x2e, 0x01, 0x00, 0x00,
}