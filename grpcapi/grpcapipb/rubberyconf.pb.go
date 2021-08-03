// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: rubberyconf.proto

package grpcapipb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RubberyConfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RubberyConfRequest) Reset() {
	*x = RubberyConfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rubberyconf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RubberyConfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RubberyConfRequest) ProtoMessage() {}

func (x *RubberyConfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rubberyconf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RubberyConfRequest.ProtoReflect.Descriptor instead.
func (*RubberyConfRequest) Descriptor() ([]byte, []int) {
	return file_rubberyconf_proto_rawDescGZIP(), []int{0}
}

func (x *RubberyConfRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RubberyConfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *RubberyConfResponse) Reset() {
	*x = RubberyConfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rubberyconf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RubberyConfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RubberyConfResponse) ProtoMessage() {}

func (x *RubberyConfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rubberyconf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RubberyConfResponse.ProtoReflect.Descriptor instead.
func (*RubberyConfResponse) Descriptor() ([]byte, []int) {
	return file_rubberyconf_proto_rawDescGZIP(), []int{1}
}

func (x *RubberyConfResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_rubberyconf_proto protoreflect.FileDescriptor

var file_rubberyconf_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x28, 0x0a, 0x12,
	0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0xe6, 0x01, 0x0a, 0x12, 0x72, 0x75,
	0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x62,
	0x62, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rubberyconf_proto_rawDescOnce sync.Once
	file_rubberyconf_proto_rawDescData = file_rubberyconf_proto_rawDesc
)

func file_rubberyconf_proto_rawDescGZIP() []byte {
	file_rubberyconf_proto_rawDescOnce.Do(func() {
		file_rubberyconf_proto_rawDescData = protoimpl.X.CompressGZIP(file_rubberyconf_proto_rawDescData)
	})
	return file_rubberyconf_proto_rawDescData
}

var file_rubberyconf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rubberyconf_proto_goTypes = []interface{}{
	(*RubberyConfRequest)(nil),  // 0: grpcapi.rubberyConfRequest
	(*RubberyConfResponse)(nil), // 1: grpcapi.rubberyConfResponse
}
var file_rubberyconf_proto_depIdxs = []int32{
	0, // 0: grpcapi.rubberyConfService.get:input_type -> grpcapi.rubberyConfRequest
	0, // 1: grpcapi.rubberyConfService.create:input_type -> grpcapi.rubberyConfRequest
	0, // 2: grpcapi.rubberyConfService.delete:input_type -> grpcapi.rubberyConfRequest
	1, // 3: grpcapi.rubberyConfService.get:output_type -> grpcapi.rubberyConfResponse
	1, // 4: grpcapi.rubberyConfService.create:output_type -> grpcapi.rubberyConfResponse
	1, // 5: grpcapi.rubberyConfService.delete:output_type -> grpcapi.rubberyConfResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rubberyconf_proto_init() }
func file_rubberyconf_proto_init() {
	if File_rubberyconf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rubberyconf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RubberyConfRequest); i {
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
		file_rubberyconf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RubberyConfResponse); i {
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
			RawDescriptor: file_rubberyconf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rubberyconf_proto_goTypes,
		DependencyIndexes: file_rubberyconf_proto_depIdxs,
		MessageInfos:      file_rubberyconf_proto_msgTypes,
	}.Build()
	File_rubberyconf_proto = out.File
	file_rubberyconf_proto_rawDesc = nil
	file_rubberyconf_proto_goTypes = nil
	file_rubberyconf_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RubberyConfServiceClient is the client API for RubberyConfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RubberyConfServiceClient interface {
	Get(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error)
	Create(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error)
	Delete(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error)
}

type rubberyConfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRubberyConfServiceClient(cc grpc.ClientConnInterface) RubberyConfServiceClient {
	return &rubberyConfServiceClient{cc}
}

func (c *rubberyConfServiceClient) Get(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error) {
	out := new(RubberyConfResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.rubberyConfService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rubberyConfServiceClient) Create(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error) {
	out := new(RubberyConfResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.rubberyConfService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rubberyConfServiceClient) Delete(ctx context.Context, in *RubberyConfRequest, opts ...grpc.CallOption) (*RubberyConfResponse, error) {
	out := new(RubberyConfResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.rubberyConfService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RubberyConfServiceServer is the server API for RubberyConfService service.
type RubberyConfServiceServer interface {
	Get(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error)
	Create(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error)
	Delete(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error)
}

// UnimplementedRubberyConfServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRubberyConfServiceServer struct {
}

func (*UnimplementedRubberyConfServiceServer) Get(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRubberyConfServiceServer) Create(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRubberyConfServiceServer) Delete(context.Context, *RubberyConfRequest) (*RubberyConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRubberyConfServiceServer(s *grpc.Server, srv RubberyConfServiceServer) {
	s.RegisterService(&_RubberyConfService_serviceDesc, srv)
}

func _RubberyConfService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RubberyConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubberyConfServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.rubberyConfService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubberyConfServiceServer).Get(ctx, req.(*RubberyConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RubberyConfService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RubberyConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubberyConfServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.rubberyConfService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubberyConfServiceServer).Create(ctx, req.(*RubberyConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RubberyConfService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RubberyConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubberyConfServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.rubberyConfService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubberyConfServiceServer).Delete(ctx, req.(*RubberyConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RubberyConfService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.rubberyConfService",
	HandlerType: (*RubberyConfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _RubberyConfService_Get_Handler,
		},
		{
			MethodName: "create",
			Handler:    _RubberyConfService_Create_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _RubberyConfService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rubberyconf.proto",
}
