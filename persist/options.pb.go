// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persist/options.proto

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	persist/options.proto

It has these top-level messages:
	QLImpl
	TypeMapping
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PersistenceOptions int32

const (
	PersistenceOptions_SQL     PersistenceOptions = 0
	PersistenceOptions_SPANNER PersistenceOptions = 1
)

var PersistenceOptions_name = map[int32]string{
	0: "SQL",
	1: "SPANNER",
}
var PersistenceOptions_value = map[string]int32{
	"SQL":     0,
	"SPANNER": 1,
}

func (x PersistenceOptions) Enum() *PersistenceOptions {
	p := new(PersistenceOptions)
	*p = x
	return p
}
func (x PersistenceOptions) String() string {
	return proto.EnumName(PersistenceOptions_name, int32(x))
}
func (x *PersistenceOptions) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PersistenceOptions_value, data, "PersistenceOptions")
	if err != nil {
		return err
	}
	*x = PersistenceOptions(value)
	return nil
}
func (PersistenceOptions) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type QLImpl struct {
	Query            []string `protobuf:"bytes,1,rep,name=query" json:"query,omitempty"`
	Arguments        []string `protobuf:"bytes,2,rep,name=arguments" json:"arguments,omitempty"`
	Before           *bool    `protobuf:"varint,10,opt,name=before" json:"before,omitempty"`
	After            *bool    `protobuf:"varint,11,opt,name=after" json:"after,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *QLImpl) Reset()                    { *m = QLImpl{} }
func (m *QLImpl) String() string            { return proto.CompactTextString(m) }
func (*QLImpl) ProtoMessage()               {}
func (*QLImpl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QLImpl) GetQuery() []string {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *QLImpl) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *QLImpl) GetBefore() bool {
	if m != nil && m.Before != nil {
		return *m.Before
	}
	return false
}

func (m *QLImpl) GetAfter() bool {
	if m != nil && m.After != nil {
		return *m.After
	}
	return false
}

type TypeMapping struct {
	Types            []*TypeMapping_TypeDescriptor `protobuf:"bytes,1,rep,name=types" json:"types,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *TypeMapping) Reset()                    { *m = TypeMapping{} }
func (m *TypeMapping) String() string            { return proto.CompactTextString(m) }
func (*TypeMapping) ProtoMessage()               {}
func (*TypeMapping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TypeMapping) GetTypes() []*TypeMapping_TypeDescriptor {
	if m != nil {
		return m.Types
	}
	return nil
}

type TypeMapping_TypeDescriptor struct {
	ProtoTypeName    *string                                     `protobuf:"bytes,1,opt,name=proto_type_name,json=protoTypeName" json:"proto_type_name,omitempty"`
	ProtoType        *google_protobuf.FieldDescriptorProto_Type  `protobuf:"varint,2,opt,name=proto_type,json=protoType,enum=google.protobuf.FieldDescriptorProto_Type" json:"proto_type,omitempty"`
	ProtoLabel       *google_protobuf.FieldDescriptorProto_Label `protobuf:"varint,3,opt,name=proto_label,json=protoLabel,enum=google.protobuf.FieldDescriptorProto_Label" json:"proto_label,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *TypeMapping_TypeDescriptor) Reset()                    { *m = TypeMapping_TypeDescriptor{} }
func (m *TypeMapping_TypeDescriptor) String() string            { return proto.CompactTextString(m) }
func (*TypeMapping_TypeDescriptor) ProtoMessage()               {}
func (*TypeMapping_TypeDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *TypeMapping_TypeDescriptor) GetProtoTypeName() string {
	if m != nil && m.ProtoTypeName != nil {
		return *m.ProtoTypeName
	}
	return ""
}

func (m *TypeMapping_TypeDescriptor) GetProtoType() google_protobuf.FieldDescriptorProto_Type {
	if m != nil && m.ProtoType != nil {
		return *m.ProtoType
	}
	return google_protobuf.FieldDescriptorProto_TYPE_DOUBLE
}

func (m *TypeMapping_TypeDescriptor) GetProtoLabel() google_protobuf.FieldDescriptorProto_Label {
	if m != nil && m.ProtoLabel != nil {
		return *m.ProtoLabel
	}
	return google_protobuf.FieldDescriptorProto_LABEL_OPTIONAL
}

var E_Pkg = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         560003,
	Name:          "persist.pkg",
	Tag:           "bytes,560003,opt,name=pkg",
	Filename:      "persist/options.proto",
}

var E_Ql = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*QLImpl)(nil),
	Field:         560000,
	Name:          "persist.ql",
	Tag:           "bytes,560000,opt,name=ql",
	Filename:      "persist/options.proto",
}

var E_Mapping = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*TypeMapping)(nil),
	Field:         560001,
	Name:          "persist.mapping",
	Tag:           "bytes,560001,opt,name=mapping",
	Filename:      "persist/options.proto",
}

var E_ServiceType = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*PersistenceOptions)(nil),
	Field:         560002,
	Name:          "persist.service_type",
	Tag:           "varint,560002,opt,name=service_type,json=serviceType,enum=persist.PersistenceOptions",
	Filename:      "persist/options.proto",
}

func init() {
	proto.RegisterType((*QLImpl)(nil), "persist.QLImpl")
	proto.RegisterType((*TypeMapping)(nil), "persist.TypeMapping")
	proto.RegisterType((*TypeMapping_TypeDescriptor)(nil), "persist.TypeMapping.TypeDescriptor")
	proto.RegisterEnum("persist.PersistenceOptions", PersistenceOptions_name, PersistenceOptions_value)
	proto.RegisterExtension(E_Pkg)
	proto.RegisterExtension(E_Ql)
	proto.RegisterExtension(E_Mapping)
	proto.RegisterExtension(E_ServiceType)
}

func init() { proto.RegisterFile("persist/options.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xab, 0xad, 0xf4, 0x04, 0xba, 0xc9, 0x1a, 0xc8, 0x1a, 0x13, 0x44, 0x45, 0x42,
	0x55, 0xd1, 0x52, 0xd4, 0x0b, 0x04, 0xe5, 0x6a, 0x08, 0x90, 0x26, 0x75, 0xa5, 0x75, 0xb9, 0xe2,
	0x66, 0x4a, 0xd3, 0xd3, 0x2c, 0xe0, 0xc4, 0xae, 0xe3, 0x20, 0xf5, 0x8e, 0x3f, 0x0f, 0xc1, 0x0d,
	0x2f, 0xc5, 0x1b, 0xa1, 0xd8, 0x49, 0x0a, 0xea, 0x24, 0x76, 0x65, 0xfb, 0xf8, 0xfb, 0x7e, 0xe7,
	0xe4, 0x9c, 0x18, 0xee, 0x49, 0x54, 0x59, 0x9c, 0xe9, 0x81, 0x90, 0x3a, 0x16, 0x69, 0xe6, 0x4b,
	0x25, 0xb4, 0x20, 0xad, 0x32, 0x7c, 0xec, 0x45, 0x42, 0x44, 0x1c, 0x07, 0x26, 0xbc, 0xc8, 0x57,
	0x83, 0x25, 0x66, 0xa1, 0x8a, 0xa5, 0x16, 0xca, 0x4a, 0xbb, 0x9f, 0x60, 0x7f, 0x36, 0x3e, 0x4f,
	0x24, 0x27, 0x47, 0xb0, 0xb7, 0xce, 0x51, 0x6d, 0xa8, 0xe3, 0x35, 0x7b, 0x6d, 0x66, 0x0f, 0xe4,
	0x04, 0xda, 0x81, 0x8a, 0xf2, 0x04, 0x53, 0x9d, 0xd1, 0x86, 0xb9, 0xd9, 0x06, 0xc8, 0x7d, 0xd8,
	0x5f, 0xe0, 0x4a, 0x28, 0xa4, 0xe0, 0x39, 0xbd, 0xdb, 0xac, 0x3c, 0x15, 0xac, 0x60, 0xa5, 0x51,
	0x51, 0xd7, 0x84, 0xed, 0xa1, 0xfb, 0xab, 0x01, 0xee, 0x87, 0x8d, 0xc4, 0x8b, 0x40, 0xca, 0x38,
	0x8d, 0xc8, 0x4b, 0xd8, 0xd3, 0x1b, 0x89, 0x99, 0xc9, 0xe8, 0x0e, 0x1f, 0xfb, 0x65, 0xd9, 0xfe,
	0x5f, 0x22, 0xb3, 0x7f, 0x53, 0x57, 0xcd, 0xac, 0xe3, 0xf8, 0xb7, 0x03, 0x9d, 0x7f, 0x6f, 0xc8,
	0x13, 0x38, 0x30, 0x9f, 0x74, 0x59, 0x28, 0x2e, 0xd3, 0x20, 0x41, 0xea, 0x78, 0x4e, 0xaf, 0xcd,
	0xee, 0x9a, 0x70, 0xa1, 0x9e, 0x04, 0x09, 0x92, 0x73, 0x80, 0xad, 0x8e, 0x36, 0x3c, 0xa7, 0xd7,
	0x19, 0xf6, 0x7d, 0xdb, 0x28, 0xbf, 0x6a, 0x94, 0xff, 0x2e, 0x46, 0xbe, 0xdc, 0xd2, 0xa7, 0x45,
	0xdc, 0xd4, 0xc2, 0xda, 0x35, 0x8e, 0x8c, 0xc1, 0xb5, 0x28, 0x1e, 0x2c, 0x90, 0xd3, 0xa6, 0x61,
	0x3d, 0xbd, 0x19, 0x6b, 0x5c, 0x58, 0x98, 0x2d, 0xc5, 0xec, 0xfb, 0x7d, 0x20, 0x53, 0xdb, 0x00,
	0x4c, 0x43, 0x7c, 0x6f, 0x27, 0x4a, 0x5a, 0xd0, 0x9c, 0xcf, 0xc6, 0x87, 0xb7, 0x88, 0x0b, 0xad,
	0xf9, 0xf4, 0x6c, 0x32, 0x79, 0xcb, 0x0e, 0x9d, 0xd1, 0x33, 0x68, 0xca, 0xcf, 0x11, 0x39, 0xb9,
	0x26, 0x17, 0xaf, 0xac, 0xf4, 0xc7, 0xcf, 0xae, 0x69, 0x40, 0x21, 0x1d, 0x9d, 0x41, 0x63, 0xcd,
	0xc9, 0xc3, 0x1d, 0xc3, 0x05, 0xea, 0x2b, 0xb1, 0xac, 0x2c, 0x5f, 0x8d, 0xc5, 0x1d, 0x1e, 0xd4,
	0xb3, 0xb0, 0x7f, 0x07, 0x6b, 0xac, 0xf9, 0x68, 0x06, 0xad, 0xa4, 0x1c, 0xdd, 0xa3, 0x1d, 0xce,
	0x1c, 0xd5, 0x97, 0xb8, 0x2e, 0x9b, 0x7e, 0x2b, 0x41, 0x47, 0xd7, 0x0d, 0x95, 0x55, 0x9c, 0x51,
	0x00, 0x77, 0x32, 0x6b, 0x34, 0xe3, 0xf8, 0x3f, 0xf7, 0xbb, 0xe1, 0x76, 0x86, 0x0f, 0x6a, 0xee,
	0x6e, 0xcf, 0x98, 0x5b, 0x32, 0x8b, 0x94, 0xaf, 0x5f, 0x7c, 0x7c, 0x1e, 0xc5, 0xfa, 0x2a, 0x5f,
	0xf8, 0xa1, 0x48, 0x06, 0x3a, 0x4c, 0x43, 0x2e, 0xf2, 0xa5, 0x7d, 0x12, 0xe1, 0x69, 0x84, 0xe9,
	0x69, 0xf5, 0x88, 0xca, 0xf5, 0x55, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x46, 0xfd, 0x6a,
	0x9a, 0x5e, 0x03, 0x00, 0x00,
}
