// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: types/operation.proto

package types

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The server-assigned name, which is only unique within the same service that
	// originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *anypb.Any `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//
	//	*Operation_Error
	//	*Operation_Response
	Result        isOperation_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Operation) Reset() {
	*x = Operation{}
	mi := &file_types_operation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_types_operation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_types_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Operation) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Operation) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Operation) GetResult() isOperation_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Operation) GetError() *status.Status {
	if x != nil {
		if x, ok := x.Result.(*Operation_Error); ok {
			return x.Error
		}
	}
	return nil
}

func (x *Operation) GetResponse() *anypb.Any {
	if x != nil {
		if x, ok := x.Result.(*Operation_Response); ok {
			return x.Response
		}
	}
	return nil
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	// The error result of the operation in case of failure or cancellation.
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	// The normal response of the operation in case of success.  If the original
	// method returns no data on success, such as `Delete`, the response is
	// `google.protobuf.Empty`.  If the original method is standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For other
	// methods, the response should have the type `XxxResponse`, where `Xxx`
	// is the original method name.  For example, if the original method name
	// is `TakeSnapshot()`, the inferred response type is
	// `TakeSnapshotResponse`.
	Response *anypb.Any `protobuf:"bytes,5,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

var File_types_operation_proto protoreflect.FileDescriptor

const file_types_operation_proto_rawDesc = "" +
	"\n" +
	"\x15types/operation.proto\x12\x0ev10proto.types\x1a\x19google/protobuf/any.proto\x1a\x17google/rpc/status.proto\"\xcf\x01\n" +
	"\tOperation\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x120\n" +
	"\bmetadata\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\bmetadata\x12\x12\n" +
	"\x04done\x18\x03 \x01(\bR\x04done\x12*\n" +
	"\x05error\x18\x04 \x01(\v2\x12.google.rpc.StatusH\x00R\x05error\x122\n" +
	"\bresponse\x18\x05 \x01(\v2\x14.google.protobuf.AnyH\x00R\bresponseB\b\n" +
	"\x06resultB#Z!github.com/drival-ai/v10-go/typesb\x06proto3"

var (
	file_types_operation_proto_rawDescOnce sync.Once
	file_types_operation_proto_rawDescData []byte
)

func file_types_operation_proto_rawDescGZIP() []byte {
	file_types_operation_proto_rawDescOnce.Do(func() {
		file_types_operation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_operation_proto_rawDesc), len(file_types_operation_proto_rawDesc)))
	})
	return file_types_operation_proto_rawDescData
}

var file_types_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_operation_proto_goTypes = []any{
	(*Operation)(nil),     // 0: v10proto.types.Operation
	(*anypb.Any)(nil),     // 1: google.protobuf.Any
	(*status.Status)(nil), // 2: google.rpc.Status
}
var file_types_operation_proto_depIdxs = []int32{
	1, // 0: v10proto.types.Operation.metadata:type_name -> google.protobuf.Any
	2, // 1: v10proto.types.Operation.error:type_name -> google.rpc.Status
	1, // 2: v10proto.types.Operation.response:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_types_operation_proto_init() }
func file_types_operation_proto_init() {
	if File_types_operation_proto != nil {
		return
	}
	file_types_operation_proto_msgTypes[0].OneofWrappers = []any{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_operation_proto_rawDesc), len(file_types_operation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_operation_proto_goTypes,
		DependencyIndexes: file_types_operation_proto_depIdxs,
		MessageInfos:      file_types_operation_proto_msgTypes,
	}.Build()
	File_types_operation_proto = out.File
	file_types_operation_proto_goTypes = nil
	file_types_operation_proto_depIdxs = nil
}
