// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/FeatureRow.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type FeatureRow struct {
	EntityKey            string               `protobuf:"bytes,1,opt,name=entityKey,proto3" json:"entityKey,omitempty"`
	Features             []*Feature           `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	EventTimestamp       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	EntityName           string               `protobuf:"bytes,4,opt,name=entityName,proto3" json:"entityName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureRow) Reset()         { *m = FeatureRow{} }
func (m *FeatureRow) String() string { return proto.CompactTextString(m) }
func (*FeatureRow) ProtoMessage()    {}
func (*FeatureRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbbea9c89787d1c7, []int{0}
}

func (m *FeatureRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRow.Unmarshal(m, b)
}
func (m *FeatureRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRow.Marshal(b, m, deterministic)
}
func (m *FeatureRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRow.Merge(m, src)
}
func (m *FeatureRow) XXX_Size() int {
	return xxx_messageInfo_FeatureRow.Size(m)
}
func (m *FeatureRow) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRow.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRow proto.InternalMessageInfo

func (m *FeatureRow) GetEntityKey() string {
	if m != nil {
		return m.EntityKey
	}
	return ""
}

func (m *FeatureRow) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureRow) GetEventTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.EventTimestamp
	}
	return nil
}

func (m *FeatureRow) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func init() {
	proto.RegisterType((*FeatureRow)(nil), "feast.types.FeatureRow")
}

func init() { proto.RegisterFile("feast/types/FeatureRow.proto", fileDescriptor_fbbea9c89787d1c7) }

var fileDescriptor_fbbea9c89787d1c7 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x85, 0x40,
	0x14, 0x86, 0xb1, 0x1b, 0xd1, 0x3d, 0x42, 0xc1, 0xd0, 0xc2, 0xe4, 0x52, 0xd2, 0xca, 0xd5, 0x9c,
	0xb8, 0x41, 0x0f, 0xe0, 0xa2, 0x4d, 0x10, 0x21, 0xd1, 0xa2, 0xdd, 0x58, 0xc7, 0xc9, 0x4a, 0x47,
	0x9c, 0x63, 0xe1, 0xdb, 0xf5, 0x68, 0xd1, 0x0c, 0x5e, 0x87, 0x68, 0x7b, 0xfe, 0xcf, 0xdf, 0x6f,
	0x7e, 0xd8, 0xd4, 0xa4, 0x2c, 0x23, 0x4f, 0x3d, 0x59, 0xbc, 0x21, 0xc5, 0xe3, 0x40, 0xa5, 0xf9,
	0x92, 0xfd, 0x60, 0xd8, 0x88, 0xd8, 0xa5, 0xd2, 0xa5, 0xe9, 0xb9, 0x36, 0x46, 0x7f, 0x10, 0xba,
	0xa8, 0x1a, 0x6b, 0xe4, 0xa6, 0x25, 0xcb, 0xaa, 0xed, 0x3d, 0x9d, 0x9e, 0xfe, 0xd3, 0xe5, 0xa3,
	0x8b, 0xef, 0x08, 0x60, 0x69, 0x17, 0x1b, 0x58, 0x53, 0xc7, 0x0d, 0x4f, 0xb7, 0x34, 0x25, 0x51,
	0x16, 0xe5, 0xeb, 0x72, 0x39, 0x88, 0x4b, 0x38, 0xac, 0x3d, 0x6b, 0x93, 0xbd, 0x6c, 0x95, 0xc7,
	0xdb, 0x13, 0x19, 0x88, 0xc8, 0xb9, 0x68, 0x47, 0x89, 0x02, 0x8e, 0xe8, 0x93, 0x3a, 0x7e, 0x98,
	0x8d, 0x92, 0x55, 0x16, 0xe5, 0xf1, 0x36, 0x95, 0xde, 0x59, 0xce, 0xce, 0x72, 0x47, 0x94, 0x7f,
	0xbe, 0x10, 0x67, 0x00, 0x5e, 0xe1, 0x4e, 0xb5, 0x94, 0xec, 0x3b, 0xa9, 0xe0, 0x52, 0x3c, 0x42,
	0xb8, 0x46, 0x71, 0xbc, 0x3c, 0xe7, 0xfe, 0xb7, 0xfc, 0xe9, 0x5a, 0x37, 0xfc, 0x3a, 0x56, 0xf2,
	0xd9, 0xb4, 0xa8, 0xcd, 0x1b, 0xbd, 0xa3, 0x9f, 0xc3, 0xfd, 0xda, 0xa2, 0xa6, 0x8e, 0x06, 0xc5,
	0xf4, 0x82, 0xda, 0x60, 0x30, 0x54, 0x75, 0xe0, 0x80, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0xcd, 0xff, 0x6c, 0x8a, 0x01, 0x00, 0x00,
}
