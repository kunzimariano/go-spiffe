// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workload.proto

package workload

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
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

type X509SVIDRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X509SVIDRequest) Reset()         { *m = X509SVIDRequest{} }
func (m *X509SVIDRequest) String() string { return proto.CompactTextString(m) }
func (*X509SVIDRequest) ProtoMessage()    {}
func (*X509SVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{0}
}

func (m *X509SVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509SVIDRequest.Unmarshal(m, b)
}
func (m *X509SVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509SVIDRequest.Marshal(b, m, deterministic)
}
func (m *X509SVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509SVIDRequest.Merge(m, src)
}
func (m *X509SVIDRequest) XXX_Size() int {
	return xxx_messageInfo_X509SVIDRequest.Size(m)
}
func (m *X509SVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_X509SVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_X509SVIDRequest proto.InternalMessageInfo

// The X509SVIDResponse message carries a set of X.509 SVIDs and their
// associated information. It also carries a set of global CRLs, and a
// TTL to inform the workload when it should check back next.
type X509SVIDResponse struct {
	// A list of X509SVID messages, each of which includes a single
	// SPIFFE Verifiable Identity Document, along with its private key
	// and bundle.
	Svids []*X509SVID `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
	// ASN.1 DER encoded
	Crl [][]byte `protobuf:"bytes,2,rep,name=crl,proto3" json:"crl,omitempty"`
	// CA certificate bundles belonging to foreign Trust Domains that the
	// workload should trust, keyed by the SPIFFE ID of the foreign
	// domain. Bundles are ASN.1 DER encoded.
	FederatedBundles     map[string][]byte `protobuf:"bytes,3,rep,name=federated_bundles,json=federatedBundles,proto3" json:"federated_bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *X509SVIDResponse) Reset()         { *m = X509SVIDResponse{} }
func (m *X509SVIDResponse) String() string { return proto.CompactTextString(m) }
func (*X509SVIDResponse) ProtoMessage()    {}
func (*X509SVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{1}
}

func (m *X509SVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509SVIDResponse.Unmarshal(m, b)
}
func (m *X509SVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509SVIDResponse.Marshal(b, m, deterministic)
}
func (m *X509SVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509SVIDResponse.Merge(m, src)
}
func (m *X509SVIDResponse) XXX_Size() int {
	return xxx_messageInfo_X509SVIDResponse.Size(m)
}
func (m *X509SVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_X509SVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_X509SVIDResponse proto.InternalMessageInfo

func (m *X509SVIDResponse) GetSvids() []*X509SVID {
	if m != nil {
		return m.Svids
	}
	return nil
}

func (m *X509SVIDResponse) GetCrl() [][]byte {
	if m != nil {
		return m.Crl
	}
	return nil
}

func (m *X509SVIDResponse) GetFederatedBundles() map[string][]byte {
	if m != nil {
		return m.FederatedBundles
	}
	return nil
}

// The X509SVID message carries a single SVID and all associated
// information, including CA bundles.
type X509SVID struct {
	// The SPIFFE ID of the SVID in this entry
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// ASN.1 DER encoded certificate chain. MAY include intermediates,
	// the leaf certificate (or SVID itself) MUST come first.
	X509Svid []byte `protobuf:"bytes,2,opt,name=x509_svid,json=x509Svid,proto3" json:"x509_svid,omitempty"`
	// ASN.1 DER encoded PKCS#8 private key. MUST be unencrypted.
	X509SvidKey []byte `protobuf:"bytes,3,opt,name=x509_svid_key,json=x509SvidKey,proto3" json:"x509_svid_key,omitempty"`
	// CA certificates belonging to the Trust Domain
	// ASN.1 DER encoded
	Bundle []byte `protobuf:"bytes,4,opt,name=bundle,proto3" json:"bundle,omitempty"`
	// List of trust domains the SVID federates with, which corresponds to
	// keys in the federated_bundles map in the X509SVIDResponse message.
	FederatesWith        []string `protobuf:"bytes,5,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X509SVID) Reset()         { *m = X509SVID{} }
func (m *X509SVID) String() string { return proto.CompactTextString(m) }
func (*X509SVID) ProtoMessage()    {}
func (*X509SVID) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{2}
}

func (m *X509SVID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509SVID.Unmarshal(m, b)
}
func (m *X509SVID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509SVID.Marshal(b, m, deterministic)
}
func (m *X509SVID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509SVID.Merge(m, src)
}
func (m *X509SVID) XXX_Size() int {
	return xxx_messageInfo_X509SVID.Size(m)
}
func (m *X509SVID) XXX_DiscardUnknown() {
	xxx_messageInfo_X509SVID.DiscardUnknown(m)
}

var xxx_messageInfo_X509SVID proto.InternalMessageInfo

func (m *X509SVID) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *X509SVID) GetX509Svid() []byte {
	if m != nil {
		return m.X509Svid
	}
	return nil
}

func (m *X509SVID) GetX509SvidKey() []byte {
	if m != nil {
		return m.X509SvidKey
	}
	return nil
}

func (m *X509SVID) GetBundle() []byte {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *X509SVID) GetFederatesWith() []string {
	if m != nil {
		return m.FederatesWith
	}
	return nil
}

type JWTSVID struct {
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Encoded using JWS Compact Serialization
	Svid                 string   `protobuf:"bytes,2,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTSVID) Reset()         { *m = JWTSVID{} }
func (m *JWTSVID) String() string { return proto.CompactTextString(m) }
func (*JWTSVID) ProtoMessage()    {}
func (*JWTSVID) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{3}
}

func (m *JWTSVID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTSVID.Unmarshal(m, b)
}
func (m *JWTSVID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTSVID.Marshal(b, m, deterministic)
}
func (m *JWTSVID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTSVID.Merge(m, src)
}
func (m *JWTSVID) XXX_Size() int {
	return xxx_messageInfo_JWTSVID.Size(m)
}
func (m *JWTSVID) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTSVID.DiscardUnknown(m)
}

var xxx_messageInfo_JWTSVID proto.InternalMessageInfo

func (m *JWTSVID) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *JWTSVID) GetSvid() string {
	if m != nil {
		return m.Svid
	}
	return ""
}

type JWTSVIDRequest struct {
	Audience []string `protobuf:"bytes,1,rep,name=audience,proto3" json:"audience,omitempty"`
	// SPIFFE ID of the JWT being requested
	// If not set, all IDs will be returned
	SpiffeId             string   `protobuf:"bytes,2,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTSVIDRequest) Reset()         { *m = JWTSVIDRequest{} }
func (m *JWTSVIDRequest) String() string { return proto.CompactTextString(m) }
func (*JWTSVIDRequest) ProtoMessage()    {}
func (*JWTSVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{4}
}

func (m *JWTSVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTSVIDRequest.Unmarshal(m, b)
}
func (m *JWTSVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTSVIDRequest.Marshal(b, m, deterministic)
}
func (m *JWTSVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTSVIDRequest.Merge(m, src)
}
func (m *JWTSVIDRequest) XXX_Size() int {
	return xxx_messageInfo_JWTSVIDRequest.Size(m)
}
func (m *JWTSVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTSVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JWTSVIDRequest proto.InternalMessageInfo

func (m *JWTSVIDRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *JWTSVIDRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

type JWTSVIDResponse struct {
	Svids                []*JWTSVID `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *JWTSVIDResponse) Reset()         { *m = JWTSVIDResponse{} }
func (m *JWTSVIDResponse) String() string { return proto.CompactTextString(m) }
func (*JWTSVIDResponse) ProtoMessage()    {}
func (*JWTSVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{5}
}

func (m *JWTSVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTSVIDResponse.Unmarshal(m, b)
}
func (m *JWTSVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTSVIDResponse.Marshal(b, m, deterministic)
}
func (m *JWTSVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTSVIDResponse.Merge(m, src)
}
func (m *JWTSVIDResponse) XXX_Size() int {
	return xxx_messageInfo_JWTSVIDResponse.Size(m)
}
func (m *JWTSVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTSVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JWTSVIDResponse proto.InternalMessageInfo

func (m *JWTSVIDResponse) GetSvids() []*JWTSVID {
	if m != nil {
		return m.Svids
	}
	return nil
}

type JWTBundlesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTBundlesRequest) Reset()         { *m = JWTBundlesRequest{} }
func (m *JWTBundlesRequest) String() string { return proto.CompactTextString(m) }
func (*JWTBundlesRequest) ProtoMessage()    {}
func (*JWTBundlesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{6}
}

func (m *JWTBundlesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTBundlesRequest.Unmarshal(m, b)
}
func (m *JWTBundlesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTBundlesRequest.Marshal(b, m, deterministic)
}
func (m *JWTBundlesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTBundlesRequest.Merge(m, src)
}
func (m *JWTBundlesRequest) XXX_Size() int {
	return xxx_messageInfo_JWTBundlesRequest.Size(m)
}
func (m *JWTBundlesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTBundlesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JWTBundlesRequest proto.InternalMessageInfo

type JWTBundlesResponse struct {
	// JWK sets, keyed by trust domain URI
	Bundles              map[string][]byte `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JWTBundlesResponse) Reset()         { *m = JWTBundlesResponse{} }
func (m *JWTBundlesResponse) String() string { return proto.CompactTextString(m) }
func (*JWTBundlesResponse) ProtoMessage()    {}
func (*JWTBundlesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{7}
}

func (m *JWTBundlesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTBundlesResponse.Unmarshal(m, b)
}
func (m *JWTBundlesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTBundlesResponse.Marshal(b, m, deterministic)
}
func (m *JWTBundlesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTBundlesResponse.Merge(m, src)
}
func (m *JWTBundlesResponse) XXX_Size() int {
	return xxx_messageInfo_JWTBundlesResponse.Size(m)
}
func (m *JWTBundlesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTBundlesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JWTBundlesResponse proto.InternalMessageInfo

func (m *JWTBundlesResponse) GetBundles() map[string][]byte {
	if m != nil {
		return m.Bundles
	}
	return nil
}

type ValidateJWTSVIDRequest struct {
	Audience string `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	// Encoded using JWS Compact Serialization
	Svid                 string   `protobuf:"bytes,2,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateJWTSVIDRequest) Reset()         { *m = ValidateJWTSVIDRequest{} }
func (m *ValidateJWTSVIDRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateJWTSVIDRequest) ProtoMessage()    {}
func (*ValidateJWTSVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{8}
}

func (m *ValidateJWTSVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateJWTSVIDRequest.Unmarshal(m, b)
}
func (m *ValidateJWTSVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateJWTSVIDRequest.Marshal(b, m, deterministic)
}
func (m *ValidateJWTSVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateJWTSVIDRequest.Merge(m, src)
}
func (m *ValidateJWTSVIDRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateJWTSVIDRequest.Size(m)
}
func (m *ValidateJWTSVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateJWTSVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateJWTSVIDRequest proto.InternalMessageInfo

func (m *ValidateJWTSVIDRequest) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *ValidateJWTSVIDRequest) GetSvid() string {
	if m != nil {
		return m.Svid
	}
	return ""
}

type ValidateJWTSVIDResponse struct {
	SpiffeId             string          `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	Claims               *_struct.Struct `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ValidateJWTSVIDResponse) Reset()         { *m = ValidateJWTSVIDResponse{} }
func (m *ValidateJWTSVIDResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateJWTSVIDResponse) ProtoMessage()    {}
func (*ValidateJWTSVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_611edb31abe0f206, []int{9}
}

func (m *ValidateJWTSVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateJWTSVIDResponse.Unmarshal(m, b)
}
func (m *ValidateJWTSVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateJWTSVIDResponse.Marshal(b, m, deterministic)
}
func (m *ValidateJWTSVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateJWTSVIDResponse.Merge(m, src)
}
func (m *ValidateJWTSVIDResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateJWTSVIDResponse.Size(m)
}
func (m *ValidateJWTSVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateJWTSVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateJWTSVIDResponse proto.InternalMessageInfo

func (m *ValidateJWTSVIDResponse) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *ValidateJWTSVIDResponse) GetClaims() *_struct.Struct {
	if m != nil {
		return m.Claims
	}
	return nil
}

func init() {
	proto.RegisterType((*X509SVIDRequest)(nil), "X509SVIDRequest")
	proto.RegisterType((*X509SVIDResponse)(nil), "X509SVIDResponse")
	proto.RegisterMapType((map[string][]byte)(nil), "X509SVIDResponse.FederatedBundlesEntry")
	proto.RegisterType((*X509SVID)(nil), "X509SVID")
	proto.RegisterType((*JWTSVID)(nil), "JWTSVID")
	proto.RegisterType((*JWTSVIDRequest)(nil), "JWTSVIDRequest")
	proto.RegisterType((*JWTSVIDResponse)(nil), "JWTSVIDResponse")
	proto.RegisterType((*JWTBundlesRequest)(nil), "JWTBundlesRequest")
	proto.RegisterType((*JWTBundlesResponse)(nil), "JWTBundlesResponse")
	proto.RegisterMapType((map[string][]byte)(nil), "JWTBundlesResponse.BundlesEntry")
	proto.RegisterType((*ValidateJWTSVIDRequest)(nil), "ValidateJWTSVIDRequest")
	proto.RegisterType((*ValidateJWTSVIDResponse)(nil), "ValidateJWTSVIDResponse")
}

func init() { proto.RegisterFile("workload.proto", fileDescriptor_611edb31abe0f206) }

var fileDescriptor_611edb31abe0f206 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8f, 0xd2, 0x4c,
	0x14, 0xce, 0xd0, 0x85, 0xa5, 0x87, 0xef, 0xd9, 0xf7, 0x5d, 0x9a, 0x6a, 0xb4, 0x69, 0x62, 0xe4,
	0x6a, 0x60, 0x31, 0x6b, 0x5c, 0xe2, 0x8d, 0xba, 0x6e, 0x64, 0xbd, 0x31, 0x85, 0x2c, 0xde, 0x91,
	0x42, 0x07, 0x68, 0xb6, 0x52, 0x6c, 0xa7, 0xac, 0xdc, 0x7a, 0xed, 0xff, 0xf0, 0x6f, 0xf9, 0x53,
	0x4c, 0xa7, 0x33, 0xd5, 0x16, 0x5c, 0x13, 0xef, 0xe6, 0x3c, 0xe7, 0x39, 0x4f, 0xcf, 0x67, 0xa1,
	0x7e, 0xe7, 0x07, 0xb7, 0x9e, 0x6f, 0x3b, 0x64, 0x13, 0xf8, 0xcc, 0xd7, 0x1f, 0x2e, 0x7d, 0x7f,
	0xe9, 0xd1, 0x2e, 0xb7, 0x66, 0xd1, 0xa2, 0x1b, 0xb2, 0x20, 0x9a, 0xb3, 0xc4, 0x6b, 0xb6, 0xa0,
	0xf1, 0xf1, 0xbc, 0x77, 0x31, 0xba, 0x19, 0x5e, 0x5a, 0xf4, 0x73, 0x44, 0x43, 0x66, 0xfe, 0x40,
	0xd0, 0xfc, 0x85, 0x85, 0x1b, 0x7f, 0x1d, 0x52, 0xfc, 0x18, 0x8a, 0xe1, 0xd6, 0x75, 0x42, 0x0d,
	0x19, 0x4a, 0xa7, 0xd2, 0x57, 0x49, 0xca, 0x48, 0x70, 0xdc, 0x04, 0x65, 0x1e, 0x78, 0x5a, 0xc1,
	0x50, 0x3a, 0x55, 0x2b, 0x7e, 0xe2, 0x31, 0xb4, 0x16, 0xd4, 0xa1, 0x81, 0xcd, 0xa8, 0x33, 0x9d,
	0x45, 0x6b, 0xc7, 0xa3, 0xa1, 0xa6, 0xf0, 0xf0, 0xa7, 0x24, 0xff, 0x01, 0x72, 0x25, 0xa9, 0xaf,
	0x13, 0xe6, 0xdb, 0x35, 0x0b, 0x76, 0x56, 0x73, 0x91, 0x83, 0xf5, 0x37, 0xf0, 0xff, 0x41, 0x6a,
	0x9c, 0xc0, 0x2d, 0xdd, 0x69, 0xc8, 0x40, 0x1d, 0xd5, 0x8a, 0x9f, 0xf8, 0x3f, 0x28, 0x6e, 0x6d,
	0x2f, 0xa2, 0x5a, 0xc1, 0x40, 0x9d, 0xaa, 0x95, 0x18, 0x83, 0xc2, 0x0b, 0x64, 0x7e, 0x47, 0x50,
	0x96, 0x19, 0xe0, 0x07, 0xa0, 0x86, 0x1b, 0x77, 0xb1, 0xa0, 0x53, 0xd7, 0x11, 0xe1, 0xe5, 0x04,
	0x18, 0x3a, 0xb1, 0xf3, 0xcb, 0x79, 0xef, 0x62, 0x1a, 0x17, 0x29, 0x74, 0xca, 0x31, 0x30, 0xda,
	0xba, 0x0e, 0x36, 0xa1, 0x96, 0x3a, 0xa7, 0xf1, 0xc7, 0x15, 0x4e, 0xa8, 0x48, 0xc2, 0x7b, 0xba,
	0xc3, 0xa7, 0x50, 0x4a, 0x6a, 0xd7, 0x8e, 0xb8, 0x53, 0x58, 0xf8, 0x09, 0xd4, 0x65, 0x6d, 0xe1,
	0xf4, 0xce, 0x65, 0x2b, 0xad, 0x68, 0x28, 0x1d, 0xd5, 0xaa, 0xa5, 0xe8, 0xc4, 0x65, 0x2b, 0x73,
	0x00, 0xc7, 0xd7, 0x93, 0xf1, 0xdf, 0xf3, 0xc4, 0x70, 0x94, 0xa6, 0xa8, 0x5a, 0xfc, 0x6d, 0x0e,
	0xa1, 0x2e, 0x62, 0xc5, 0x68, 0xb1, 0x0e, 0x65, 0x3b, 0x72, 0x5c, 0xba, 0x9e, 0x53, 0x3e, 0x48,
	0xd5, 0x4a, 0xed, 0xac, 0x7c, 0x21, 0x2b, 0x6f, 0x9e, 0x41, 0x23, 0x95, 0x12, 0x1b, 0xf1, 0x28,
	0xbb, 0x11, 0x65, 0x22, 0x09, 0x09, 0x6c, 0x9e, 0x40, 0xeb, 0x7a, 0x32, 0x16, 0x23, 0x92, 0xbb,
	0xf5, 0x0d, 0x01, 0xfe, 0x1d, 0x15, 0x5a, 0x03, 0x38, 0x96, 0x0b, 0x92, 0xa8, 0x19, 0x64, 0x9f,
	0x45, 0x32, 0x9b, 0x21, 0x03, 0xf4, 0x01, 0x54, 0xff, 0x79, 0x0f, 0xde, 0xc1, 0xe9, 0x8d, 0xed,
	0xb9, 0x8e, 0xcd, 0xe8, 0xbd, 0x9d, 0x42, 0x99, 0x4e, 0x1d, 0xea, 0xf5, 0x12, 0xda, 0x7b, 0x4a,
	0xa2, 0xb8, 0x7b, 0xe7, 0xd6, 0x85, 0xd2, 0xdc, 0xb3, 0xdd, 0x4f, 0x21, 0x57, 0xab, 0xf4, 0xdb,
	0x24, 0x39, 0x57, 0x22, 0xcf, 0x95, 0x8c, 0xf8, 0xb9, 0x5a, 0x82, 0xd6, 0xff, 0x5a, 0x80, 0xd6,
	0x88, 0x47, 0x4f, 0xc4, 0x9d, 0xbf, 0xfa, 0x30, 0xc4, 0x67, 0x50, 0xbd, 0xa2, 0x6c, 0xbe, 0x92,
	0xbb, 0xd2, 0x20, 0xd9, 0x7a, 0xf4, 0x26, 0xc9, 0xa7, 0xf5, 0x12, 0x1a, 0x32, 0x44, 0xf4, 0x0f,
	0x63, 0xb2, 0x37, 0x31, 0xfd, 0xe4, 0xc0, 0x24, 0x7a, 0x08, 0x5f, 0x42, 0x23, 0x57, 0x2f, 0x6e,
	0x93, 0xc3, 0xbd, 0xd4, 0x35, 0xf2, 0xa7, 0xd6, 0x3c, 0x87, 0x1a, 0xcf, 0x21, 0xbd, 0xc5, 0x26,
	0xc9, 0xfd, 0x8d, 0xf4, 0xd6, 0xde, 0xaf, 0xa2, 0x87, 0x66, 0x25, 0xde, 0x9d, 0x67, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0x6b, 0xd6, 0x93, 0xec, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpiffeWorkloadAPIClient is the client API for SpiffeWorkloadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpiffeWorkloadAPIClient interface {
	// JWT-SVID Profile
	FetchJWTSVID(ctx context.Context, in *JWTSVIDRequest, opts ...grpc.CallOption) (*JWTSVIDResponse, error)
	FetchJWTBundles(ctx context.Context, in *JWTBundlesRequest, opts ...grpc.CallOption) (SpiffeWorkloadAPI_FetchJWTBundlesClient, error)
	ValidateJWTSVID(ctx context.Context, in *ValidateJWTSVIDRequest, opts ...grpc.CallOption) (*ValidateJWTSVIDResponse, error)
	// X.509-SVID Profile
	// Fetch all SPIFFE identities the workload is entitled to, as
	// well as related information like trust bundles and CRLs. As
	// this information changes, subsequent messages will be sent.
	FetchX509SVID(ctx context.Context, in *X509SVIDRequest, opts ...grpc.CallOption) (SpiffeWorkloadAPI_FetchX509SVIDClient, error)
}

type spiffeWorkloadAPIClient struct {
	cc *grpc.ClientConn
}

func NewSpiffeWorkloadAPIClient(cc *grpc.ClientConn) SpiffeWorkloadAPIClient {
	return &spiffeWorkloadAPIClient{cc}
}

func (c *spiffeWorkloadAPIClient) FetchJWTSVID(ctx context.Context, in *JWTSVIDRequest, opts ...grpc.CallOption) (*JWTSVIDResponse, error) {
	out := new(JWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/SpiffeWorkloadAPI/FetchJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiffeWorkloadAPIClient) FetchJWTBundles(ctx context.Context, in *JWTBundlesRequest, opts ...grpc.CallOption) (SpiffeWorkloadAPI_FetchJWTBundlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SpiffeWorkloadAPI_serviceDesc.Streams[0], "/SpiffeWorkloadAPI/FetchJWTBundles", opts...)
	if err != nil {
		return nil, err
	}
	x := &spiffeWorkloadAPIFetchJWTBundlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpiffeWorkloadAPI_FetchJWTBundlesClient interface {
	Recv() (*JWTBundlesResponse, error)
	grpc.ClientStream
}

type spiffeWorkloadAPIFetchJWTBundlesClient struct {
	grpc.ClientStream
}

func (x *spiffeWorkloadAPIFetchJWTBundlesClient) Recv() (*JWTBundlesResponse, error) {
	m := new(JWTBundlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *spiffeWorkloadAPIClient) ValidateJWTSVID(ctx context.Context, in *ValidateJWTSVIDRequest, opts ...grpc.CallOption) (*ValidateJWTSVIDResponse, error) {
	out := new(ValidateJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/SpiffeWorkloadAPI/ValidateJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiffeWorkloadAPIClient) FetchX509SVID(ctx context.Context, in *X509SVIDRequest, opts ...grpc.CallOption) (SpiffeWorkloadAPI_FetchX509SVIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SpiffeWorkloadAPI_serviceDesc.Streams[1], "/SpiffeWorkloadAPI/FetchX509SVID", opts...)
	if err != nil {
		return nil, err
	}
	x := &spiffeWorkloadAPIFetchX509SVIDClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpiffeWorkloadAPI_FetchX509SVIDClient interface {
	Recv() (*X509SVIDResponse, error)
	grpc.ClientStream
}

type spiffeWorkloadAPIFetchX509SVIDClient struct {
	grpc.ClientStream
}

func (x *spiffeWorkloadAPIFetchX509SVIDClient) Recv() (*X509SVIDResponse, error) {
	m := new(X509SVIDResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpiffeWorkloadAPIServer is the server API for SpiffeWorkloadAPI service.
type SpiffeWorkloadAPIServer interface {
	// JWT-SVID Profile
	FetchJWTSVID(context.Context, *JWTSVIDRequest) (*JWTSVIDResponse, error)
	FetchJWTBundles(*JWTBundlesRequest, SpiffeWorkloadAPI_FetchJWTBundlesServer) error
	ValidateJWTSVID(context.Context, *ValidateJWTSVIDRequest) (*ValidateJWTSVIDResponse, error)
	// X.509-SVID Profile
	// Fetch all SPIFFE identities the workload is entitled to, as
	// well as related information like trust bundles and CRLs. As
	// this information changes, subsequent messages will be sent.
	FetchX509SVID(*X509SVIDRequest, SpiffeWorkloadAPI_FetchX509SVIDServer) error
}

func RegisterSpiffeWorkloadAPIServer(s *grpc.Server, srv SpiffeWorkloadAPIServer) {
	s.RegisterService(&_SpiffeWorkloadAPI_serviceDesc, srv)
}

func _SpiffeWorkloadAPI_FetchJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiffeWorkloadAPIServer).FetchJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SpiffeWorkloadAPI/FetchJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiffeWorkloadAPIServer).FetchJWTSVID(ctx, req.(*JWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiffeWorkloadAPI_FetchJWTBundles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JWTBundlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpiffeWorkloadAPIServer).FetchJWTBundles(m, &spiffeWorkloadAPIFetchJWTBundlesServer{stream})
}

type SpiffeWorkloadAPI_FetchJWTBundlesServer interface {
	Send(*JWTBundlesResponse) error
	grpc.ServerStream
}

type spiffeWorkloadAPIFetchJWTBundlesServer struct {
	grpc.ServerStream
}

func (x *spiffeWorkloadAPIFetchJWTBundlesServer) Send(m *JWTBundlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SpiffeWorkloadAPI_ValidateJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiffeWorkloadAPIServer).ValidateJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SpiffeWorkloadAPI/ValidateJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiffeWorkloadAPIServer).ValidateJWTSVID(ctx, req.(*ValidateJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiffeWorkloadAPI_FetchX509SVID_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(X509SVIDRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpiffeWorkloadAPIServer).FetchX509SVID(m, &spiffeWorkloadAPIFetchX509SVIDServer{stream})
}

type SpiffeWorkloadAPI_FetchX509SVIDServer interface {
	Send(*X509SVIDResponse) error
	grpc.ServerStream
}

type spiffeWorkloadAPIFetchX509SVIDServer struct {
	grpc.ServerStream
}

func (x *spiffeWorkloadAPIFetchX509SVIDServer) Send(m *X509SVIDResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SpiffeWorkloadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SpiffeWorkloadAPI",
	HandlerType: (*SpiffeWorkloadAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchJWTSVID",
			Handler:    _SpiffeWorkloadAPI_FetchJWTSVID_Handler,
		},
		{
			MethodName: "ValidateJWTSVID",
			Handler:    _SpiffeWorkloadAPI_ValidateJWTSVID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchJWTBundles",
			Handler:       _SpiffeWorkloadAPI_FetchJWTBundles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FetchX509SVID",
			Handler:       _SpiffeWorkloadAPI_FetchX509SVID_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "workload.proto",
}