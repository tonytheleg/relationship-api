// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: relationships/v1/relationship_data.proto

package relationships

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// the aggregate status of the cluster
type IsPropagatedToRelationshipData_Status int32

const (
	IsPropagatedToRelationshipData_STATUS_UNSPECIFIED IsPropagatedToRelationshipData_Status = 0
	IsPropagatedToRelationshipData_STATUS_OTHER       IsPropagatedToRelationshipData_Status = 1
	IsPropagatedToRelationshipData_VIOLATIONS         IsPropagatedToRelationshipData_Status = 2
	IsPropagatedToRelationshipData_NO_VIOLATIONS      IsPropagatedToRelationshipData_Status = 3
)

// Enum value maps for IsPropagatedToRelationshipData_Status.
var (
	IsPropagatedToRelationshipData_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_OTHER",
		2: "VIOLATIONS",
		3: "NO_VIOLATIONS",
	}
	IsPropagatedToRelationshipData_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_OTHER":       1,
		"VIOLATIONS":         2,
		"NO_VIOLATIONS":      3,
	}
)

func (x IsPropagatedToRelationshipData_Status) Enum() *IsPropagatedToRelationshipData_Status {
	p := new(IsPropagatedToRelationshipData_Status)
	*p = x
	return p
}

func (x IsPropagatedToRelationshipData_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IsPropagatedToRelationshipData_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_relationships_v1_relationship_data_proto_enumTypes[0].Descriptor()
}

func (IsPropagatedToRelationshipData_Status) Type() protoreflect.EnumType {
	return &file_relationships_v1_relationship_data_proto_enumTypes[0]
}

func (x IsPropagatedToRelationshipData_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IsPropagatedToRelationshipData_Status.Descriptor instead.
func (IsPropagatedToRelationshipData_Status) EnumDescriptor() ([]byte, []int) {
	return file_relationships_v1_relationship_data_proto_rawDescGZIP(), []int{0, 0}
}

type RegisteredToRelationshipData_RegistrationStatus int32

const (
	RegisteredToRelationshipData_REGISTERED   RegisteredToRelationshipData_RegistrationStatus = 0
	RegisteredToRelationshipData_UNREGISTERED RegisteredToRelationshipData_RegistrationStatus = 1
	RegisteredToRelationshipData_UNKNOWN      RegisteredToRelationshipData_RegistrationStatus = 3
)

// Enum value maps for RegisteredToRelationshipData_RegistrationStatus.
var (
	RegisteredToRelationshipData_RegistrationStatus_name = map[int32]string{
		0: "REGISTERED",
		1: "UNREGISTERED",
		3: "UNKNOWN",
	}
	RegisteredToRelationshipData_RegistrationStatus_value = map[string]int32{
		"REGISTERED":   0,
		"UNREGISTERED": 1,
		"UNKNOWN":      3,
	}
)

func (x RegisteredToRelationshipData_RegistrationStatus) Enum() *RegisteredToRelationshipData_RegistrationStatus {
	p := new(RegisteredToRelationshipData_RegistrationStatus)
	*p = x
	return p
}

func (x RegisteredToRelationshipData_RegistrationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisteredToRelationshipData_RegistrationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_relationships_v1_relationship_data_proto_enumTypes[1].Descriptor()
}

func (RegisteredToRelationshipData_RegistrationStatus) Type() protoreflect.EnumType {
	return &file_relationships_v1_relationship_data_proto_enumTypes[1]
}

func (x RegisteredToRelationshipData_RegistrationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisteredToRelationshipData_RegistrationStatus.Descriptor instead.
func (RegisteredToRelationshipData_RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return file_relationships_v1_relationship_data_proto_rawDescGZIP(), []int{1, 0}
}

type IsPropagatedToRelationshipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status IsPropagatedToRelationshipData_Status `protobuf:"varint,355610639,opt,name=status,proto3,enum=relationships.v1.IsPropagatedToRelationshipData_Status" json:"status,omitempty"`
}

func (x *IsPropagatedToRelationshipData) Reset() {
	*x = IsPropagatedToRelationshipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relationships_v1_relationship_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsPropagatedToRelationshipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsPropagatedToRelationshipData) ProtoMessage() {}

func (x *IsPropagatedToRelationshipData) ProtoReflect() protoreflect.Message {
	mi := &file_relationships_v1_relationship_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsPropagatedToRelationshipData.ProtoReflect.Descriptor instead.
func (*IsPropagatedToRelationshipData) Descriptor() ([]byte, []int) {
	return file_relationships_v1_relationship_data_proto_rawDescGZIP(), []int{0}
}

func (x *IsPropagatedToRelationshipData) GetStatus() IsPropagatedToRelationshipData_Status {
	if x != nil {
		return x.Status
	}
	return IsPropagatedToRelationshipData_STATUS_UNSPECIFIED
}

type RegisteredToRelationshipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status RegisteredToRelationshipData_RegistrationStatus `protobuf:"varint,355610639,opt,name=status,proto3,enum=relationships.v1.RegisteredToRelationshipData_RegistrationStatus" json:"status,omitempty"`
}

func (x *RegisteredToRelationshipData) Reset() {
	*x = RegisteredToRelationshipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relationships_v1_relationship_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredToRelationshipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredToRelationshipData) ProtoMessage() {}

func (x *RegisteredToRelationshipData) ProtoReflect() protoreflect.Message {
	mi := &file_relationships_v1_relationship_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredToRelationshipData.ProtoReflect.Descriptor instead.
func (*RegisteredToRelationshipData) Descriptor() ([]byte, []int) {
	return file_relationships_v1_relationship_data_proto_rawDescGZIP(), []int{1}
}

func (x *RegisteredToRelationshipData) GetStatus() RegisteredToRelationshipData_RegistrationStatus {
	if x != nil {
		return x.Status
	}
	return RegisteredToRelationshipData_REGISTERED
}

var File_relationships_v1_relationship_data_proto protoreflect.FileDescriptor

var file_relationships_v1_relationship_data_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x1e, 0x49, 0x73,
	0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x62, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x8f, 0xe0, 0xc8, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x37, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01,
	0x01, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x55, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x03, 0x22, 0xd1, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x6c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x8f, 0xe0, 0xc8, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x42, 0x5d, 0x0a, 0x11, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x6e, 0x79, 0x74, 0x68, 0x65, 0x6c, 0x65, 0x67, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x3b, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_relationships_v1_relationship_data_proto_rawDescOnce sync.Once
	file_relationships_v1_relationship_data_proto_rawDescData = file_relationships_v1_relationship_data_proto_rawDesc
)

func file_relationships_v1_relationship_data_proto_rawDescGZIP() []byte {
	file_relationships_v1_relationship_data_proto_rawDescOnce.Do(func() {
		file_relationships_v1_relationship_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_relationships_v1_relationship_data_proto_rawDescData)
	})
	return file_relationships_v1_relationship_data_proto_rawDescData
}

