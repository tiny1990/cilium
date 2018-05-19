// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/discovery.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/gogo/protobuf/gogoproto"
import any "github.com/golang/protobuf/ptypes/any"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A DiscoveryRequest requests a set of versioned resources of the same type for
// a given Envoy node on some API.
type DiscoveryRequest struct {
	// The version_info provided in the request messages will be the version_info
	// received with the most recent successfully processed response or empty on
	// the first request. It is expected that no new request is sent after a
	// response is received until the Envoy instance is ready to ACK/NACK the new
	// configuration. ACK/NACK takes place by returning the new API config version
	// as applied or the previous API config version respectively. Each type_url
	// (see below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	// The node making the request.
	Node *core.Node `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	// List of resources to subscribe to, e.g. list of cluster names or a route
	// configuration name. If this is empty, all resources for the API are
	// returned. LDS/CDS expect empty resource_names, since this is global
	// discovery for the Envoy instance. The LDS and CDS responses will then imply
	// a number of resources that need to be fetched via EDS/RDS, which will be
	// explicitly enumerated in resource_names.
	ResourceNames []string `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames" json:"resource_names,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	// nonce corresponding to DiscoveryResponse being ACK/NACKed. See above
	// discussion on version_info and the DiscoveryResponse nonce comment. This
	// may be empty if no nonce is available, e.g. at startup or for non-stream
	// xDS implementations.
	ResponseNonce string `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details* provides the Envoy
	// internal exception related to the failure. It is only intended for consumption during manual
	// debugging, the string provided is not guaranteed to be stable across Envoy versions.
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_b412db76c3d05c43, []int{0}
}
func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(dst, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	// The response resources. These resources are typed and depend on the API being called.
	Resources []*any.Any `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
	// [#not-implemented-hide:]
	// Canary is used to support two Envoy command line flags:
	//
	// * --terminate-on-canary-transition-failure. When set, Envoy is able to
	//   terminate if it detects that configuration is stuck at canary. Consider
	//   this example sequence of updates:
	//   - Management server applies a canary config successfully.
	//   - Management server rolls back to a production config.
	//   - Envoy rejects the new production config.
	//   Since there is no sensible way to continue receiving configuration
	//   updates, Envoy will then terminate and apply production config from a
	//   clean slate.
	// * --dry-run-canary. When set, a canary response will never be applied, only
	//   validated via a dry run.
	Canary bool `protobuf:"varint,3,opt,name=canary" json:"canary,omitempty"`
	// Type URL for resources. This must be consistent with the type_url in the
	// Any messages for resources if resources is non-empty. This effectively
	// identifies the xDS API when muxing over ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	// For gRPC based subscriptions, the nonce provides a way to explicitly ack a
	// specific DiscoveryResponse in a following DiscoveryRequest. Additional
	// messages may have been sent by Envoy to the management server for the
	// previous version on the stream prior to this DiscoveryResponse, that were
	// unprocessed at response send time. The nonce allows the management server
	// to ignore any further DiscoveryRequests for the previous version until a
	// DiscoveryRequest bearing the nonce. The nonce is optional and is not
	// required for non-stream based xDS implementations.
	Nonce                string   `protobuf:"bytes,5,opt,name=nonce" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_b412db76c3d05c43, []int{1}
}
func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(dst, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*any.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
}

func init() {
	proto.RegisterFile("envoy/api/v2/discovery.proto", fileDescriptor_discovery_b412db76c3d05c43)
}

var fileDescriptor_discovery_b412db76c3d05c43 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0xea, 0xd3, 0x40,
	0x10, 0xc6, 0x4d, 0xd2, 0x7f, 0x6d, 0x37, 0xad, 0xe8, 0x52, 0xec, 0xb6, 0x88, 0xc6, 0x82, 0x10,
	0x10, 0x76, 0x21, 0x22, 0x78, 0xb5, 0xf4, 0xe2, 0xa5, 0x87, 0x88, 0x17, 0x2f, 0x61, 0x9b, 0x4c,
	0x43, 0x20, 0xee, 0xc4, 0xdd, 0x24, 0x90, 0xab, 0x4f, 0xe3, 0x2b, 0xf8, 0x06, 0x3e, 0x85, 0x07,
	0x9f, 0x44, 0xb2, 0x49, 0xb0, 0x5e, 0xe4, 0x7f, 0xdb, 0xf9, 0xe6, 0xc7, 0xcc, 0x37, 0xdf, 0x92,
	0x67, 0xa0, 0x5a, 0xec, 0x84, 0xac, 0x0a, 0xd1, 0x46, 0x22, 0x2b, 0x4c, 0x8a, 0x2d, 0xe8, 0x8e,
	0x57, 0x1a, 0x6b, 0xa4, 0x2b, 0xdb, 0xe5, 0xb2, 0x2a, 0x78, 0x1b, 0xed, 0xff, 0x65, 0x53, 0xd4,
	0x20, 0x2e, 0xd2, 0xc0, 0xc0, 0xee, 0x77, 0x39, 0x62, 0x5e, 0x82, 0xb0, 0xd5, 0xa5, 0xb9, 0x0a,
	0xa9, 0xc6, 0x31, 0xfb, 0xed, 0xd8, 0xd2, 0x55, 0x2a, 0x4c, 0x2d, 0xeb, 0xc6, 0x8c, 0x8d, 0x4d,
	0x8e, 0x39, 0xda, 0xa7, 0xe8, 0x5f, 0x83, 0x7a, 0xf8, 0xe6, 0x92, 0xc7, 0xa7, 0xc9, 0x49, 0x0c,
	0x5f, 0x1b, 0x30, 0x35, 0x7d, 0x49, 0x56, 0x2d, 0x68, 0x53, 0xa0, 0x4a, 0x0a, 0x75, 0x45, 0xe6,
	0x04, 0x4e, 0xb8, 0x8c, 0xfd, 0x51, 0xfb, 0xa0, 0xae, 0x48, 0x5f, 0x93, 0x99, 0xc2, 0x0c, 0x98,
	0x1b, 0x38, 0xa1, 0x1f, 0x6d, 0xf9, 0xad, 0x79, 0xde, 0xdb, 0xe5, 0x67, 0xcc, 0x20, 0xb6, 0x10,
	0x7d, 0x45, 0x1e, 0x69, 0x30, 0xd8, 0xe8, 0x14, 0x12, 0x25, 0xbf, 0x80, 0x61, 0x5e, 0xe0, 0x85,
	0xcb, 0x78, 0x3d, 0xa9, 0xe7, 0x5e, 0xa4, 0x3b, 0xb2, 0xa8, 0xbb, 0x0a, 0x92, 0x46, 0x97, 0x6c,
	0x66, 0x57, 0x3e, 0xec, 0xeb, 0x4f, 0xba, 0x1c, 0x27, 0x54, 0xa8, 0x0c, 0x24, 0x0a, 0x55, 0x0a,
	0xec, 0xce, 0x02, 0xeb, 0x49, 0x3d, 0xf7, 0x22, 0x7d, 0x4b, 0x56, 0xa0, 0x35, 0xea, 0x24, 0x83,
	0x5a, 0x16, 0x25, 0x9b, 0x5b, 0x77, 0x94, 0x0f, 0x99, 0x70, 0x5d, 0xa5, 0xfc, 0xa3, 0xcd, 0x24,
	0xf6, 0x2d, 0x77, 0xb2, 0xd8, 0xe1, 0x87, 0x43, 0x9e, 0xdc, 0x84, 0x30, 0x4c, 0xbc, 0x4f, 0x0a,
	0xef, 0xc8, 0x72, 0x3a, 0xc1, 0x30, 0x37, 0xf0, 0x42, 0x3f, 0xda, 0x4c, 0xcb, 0xa6, 0xbf, 0xe1,
	0xef, 0x55, 0x77, 0x9c, 0xfd, 0xfc, 0xf5, 0xe2, 0x41, 0xfc, 0x17, 0xa6, 0x4f, 0xc9, 0x3c, 0x95,
	0x4a, 0xea, 0x8e, 0x79, 0x81, 0x13, 0x2e, 0xe2, 0xb1, 0xfa, 0x5f, 0x06, 0x1b, 0x72, 0x77, 0x7b,
	0xfa, 0x50, 0x1c, 0x17, 0xdf, 0x7f, 0x3f, 0x77, 0x3e, 0xbb, 0x6d, 0x74, 0x99, 0xdb, 0x8d, 0x6f,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xab, 0x9c, 0xf3, 0x67, 0x02, 0x00, 0x00,
}
