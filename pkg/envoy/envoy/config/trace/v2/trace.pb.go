// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _struct "github.com/golang/protobuf/ptypes/struct"
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

// The tracing configuration specifies global
// settings for the HTTP tracer used by Envoy. The configuration is defined by
// the :ref:`Bootstrap <envoy_api_msg_config.bootstrap.v2.Bootstrap>` :ref:`tracing
// <envoy_api_field_config.bootstrap.v2.Bootstrap.tracing>` field. Envoy may support other tracers
// in the future, but right now the HTTP tracer is the only one supported.
type Tracing struct {
	// Provides configuration for the HTTP tracer.
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{0}
}
func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (dst *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(dst, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	// The name of the HTTP trace driver to instantiate. The name must match a
	// supported HTTP trace driver. *envoy.lightstep*, *envoy.zipkin*, and
	// *envoy.dynamic.ot* are built-in trace drivers.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being
	// instantiated. See the :ref:`LightstepConfig
	// <envoy_api_msg_config.trace.v2.LightstepConfig>`, :ref:`ZipkinConfig
	// <envoy_api_msg_config.trace.v2.ZipkinConfig>`, and :ref:`DynamicOtConfig
	// <envoy_api_msg_config.trace.v2.DynamicOtConfig>` trace drivers for examples.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{0, 0}
}
func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (dst *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(dst, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tracing_Http) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <http://lightstep.com/>`_ API.
	AccessTokenFile      string   `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile" json:"access_token_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{1}
}
func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (dst *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(dst, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	// The cluster manager cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster" json:"collector_cluster,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint    string   `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint" json:"collector_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{2}
}
func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (dst *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(dst, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
type DynamicOtConfig struct {
	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{3}
}
func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (dst *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(dst, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration structure.
type TraceServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_03ceb9ff1c87ef73, []int{4}
}
func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (dst *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(dst, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_trace_03ceb9ff1c87ef73)
}

var fileDescriptor_trace_03ceb9ff1c87ef73 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbd, 0x6e, 0x14, 0x31,
	0x14, 0x85, 0x35, 0xcb, 0x90, 0x28, 0x4e, 0xa4, 0x4d, 0x2c, 0xa1, 0x44, 0x2b, 0x40, 0x61, 0x43,
	0x41, 0x65, 0x4b, 0x8b, 0xf8, 0xa9, 0x13, 0xfe, 0x0a, 0x10, 0xd2, 0x24, 0xa2, 0x48, 0x33, 0xf2,
	0x3a, 0x77, 0x1d, 0x2b, 0x8e, 0x6d, 0x79, 0xee, 0x8e, 0xb4, 0x1d, 0x35, 0x8f, 0xc0, 0xa3, 0x50,
	0xf1, 0x3a, 0xbc, 0x05, 0xf2, 0xcf, 0x90, 0xa0, 0xa5, 0x41, 0xe9, 0x6c, 0xdf, 0x73, 0xee, 0xf9,
	0x34, 0x73, 0xc8, 0x13, 0xb0, 0xbd, 0x5b, 0x71, 0xe9, 0xec, 0x42, 0x2b, 0x8e, 0x41, 0x48, 0xe0,
	0xfd, 0x2c, 0x1f, 0x98, 0x0f, 0x0e, 0x1d, 0x7d, 0x90, 0x24, 0x2c, 0x4b, 0x58, 0x9e, 0xf4, 0xb3,
	0xc9, 0xd3, 0xec, 0x14, 0x5e, 0x47, 0x83, 0x74, 0x01, 0xb8, 0x0a, 0x5e, 0xb6, 0x1d, 0x84, 0x5e,
	0x0f, 0xe6, 0xc9, 0x43, 0xe5, 0x9c, 0x32, 0xc0, 0xd3, 0x6d, 0xbe, 0x5c, 0xf0, 0x0e, 0xc3, 0x52,
	0x62, 0x99, 0xee, 0xf7, 0xc2, 0xe8, 0x0b, 0x81, 0xc0, 0x87, 0x43, 0x1e, 0x4c, 0xbf, 0x57, 0x64,
	0xf3, 0x2c, 0x08, 0xa9, 0xad, 0xa2, 0xaf, 0x48, 0x7d, 0x89, 0xe8, 0x0f, 0xaa, 0xc3, 0xea, 0xd9,
	0xf6, 0xec, 0x88, 0xfd, 0x13, 0x87, 0x15, 0x35, 0xfb, 0x80, 0xe8, 0x9b, 0x64, 0x98, 0x7c, 0x21,
	0x75, 0xbc, 0xd1, 0x47, 0xa4, 0xb6, 0xe2, 0x1a, 0xd2, 0x82, 0xad, 0xe3, 0xad, 0x1f, 0xbf, 0x7e,
	0xde, 0xab, 0xc3, 0xe8, 0xb0, 0x6a, 0xd2, 0x33, 0xe5, 0x64, 0x23, 0x2f, 0x3b, 0x18, 0xa5, 0x84,
	0x7d, 0x96, 0x99, 0xd9, 0xc0, 0xcc, 0x4e, 0x13, 0x73, 0x53, 0x64, 0xd3, 0xaf, 0x15, 0x19, 0x7f,
	0xd4, 0xea, 0x12, 0x3b, 0x04, 0x7f, 0x92, 0xde, 0xe8, 0x4b, 0xb2, 0x27, 0x9d, 0x31, 0x20, 0xd1,
	0x85, 0x56, 0x9a, 0x65, 0x87, 0x10, 0xd6, 0x03, 0x77, 0xff, 0x68, 0x4e, 0xb2, 0x84, 0xbe, 0x20,
	0x7b, 0x42, 0x4a, 0xe8, 0xba, 0x16, 0xdd, 0x15, 0xd8, 0x76, 0xa1, 0x0d, 0x24, 0x8e, 0xbf, 0x7c,
	0xe3, 0xac, 0x39, 0x8b, 0x92, 0x77, 0xda, 0x40, 0x44, 0xd8, 0x39, 0xd7, 0xfe, 0x4a, 0xdb, 0x3b,
	0xe6, 0xbf, 0x26, 0xf4, 0xc6, 0x07, 0xf6, 0xc2, 0x3b, 0x6d, 0x71, 0x1d, 0xe0, 0x66, 0xf9, 0xdb,
	0xa2, 0x99, 0x2a, 0x32, 0x7e, 0xb3, 0xb2, 0xe2, 0x5a, 0xcb, 0xcf, 0x58, 0x20, 0x8e, 0xc8, 0xa6,
	0xd1, 0xf3, 0x20, 0xc2, 0x6a, 0x3d, 0x7a, 0x98, 0xfc, 0xff, 0xe7, 0x96, 0x84, 0xc6, 0x9f, 0x0b,
	0xa7, 0xb9, 0x58, 0x25, 0xeb, 0x13, 0xd9, 0xb9, 0x5d, 0xb7, 0xd2, 0x8e, 0xc7, 0xa5, 0x1d, 0xc2,
	0xeb, 0x58, 0x8a, 0xd8, 0x4a, 0xf6, 0x3e, 0x78, 0x59, 0xbc, 0xc7, 0x24, 0x02, 0xdd, 0xff, 0x56,
	0x8d, 0x76, 0xab, 0x66, 0x5b, 0xdd, 0x1a, 0xd4, 0xe7, 0xa3, 0x7e, 0x36, 0xdf, 0x48, 0x0c, 0xcf,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x93, 0x3a, 0x40, 0x16, 0x03, 0x00, 0x00,
}
