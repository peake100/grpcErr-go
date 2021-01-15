// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: grpeakec_proto/pkerr/error.proto

package protogen

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	cerealMessages "github.com/illuscio-dev/protoCereal-go/cerealMessages"
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

// Trace info holds information about where an error occurred.
type TraceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app name is the name of the service or process that generated the error.
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	// Identifier for the host the app is running on.
	AppHost string `protobuf:"bytes,2,opt,name=app_host,json=appHost,proto3" json:"app_host,omitempty"`
	// stack_trace of the error (can be ""). No enforced format. This is meant for human
	// reference).
	StackTrace string `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	// Some languages let an error continue to gather context, with wrapping errors in go.
	// This field can be used to store additional context added to the error after it
	// was created.
	AdditionalContext string `protobuf:"bytes,4,opt,name=additional_context,json=additionalContext,proto3" json:"additional_context,omitempty"`
}

func (x *TraceInfo) Reset() {
	*x = TraceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpeakec_proto_pkerr_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceInfo) ProtoMessage() {}

func (x *TraceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpeakec_proto_pkerr_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceInfo.ProtoReflect.Descriptor instead.
func (*TraceInfo) Descriptor() ([]byte, []int) {
	return file_grpeakec_proto_pkerr_error_proto_rawDescGZIP(), []int{0}
}

func (x *TraceInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *TraceInfo) GetAppHost() string {
	if x != nil {
		return x.AppHost
	}
	return ""
}

func (x *TraceInfo) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *TraceInfo) GetAdditionalContext() string {
	if x != nil {
		return x.AdditionalContext
	}
	return ""
}

// Error holds information about an error.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a uuid that uniquely identifies this error.
	Id *cerealMessages.UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// issuer is the issuer of a code. If multiple services use this library, they
	// can differentiate their error codes by having unique issuers. If a number of
	// services working together in the same backend coordinate to ensure their error
	// definitions do not overlap, they can share an Issuer value.
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// code is the API Error code.
	Code uint32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	// grpc_code is the grpc status code associated with this error.
	GrpcCode int32 `protobuf:"varint,4,opt,name=grpc_code,json=grpcCode,proto3" json:"grpc_code,omitempty"`
	// name is the human-readable API Error name tied to code.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// message is the error message.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// time is the time that the error occurred
	Time *timestamp.Timestamp `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	// details are arbitrary information about the error.
	Details []*any.Any `protobuf:"bytes,8,rep,name=details,proto3" json:"details,omitempty"`
	// Trace holds a stack of TraceInfo objects. Each time an app encounters an error,
	// it can add it's TraceInfo object to the end of the trace.
	Trace []*TraceInfo `protobuf:"bytes,9,rep,name=trace,proto3" json:"trace,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpeakec_proto_pkerr_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_grpeakec_proto_pkerr_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_grpeakec_proto_pkerr_error_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetId() *cerealMessages.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Error) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetGrpcCode() int32 {
	if x != nil {
		return x.GrpcCode
	}
	return 0
}

func (x *Error) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Error) GetDetails() []*any.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Error) GetTrace() []*TraceInfo {
	if x != nil {
		return x.Trace
	}
	return nil
}

var File_grpeakec_proto_pkerr_error_proto protoreflect.FileDescriptor

var file_grpeakec_proto_pkerr_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6b, 0x65, 0x72, 0x72, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x63, 0x1a, 0x17, 0x63, 0x65,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0xa7, 0x02, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x65,
	0x72, 0x65, 0x61, 0x6c, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x63, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65,
	0x61, 0x6b, 0x65, 0x31, 0x30, 0x30, 0x2f, 0x67, 0x52, 0x50, 0x45, 0x41, 0x4b, 0x45, 0x43, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x65, 0x72, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpeakec_proto_pkerr_error_proto_rawDescOnce sync.Once
	file_grpeakec_proto_pkerr_error_proto_rawDescData = file_grpeakec_proto_pkerr_error_proto_rawDesc
)

func file_grpeakec_proto_pkerr_error_proto_rawDescGZIP() []byte {
	file_grpeakec_proto_pkerr_error_proto_rawDescOnce.Do(func() {
		file_grpeakec_proto_pkerr_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpeakec_proto_pkerr_error_proto_rawDescData)
	})
	return file_grpeakec_proto_pkerr_error_proto_rawDescData
}

var file_grpeakec_proto_pkerr_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpeakec_proto_pkerr_error_proto_goTypes = []interface{}{
	(*TraceInfo)(nil),           // 0: grpeakec.TraceInfo
	(*Error)(nil),               // 1: grpeakec.Error
	(*cerealMessages.UUID)(nil), // 2: cereal.UUID
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*any.Any)(nil),             // 4: google.protobuf.Any
}
var file_grpeakec_proto_pkerr_error_proto_depIdxs = []int32{
	2, // 0: grpeakec.Error.id:type_name -> cereal.UUID
	3, // 1: grpeakec.Error.time:type_name -> google.protobuf.Timestamp
	4, // 2: grpeakec.Error.details:type_name -> google.protobuf.Any
	0, // 3: grpeakec.Error.trace:type_name -> grpeakec.TraceInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grpeakec_proto_pkerr_error_proto_init() }
func file_grpeakec_proto_pkerr_error_proto_init() {
	if File_grpeakec_proto_pkerr_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpeakec_proto_pkerr_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceInfo); i {
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
		file_grpeakec_proto_pkerr_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_grpeakec_proto_pkerr_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpeakec_proto_pkerr_error_proto_goTypes,
		DependencyIndexes: file_grpeakec_proto_pkerr_error_proto_depIdxs,
		MessageInfos:      file_grpeakec_proto_pkerr_error_proto_msgTypes,
	}.Build()
	File_grpeakec_proto_pkerr_error_proto = out.File
	file_grpeakec_proto_pkerr_error_proto_rawDesc = nil
	file_grpeakec_proto_pkerr_error_proto_goTypes = nil
	file_grpeakec_proto_pkerr_error_proto_depIdxs = nil
}
