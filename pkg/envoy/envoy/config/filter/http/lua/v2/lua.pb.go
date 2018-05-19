// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v2/lua.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Lua struct {
	// The Lua code that Envoy will execute. This can be a very small script that
	// further loads code from disk if desired. Note that if JSON configuration is used, the code must
	// be properly escaped. YAML configuration may be easier to read since YAML supports multi-line
	// strings so complex scripts can be easily expressed inline in the configuration.
	InlineCode           string   `protobuf:"bytes,1,opt,name=inline_code,json=inlineCode" json:"inline_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lua) Reset()         { *m = Lua{} }
func (m *Lua) String() string { return proto.CompactTextString(m) }
func (*Lua) ProtoMessage()    {}
func (*Lua) Descriptor() ([]byte, []int) {
	return fileDescriptor_lua_9c1a4a0b22b06723, []int{0}
}
func (m *Lua) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lua.Unmarshal(m, b)
}
func (m *Lua) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lua.Marshal(b, m, deterministic)
}
func (dst *Lua) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lua.Merge(dst, src)
}
func (m *Lua) XXX_Size() int {
	return xxx_messageInfo_Lua.Size(m)
}
func (m *Lua) XXX_DiscardUnknown() {
	xxx_messageInfo_Lua.DiscardUnknown(m)
}

var xxx_messageInfo_Lua proto.InternalMessageInfo

func (m *Lua) GetInlineCode() string {
	if m != nil {
		return m.InlineCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Lua)(nil), "envoy.config.filter.http.lua.v2.Lua")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/lua/v2/lua.proto", fileDescriptor_lua_9c1a4a0b22b06723)
}

var fileDescriptor_lua_9c1a4a0b22b06723 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x29, 0x4d, 0xd4, 0x2f, 0x33, 0x02, 0x51, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xf2, 0x60, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x20, 0xa5,
	0x7a, 0x20, 0x35, 0x65, 0x46, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa,
	0x30, 0x06, 0x44, 0xa7, 0x92, 0x21, 0x17, 0xb3, 0x4f, 0x69, 0xa2, 0x90, 0x16, 0x17, 0x77, 0x66,
	0x5e, 0x4e, 0x66, 0x5e, 0x6a, 0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7,
	0x13, 0xe7, 0xae, 0x97, 0x07, 0x98, 0x59, 0x8a, 0x98, 0x14, 0x18, 0x83, 0xb8, 0x20, 0xb2, 0xce,
	0xf9, 0x29, 0xa9, 0x4e, 0x2c, 0x51, 0x4c, 0x65, 0x46, 0x49, 0x6c, 0x60, 0xfd, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0x2c, 0x11, 0xa8, 0xa6, 0x00, 0x00, 0x00,
}
