// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: task_service.proto

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error)
	MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error) {
	out := new(TaskStagesResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/TaskStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error) {
	out := new(MemberProjectResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/MemberProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error)
	MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskStages not implemented")
}
func (UnimplementedTaskServiceServer) MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberProjectList not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_TaskStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/TaskStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskStages(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_MemberProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).MemberProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/MemberProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).MemberProjectList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.service.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskStages",
			Handler:    _TaskService_TaskStages_Handler,
		},
		{
			MethodName: "MemberProjectList",
			Handler:    _TaskService_MemberProjectList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_service.proto",
}