var file_relationships_v1_relationship_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_relationships_v1_relationship_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relationships_v1_relationship_data_proto_goTypes = []any{
	(IsPropagatedToRelationshipData_Status)(0),           // 0: relationships.v1.IsPropagatedToRelationshipData.Status
	(RegisteredToRelationshipData_RegistrationStatus)(0), // 1: relationships.v1.RegisteredToRelationshipData.RegistrationStatus
	(*IsPropagatedToRelationshipData)(nil),               // 2: relationships.v1.IsPropagatedToRelationshipData
	(*RegisteredToRelationshipData)(nil),                 // 3: relationships.v1.RegisteredToRelationshipData
}
var file_relationships_v1_relationship_data_proto_depIdxs = []int32{
	0, // 0: relationships.v1.IsPropagatedToRelationshipData.status:type_name -> relationships.v1.IsPropagatedToRelationshipData.Status
	1, // 1: relationships.v1.RegisteredToRelationshipData.status:type_name -> relationships.v1.RegisteredToRelationshipData.RegistrationStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relationships_v1_relationship_data_proto_init() }
func file_relationships_v1_relationship_data_proto_init() {
	if File_relationships_v1_relationship_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relationships_v1_relationship_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IsPropagatedToRelationshipData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relationships_v1_relationship_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisteredToRelationshipData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relationships_v1_relationship_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relationships_v1_relationship_data_proto_goTypes,
		DependencyIndexes: file_relationships_v1_relationship_data_proto_depIdxs,
		EnumInfos:         file_relationships_v1_relationship_data_proto_enumTypes,
		MessageInfos:      file_relationships_v1_relationship_data_proto_msgTypes,
	}.Build()
	File_relationships_v1_relationship_data_proto = out.File
	file_relationships_v1_relationship_data_proto_rawDesc = nil
	file_relationships_v1_relationship_data_proto_goTypes = nil
	file_relationships_v1_relationship_data_proto_depIdxs = nil
}
