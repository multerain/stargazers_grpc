// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: stargazers/stargazers.proto

package stargazerspb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the project and repository to query
type StargazerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project    string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *StargazerRequest) Reset() {
	*x = StargazerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stargazers_stargazers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StargazerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StargazerRequest) ProtoMessage() {}

func (x *StargazerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stargazers_stargazers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StargazerRequest.ProtoReflect.Descriptor instead.
func (*StargazerRequest) Descriptor() ([]byte, []int) {
	return file_stargazers_stargazers_proto_rawDescGZIP(), []int{0}
}

func (x *StargazerRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *StargazerRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

// The response is the number of "stargazers" on a given GitHub repository
type StargazersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StargazersResponse) Reset() {
	*x = StargazersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stargazers_stargazers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StargazersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StargazersResponse) ProtoMessage() {}

func (x *StargazersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stargazers_stargazers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StargazersResponse.ProtoReflect.Descriptor instead.
func (*StargazersResponse) Descriptor() ([]byte, []int) {
	return file_stargazers_stargazers_proto_rawDescGZIP(), []int{1}
}

func (x *StargazersResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_stargazers_stargazers_proto protoreflect.FileDescriptor

var file_stargazers_stargazers_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x67,
	0x61, 0x7a, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0x5d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72,
	0x73, 0x12, 0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65,
	0x72, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73,
	0x3b, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stargazers_stargazers_proto_rawDescOnce sync.Once
	file_stargazers_stargazers_proto_rawDescData = file_stargazers_stargazers_proto_rawDesc
)

func file_stargazers_stargazers_proto_rawDescGZIP() []byte {
	file_stargazers_stargazers_proto_rawDescOnce.Do(func() {
		file_stargazers_stargazers_proto_rawDescData = protoimpl.X.CompressGZIP(file_stargazers_stargazers_proto_rawDescData)
	})
	return file_stargazers_stargazers_proto_rawDescData
}

var file_stargazers_stargazers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stargazers_stargazers_proto_goTypes = []interface{}{
	(*StargazerRequest)(nil),   // 0: stargazers.StargazerRequest
	(*StargazersResponse)(nil), // 1: stargazers.StargazersResponse
}
var file_stargazers_stargazers_proto_depIdxs = []int32{
	0, // 0: stargazers.Stargazers.GetStargazers:input_type -> stargazers.StargazerRequest
	1, // 1: stargazers.Stargazers.GetStargazers:output_type -> stargazers.StargazersResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stargazers_stargazers_proto_init() }
func file_stargazers_stargazers_proto_init() {
	if File_stargazers_stargazers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stargazers_stargazers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StargazerRequest); i {
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
		file_stargazers_stargazers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StargazersResponse); i {
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
			RawDescriptor: file_stargazers_stargazers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stargazers_stargazers_proto_goTypes,
		DependencyIndexes: file_stargazers_stargazers_proto_depIdxs,
		MessageInfos:      file_stargazers_stargazers_proto_msgTypes,
	}.Build()
	File_stargazers_stargazers_proto = out.File
	file_stargazers_stargazers_proto_rawDesc = nil
	file_stargazers_stargazers_proto_goTypes = nil
	file_stargazers_stargazers_proto_depIdxs = nil
}
