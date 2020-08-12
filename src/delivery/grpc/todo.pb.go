// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: src/delivery/grpc/todo.proto

package todo_grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodoListRequest) Reset() {
	*x = TodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListRequest) ProtoMessage() {}

func (x *TodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListRequest.ProtoReflect.Descriptor instead.
func (*TodoListRequest) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_todo_proto_rawDescGZIP(), []int{1}
}

type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoListResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_src_delivery_grpc_todo_proto protoreflect.FileDescriptor

var file_src_delivery_grpc_todo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x72, 0x63, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x70, 0x64, 0x62, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x5f, 0x70, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x32, 0x50, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x5f, 0x70, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x70,
	0x64, 0x62, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_delivery_grpc_todo_proto_rawDescOnce sync.Once
	file_src_delivery_grpc_todo_proto_rawDescData = file_src_delivery_grpc_todo_proto_rawDesc
)

func file_src_delivery_grpc_todo_proto_rawDescGZIP() []byte {
	file_src_delivery_grpc_todo_proto_rawDescOnce.Do(func() {
		file_src_delivery_grpc_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_delivery_grpc_todo_proto_rawDescData)
	})
	return file_src_delivery_grpc_todo_proto_rawDescData
}

var file_src_delivery_grpc_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_delivery_grpc_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),             // 0: todo_pdb.Todo
	(*TodoListRequest)(nil),  // 1: todo_pdb.TodoListRequest
	(*TodoListResponse)(nil), // 2: todo_pdb.TodoListResponse
}
var file_src_delivery_grpc_todo_proto_depIdxs = []int32{
	0, // 0: todo_pdb.TodoListResponse.todos:type_name -> todo_pdb.Todo
	1, // 1: todo_pdb.TodoService.TodoList:input_type -> todo_pdb.TodoListRequest
	2, // 2: todo_pdb.TodoService.TodoList:output_type -> todo_pdb.TodoListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_delivery_grpc_todo_proto_init() }
func file_src_delivery_grpc_todo_proto_init() {
	if File_src_delivery_grpc_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_delivery_grpc_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_src_delivery_grpc_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListRequest); i {
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
		file_src_delivery_grpc_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListResponse); i {
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
			RawDescriptor: file_src_delivery_grpc_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_delivery_grpc_todo_proto_goTypes,
		DependencyIndexes: file_src_delivery_grpc_todo_proto_depIdxs,
		MessageInfos:      file_src_delivery_grpc_todo_proto_msgTypes,
	}.Build()
	File_src_delivery_grpc_todo_proto = out.File
	file_src_delivery_grpc_todo_proto_rawDesc = nil
	file_src_delivery_grpc_todo_proto_goTypes = nil
	file_src_delivery_grpc_todo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	TodoList(ctx context.Context, in *TodoListRequest, opts ...grpc.CallOption) (*TodoListResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) TodoList(ctx context.Context, in *TodoListRequest, opts ...grpc.CallOption) (*TodoListResponse, error) {
	out := new(TodoListResponse)
	err := c.cc.Invoke(ctx, "/todo_pdb.TodoService/TodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	TodoList(context.Context, *TodoListRequest) (*TodoListResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) TodoList(context.Context, *TodoListRequest) (*TodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TodoList not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_TodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).TodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo_pdb.TodoService/TodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).TodoList(ctx, req.(*TodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo_pdb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TodoList",
			Handler:    _TodoService_TodoList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/delivery/grpc/todo.proto",
}