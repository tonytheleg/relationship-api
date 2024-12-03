// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.3
// source: api/resources/v1/resource_types.proto

package resources

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RhelHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata     *Metadata             `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ResourceData *RhelHostResourceData `protobuf:"bytes,2,opt,name=resource_data,json=resourceData,proto3" json:"resource_data,omitempty"`
}

func (x *RhelHost) Reset() {
	*x = RhelHost{}
	mi := &file_api_resources_v1_resource_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RhelHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RhelHost) ProtoMessage() {}

func (x *RhelHost) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resource_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RhelHost.ProtoReflect.Descriptor instead.
func (*RhelHost) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resource_types_proto_rawDescGZIP(), []int{0}
}

func (x *RhelHost) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RhelHost) GetResourceData() *RhelHostResourceData {
	if x != nil {
		return x.ResourceData
	}
	return nil
}

type K8SCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata     *Metadata               `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ResourceData *K8SClusterResourceData `protobuf:"bytes,2,opt,name=resource_data,json=resourceData,proto3" json:"resource_data,omitempty"`
}

func (x *K8SCluster) Reset() {
	*x = K8SCluster{}
	mi := &file_api_resources_v1_resource_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *K8SCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SCluster) ProtoMessage() {}

func (x *K8SCluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resource_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SCluster.ProtoReflect.Descriptor instead.
func (*K8SCluster) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resource_types_proto_rawDescGZIP(), []int{1}
}

func (x *K8SCluster) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *K8SCluster) GetResourceData() *K8SClusterResourceData {
	if x != nil {
		return x.ResourceData
	}
	return nil
}

var File_api_resources_v1_resource_types_proto protoreflect.FileDescriptor

var file_api_resources_v1_resource_types_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x08,
	0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x0a, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x37, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_resources_v1_resource_types_proto_rawDescOnce sync.Once
	file_api_resources_v1_resource_types_proto_rawDescData = file_api_resources_v1_resource_types_proto_rawDesc
)

func file_api_resources_v1_resource_types_proto_rawDescGZIP() []byte {
	file_api_resources_v1_resource_types_proto_rawDescOnce.Do(func() {
		file_api_resources_v1_resource_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_resources_v1_resource_types_proto_rawDescData)
	})
	return file_api_resources_v1_resource_types_proto_rawDescData
}

var file_api_resources_v1_resource_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_resources_v1_resource_types_proto_goTypes = []any{
	(*RhelHost)(nil),               // 0: resources.v1.RhelHost
	(*K8SCluster)(nil),             // 1: resources.v1.K8sCluster
	(*Metadata)(nil),               // 2: resources.v1.Metadata
	(*RhelHostResourceData)(nil),   // 3: resources.v1.RhelHostResourceData
	(*K8SClusterResourceData)(nil), // 4: resources.v1.K8sClusterResourceData
}
var file_api_resources_v1_resource_types_proto_depIdxs = []int32{
	2, // 0: resources.v1.RhelHost.metadata:type_name -> resources.v1.Metadata
	3, // 1: resources.v1.RhelHost.resource_data:type_name -> resources.v1.RhelHostResourceData
	2, // 2: resources.v1.K8sCluster.metadata:type_name -> resources.v1.Metadata
	4, // 3: resources.v1.K8sCluster.resource_data:type_name -> resources.v1.K8sClusterResourceData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_resources_v1_resource_types_proto_init() }
func file_api_resources_v1_resource_types_proto_init() {
	if File_api_resources_v1_resource_types_proto != nil {
		return
	}
	file_api_resources_v1_common_proto_init()
	file_api_resources_v1_resource_data_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_resources_v1_resource_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_resources_v1_resource_types_proto_goTypes,
		DependencyIndexes: file_api_resources_v1_resource_types_proto_depIdxs,
		MessageInfos:      file_api_resources_v1_resource_types_proto_msgTypes,
	}.Build()
	File_api_resources_v1_resource_types_proto = out.File
	file_api_resources_v1_resource_types_proto_rawDesc = nil
	file_api_resources_v1_resource_types_proto_goTypes = nil
	file_api_resources_v1_resource_types_proto_depIdxs = nil
}
