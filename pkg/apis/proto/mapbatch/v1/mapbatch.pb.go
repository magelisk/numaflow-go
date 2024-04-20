// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: pkg/apis/proto/mapbatch/v1/mapbatch.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/numaproj/numaflow-go/pkg/apis/proto/map/v1"
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

// *
// MapRequest represents a request element.
type MapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string             `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *timestamp.Timestamp `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
	Headers   map[string]string    `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapRequest) Reset() {
	*x = MapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequest) ProtoMessage() {}

func (x *MapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequest.ProtoReflect.Descriptor instead.
func (*MapRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescGZIP(), []int{0}
}

func (x *MapRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *MapRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MapRequest) GetEventTime() *timestamp.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *MapRequest) GetWatermark() *timestamp.Timestamp {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *MapRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type MapBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MapRequest `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MapBatchRequest) Reset() {
	*x = MapBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapBatchRequest) ProtoMessage() {}

func (x *MapBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapBatchRequest.ProtoReflect.Descriptor instead.
func (*MapBatchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescGZIP(), []int{1}
}

func (x *MapBatchRequest) GetMessages() []*MapRequest {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_pkg_apis_proto_mapbatch_v1_mapbatch_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x61, 0x70, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x70,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65,
	0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x70, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x46, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x70, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x87, 0x01, 0x0a, 0x08, 0x4d, 0x61,
	0x70, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x41, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x46, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x70, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x07, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x6d,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x66,
	0x6c, 0x6f, 0x77, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescData = file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDesc
)

func file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescData)
	})
	return file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDescData
}

var file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_goTypes = []interface{}{
	(*MapRequest)(nil),          // 0: mapbatch.v1.MapRequest
	(*MapBatchRequest)(nil),     // 1: mapbatch.v1.MapBatchRequest
	nil,                         // 2: mapbatch.v1.MapRequest.HeadersEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
	(*v1.MapResponse)(nil),      // 5: map.v1.MapResponse
	(*v1.ReadyResponse)(nil),    // 6: map.v1.ReadyResponse
}
var file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_depIdxs = []int32{
	3, // 0: mapbatch.v1.MapRequest.event_time:type_name -> google.protobuf.Timestamp
	3, // 1: mapbatch.v1.MapRequest.watermark:type_name -> google.protobuf.Timestamp
	2, // 2: mapbatch.v1.MapRequest.headers:type_name -> mapbatch.v1.MapRequest.HeadersEntry
	0, // 3: mapbatch.v1.MapBatchRequest.messages:type_name -> mapbatch.v1.MapRequest
	1, // 4: mapbatch.v1.MapBatch.MapBatchFn:input_type -> mapbatch.v1.MapBatchRequest
	4, // 5: mapbatch.v1.MapBatch.IsReady:input_type -> google.protobuf.Empty
	5, // 6: mapbatch.v1.MapBatch.MapBatchFn:output_type -> map.v1.MapResponse
	6, // 7: mapbatch.v1.MapBatch.IsReady:output_type -> map.v1.ReadyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_init() }
func file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_init() {
	if File_pkg_apis_proto_mapbatch_v1_mapbatch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRequest); i {
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
		file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapBatchRequest); i {
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
			RawDescriptor: file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_depIdxs,
		MessageInfos:      file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_mapbatch_v1_mapbatch_proto = out.File
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_rawDesc = nil
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_goTypes = nil
	file_pkg_apis_proto_mapbatch_v1_mapbatch_proto_depIdxs = nil
}
