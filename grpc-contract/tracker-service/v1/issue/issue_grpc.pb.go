// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package issue

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

// IssueServiceClient is the client API for IssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IssueServiceClient interface {
	GetIssueById(ctx context.Context, in *IssueByIdRequest, opts ...grpc.CallOption) (*IssueResponse, error)
	GetIssuesByProject(ctx context.Context, in *IssuesByProjectRequest, opts ...grpc.CallOption) (*IssuesResponse, error)
	GetIssuesByUser(ctx context.Context, in *IssuesByUserRequest, opts ...grpc.CallOption) (*IssuesResponse, error)
	CreateIssue(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	DeleteIssue(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	UpdateDescription(ctx context.Context, in *UpdateDescriptionRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	UpdateBugTrace(ctx context.Context, in *UpdateBugTraceRequest, opts ...grpc.CallOption) (*GenericResponse, error)
}

type issueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueServiceClient(cc grpc.ClientConnInterface) IssueServiceClient {
	return &issueServiceClient{cc}
}

func (c *issueServiceClient) GetIssueById(ctx context.Context, in *IssueByIdRequest, opts ...grpc.CallOption) (*IssueResponse, error) {
	out := new(IssueResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/GetIssueById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) GetIssuesByProject(ctx context.Context, in *IssuesByProjectRequest, opts ...grpc.CallOption) (*IssuesResponse, error) {
	out := new(IssuesResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/GetIssuesByProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) GetIssuesByUser(ctx context.Context, in *IssuesByUserRequest, opts ...grpc.CallOption) (*IssuesResponse, error) {
	out := new(IssuesResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/GetIssuesByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) CreateIssue(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/CreateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) DeleteIssue(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/DeleteIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) UpdateDescription(ctx context.Context, in *UpdateDescriptionRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/UpdateDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) UpdateBugTrace(ctx context.Context, in *UpdateBugTraceRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/issue.IssueService/UpdateBugTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueServiceServer is the server API for IssueService service.
// All implementations must embed UnimplementedIssueServiceServer
// for forward compatibility
type IssueServiceServer interface {
	GetIssueById(context.Context, *IssueByIdRequest) (*IssueResponse, error)
	GetIssuesByProject(context.Context, *IssuesByProjectRequest) (*IssuesResponse, error)
	GetIssuesByUser(context.Context, *IssuesByUserRequest) (*IssuesResponse, error)
	CreateIssue(context.Context, *CreateRequest) (*GenericResponse, error)
	DeleteIssue(context.Context, *DeleteRequest) (*GenericResponse, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*GenericResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*GenericResponse, error)
	UpdateDescription(context.Context, *UpdateDescriptionRequest) (*GenericResponse, error)
	UpdateBugTrace(context.Context, *UpdateBugTraceRequest) (*GenericResponse, error)
	mustEmbedUnimplementedIssueServiceServer()
}

// UnimplementedIssueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIssueServiceServer struct {
}

func (UnimplementedIssueServiceServer) GetIssueById(context.Context, *IssueByIdRequest) (*IssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssueById not implemented")
}
func (UnimplementedIssueServiceServer) GetIssuesByProject(context.Context, *IssuesByProjectRequest) (*IssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssuesByProject not implemented")
}
func (UnimplementedIssueServiceServer) GetIssuesByUser(context.Context, *IssuesByUserRequest) (*IssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssuesByUser not implemented")
}
func (UnimplementedIssueServiceServer) CreateIssue(context.Context, *CreateRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssue not implemented")
}
func (UnimplementedIssueServiceServer) DeleteIssue(context.Context, *DeleteRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIssue not implemented")
}
func (UnimplementedIssueServiceServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedIssueServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedIssueServiceServer) UpdateDescription(context.Context, *UpdateDescriptionRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDescription not implemented")
}
func (UnimplementedIssueServiceServer) UpdateBugTrace(context.Context, *UpdateBugTraceRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBugTrace not implemented")
}
func (UnimplementedIssueServiceServer) mustEmbedUnimplementedIssueServiceServer() {}

// UnsafeIssueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IssueServiceServer will
// result in compilation errors.
type UnsafeIssueServiceServer interface {
	mustEmbedUnimplementedIssueServiceServer()
}

func RegisterIssueServiceServer(s grpc.ServiceRegistrar, srv IssueServiceServer) {
	s.RegisterService(&IssueService_ServiceDesc, srv)
}

func _IssueService_GetIssueById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).GetIssueById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/GetIssueById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).GetIssueById(ctx, req.(*IssueByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_GetIssuesByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuesByProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).GetIssuesByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/GetIssuesByProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).GetIssuesByProject(ctx, req.(*IssuesByProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_GetIssuesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).GetIssuesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/GetIssuesByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).GetIssuesByUser(ctx, req.(*IssuesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_CreateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).CreateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/CreateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).CreateIssue(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_DeleteIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).DeleteIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/DeleteIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).DeleteIssue(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_UpdateDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).UpdateDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/UpdateDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).UpdateDescription(ctx, req.(*UpdateDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_UpdateBugTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBugTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).UpdateBugTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/UpdateBugTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).UpdateBugTrace(ctx, req.(*UpdateBugTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IssueService_ServiceDesc is the grpc.ServiceDesc for IssueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IssueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "issue.IssueService",
	HandlerType: (*IssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIssueById",
			Handler:    _IssueService_GetIssueById_Handler,
		},
		{
			MethodName: "GetIssuesByProject",
			Handler:    _IssueService_GetIssuesByProject_Handler,
		},
		{
			MethodName: "GetIssuesByUser",
			Handler:    _IssueService_GetIssuesByUser_Handler,
		},
		{
			MethodName: "CreateIssue",
			Handler:    _IssueService_CreateIssue_Handler,
		},
		{
			MethodName: "DeleteIssue",
			Handler:    _IssueService_DeleteIssue_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _IssueService_UpdateStatus_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _IssueService_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateDescription",
			Handler:    _IssueService_UpdateDescription_Handler,
		},
		{
			MethodName: "UpdateBugTrace",
			Handler:    _IssueService_UpdateBugTrace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-contract/tracker-service/v1/issue/issue.proto",
}
