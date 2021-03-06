// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/campaign_shared_set_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum_CampaignSharedSetError int32

const (
	// Enum unspecified.
	CampaignSharedSetErrorEnum_UNSPECIFIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 0
	// The received error code is not known in this version.
	CampaignSharedSetErrorEnum_UNKNOWN CampaignSharedSetErrorEnum_CampaignSharedSetError = 1
	// The shared set belongs to another customer and permission isn't granted.
	CampaignSharedSetErrorEnum_SHARED_SET_ACCESS_DENIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 2
)

var CampaignSharedSetErrorEnum_CampaignSharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SHARED_SET_ACCESS_DENIED",
}
var CampaignSharedSetErrorEnum_CampaignSharedSetError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"SHARED_SET_ACCESS_DENIED": 2,
}

func (x CampaignSharedSetErrorEnum_CampaignSharedSetError) String() string {
	return proto.EnumName(CampaignSharedSetErrorEnum_CampaignSharedSetError_name, int32(x))
}
func (CampaignSharedSetErrorEnum_CampaignSharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_campaign_shared_set_error_6fb036bae1aa3bd9, []int{0, 0}
}

// Container for enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetErrorEnum) Reset()         { *m = CampaignSharedSetErrorEnum{} }
func (m *CampaignSharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetErrorEnum) ProtoMessage()    {}
func (*CampaignSharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_shared_set_error_6fb036bae1aa3bd9, []int{0}
}
func (m *CampaignSharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CampaignSharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetErrorEnum.Merge(dst, src)
}
func (m *CampaignSharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Size(m)
}
func (m *CampaignSharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignSharedSetErrorEnum)(nil), "google.ads.googleads.v0.errors.CampaignSharedSetErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CampaignSharedSetErrorEnum_CampaignSharedSetError", CampaignSharedSetErrorEnum_CampaignSharedSetError_name, CampaignSharedSetErrorEnum_CampaignSharedSetError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/campaign_shared_set_error.proto", fileDescriptor_campaign_shared_set_error_6fb036bae1aa3bd9)
}

var fileDescriptor_campaign_shared_set_error_6fb036bae1aa3bd9 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x15, 0x14, 0xb2, 0x83, 0xa3, 0x07, 0x11, 0x95, 0x1d, 0xfa, 0x00, 0x69, 0xc1, 0xa3,
	0x20, 0x64, 0x6d, 0x9c, 0x43, 0x88, 0xc5, 0x6c, 0x13, 0xa4, 0x10, 0xe2, 0x12, 0xe2, 0x60, 0x6d,
	0x4a, 0x52, 0xf7, 0x40, 0x1e, 0x7d, 0x14, 0x6f, 0xbe, 0x91, 0x24, 0xd9, 0x76, 0x9a, 0x9e, 0xf2,
	0xc1, 0xf7, 0x2f, 0xbf, 0x0f, 0xdc, 0x29, 0xad, 0xd5, 0x5a, 0x66, 0x5c, 0xd8, 0x2c, 0x40, 0x87,
	0x36, 0x79, 0x26, 0x8d, 0xd1, 0xc6, 0x66, 0x4b, 0xde, 0x74, 0x7c, 0xa5, 0x5a, 0x66, 0xdf, 0xb9,
	0x91, 0x82, 0x59, 0xd9, 0x33, 0x4f, 0xc1, 0xce, 0xe8, 0x5e, 0x27, 0xa3, 0x60, 0x82, 0x5c, 0x58,
	0xb8, 0xf7, 0xc3, 0x4d, 0x0e, 0x83, 0x3f, 0x35, 0xe0, 0xb2, 0xd8, 0x46, 0x50, 0x9f, 0x40, 0x65,
	0x8f, 0x1d, 0x85, 0xdb, 0x8f, 0x26, 0x9d, 0x81, 0xf3, 0xc3, 0x6c, 0x72, 0x06, 0x06, 0x73, 0x42,
	0x2b, 0x5c, 0x4c, 0xef, 0xa7, 0xb8, 0x1c, 0x1e, 0x25, 0x03, 0x70, 0x3a, 0x27, 0x8f, 0xe4, 0xe9,
	0x85, 0x0c, 0xa3, 0xe4, 0x1a, 0x5c, 0xd0, 0x07, 0xf4, 0x8c, 0x4b, 0x46, 0xf1, 0x8c, 0xa1, 0xa2,
	0xc0, 0x94, 0xb2, 0x12, 0x13, 0x27, 0x8d, 0xc7, 0x3f, 0x11, 0x48, 0x97, 0xba, 0x81, 0xff, 0x7f,
	0x6d, 0x7c, 0x75, 0xb8, 0xba, 0x72, 0x77, 0x55, 0xd1, 0x6b, 0xb9, 0xb5, 0x2b, 0xbd, 0xe6, 0xad,
	0x82, 0xda, 0xa8, 0x4c, 0xc9, 0xd6, 0x5f, 0xbd, 0x5b, 0xaa, 0x5b, 0xd9, 0xbf, 0x86, 0xbb, 0x0d,
	0xcf, 0x67, 0x7c, 0x3c, 0x41, 0xe8, 0x2b, 0x1e, 0x4d, 0x42, 0x18, 0x12, 0x16, 0x06, 0xe8, 0xd0,
	0x22, 0x87, 0xbe, 0xd2, 0x7e, 0xef, 0x04, 0x35, 0x12, 0xb6, 0xde, 0x0b, 0xea, 0x45, 0x5e, 0x07,
	0xc1, 0xdb, 0x89, 0x2f, 0xbe, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x05, 0x13, 0x71, 0xbb, 0xb0,
	0x01, 0x00, 0x00,
}
