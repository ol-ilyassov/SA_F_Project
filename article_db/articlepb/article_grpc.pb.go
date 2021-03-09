// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package articlepb

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

// ArticlesServiceClient is the client API for ArticlesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesServiceClient interface {
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (ArticlesService_GetArticlesClient, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error)
	InsertArticle(ctx context.Context, in *InsertArticleRequest, opts ...grpc.CallOption) (*InsertArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
	SearchArticles(ctx context.Context, in *SearchArticlesRequest, opts ...grpc.CallOption) (ArticlesService_SearchArticlesClient, error)
	EditArticle(ctx context.Context, in *EditArticleRequest, opts ...grpc.CallOption) (*EditArticleResponse, error)
}

type articlesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesServiceClient(cc grpc.ClientConnInterface) ArticlesServiceClient {
	return &articlesServiceClient{cc}
}

func (c *articlesServiceClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (ArticlesService_GetArticlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArticlesService_ServiceDesc.Streams[0], "/articlepb.ArticlesService/GetArticles", opts...)
	if err != nil {
		return nil, err
	}
	x := &articlesServiceGetArticlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArticlesService_GetArticlesClient interface {
	Recv() (*GetArticlesResponse, error)
	grpc.ClientStream
}

type articlesServiceGetArticlesClient struct {
	grpc.ClientStream
}

func (x *articlesServiceGetArticlesClient) Recv() (*GetArticlesResponse, error) {
	m := new(GetArticlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *articlesServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	out := new(GetArticleResponse)
	err := c.cc.Invoke(ctx, "/articlepb.ArticlesService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesServiceClient) InsertArticle(ctx context.Context, in *InsertArticleRequest, opts ...grpc.CallOption) (*InsertArticleResponse, error) {
	out := new(InsertArticleResponse)
	err := c.cc.Invoke(ctx, "/articlepb.ArticlesService/InsertArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, "/articlepb.ArticlesService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesServiceClient) SearchArticles(ctx context.Context, in *SearchArticlesRequest, opts ...grpc.CallOption) (ArticlesService_SearchArticlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArticlesService_ServiceDesc.Streams[1], "/articlepb.ArticlesService/SearchArticles", opts...)
	if err != nil {
		return nil, err
	}
	x := &articlesServiceSearchArticlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArticlesService_SearchArticlesClient interface {
	Recv() (*SearchArticlesResponse, error)
	grpc.ClientStream
}

type articlesServiceSearchArticlesClient struct {
	grpc.ClientStream
}

func (x *articlesServiceSearchArticlesClient) Recv() (*SearchArticlesResponse, error) {
	m := new(SearchArticlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *articlesServiceClient) EditArticle(ctx context.Context, in *EditArticleRequest, opts ...grpc.CallOption) (*EditArticleResponse, error) {
	out := new(EditArticleResponse)
	err := c.cc.Invoke(ctx, "/articlepb.ArticlesService/EditArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServiceServer is the server API for ArticlesService service.
// All implementations must embed UnimplementedArticlesServiceServer
// for forward compatibility
type ArticlesServiceServer interface {
	GetArticles(*GetArticlesRequest, ArticlesService_GetArticlesServer) error
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error)
	InsertArticle(context.Context, *InsertArticleRequest) (*InsertArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	SearchArticles(*SearchArticlesRequest, ArticlesService_SearchArticlesServer) error
	EditArticle(context.Context, *EditArticleRequest) (*EditArticleResponse, error)
	mustEmbedUnimplementedArticlesServiceServer()
}

// UnimplementedArticlesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServiceServer struct {
}

func (UnimplementedArticlesServiceServer) GetArticles(*GetArticlesRequest, ArticlesService_GetArticlesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedArticlesServiceServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticlesServiceServer) InsertArticle(context.Context, *InsertArticleRequest) (*InsertArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertArticle not implemented")
}
func (UnimplementedArticlesServiceServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticlesServiceServer) SearchArticles(*SearchArticlesRequest, ArticlesService_SearchArticlesServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchArticles not implemented")
}
func (UnimplementedArticlesServiceServer) EditArticle(context.Context, *EditArticleRequest) (*EditArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditArticle not implemented")
}
func (UnimplementedArticlesServiceServer) mustEmbedUnimplementedArticlesServiceServer() {}

// UnsafeArticlesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServiceServer will
// result in compilation errors.
type UnsafeArticlesServiceServer interface {
	mustEmbedUnimplementedArticlesServiceServer()
}

func RegisterArticlesServiceServer(s grpc.ServiceRegistrar, srv ArticlesServiceServer) {
	s.RegisterService(&ArticlesService_ServiceDesc, srv)
}

func _ArticlesService_GetArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetArticlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArticlesServiceServer).GetArticles(m, &articlesServiceGetArticlesServer{stream})
}

type ArticlesService_GetArticlesServer interface {
	Send(*GetArticlesResponse) error
	grpc.ServerStream
}

type articlesServiceGetArticlesServer struct {
	grpc.ServerStream
}

func (x *articlesServiceGetArticlesServer) Send(m *GetArticlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ArticlesService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlepb.ArticlesService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesService_InsertArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServiceServer).InsertArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlepb.ArticlesService/InsertArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServiceServer).InsertArticle(ctx, req.(*InsertArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlepb.ArticlesService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServiceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesService_SearchArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchArticlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArticlesServiceServer).SearchArticles(m, &articlesServiceSearchArticlesServer{stream})
}

type ArticlesService_SearchArticlesServer interface {
	Send(*SearchArticlesResponse) error
	grpc.ServerStream
}

type articlesServiceSearchArticlesServer struct {
	grpc.ServerStream
}

func (x *articlesServiceSearchArticlesServer) Send(m *SearchArticlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ArticlesService_EditArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServiceServer).EditArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlepb.ArticlesService/EditArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServiceServer).EditArticle(ctx, req.(*EditArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticlesService_ServiceDesc is the grpc.ServiceDesc for ArticlesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticlesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "articlepb.ArticlesService",
	HandlerType: (*ArticlesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticle",
			Handler:    _ArticlesService_GetArticle_Handler,
		},
		{
			MethodName: "InsertArticle",
			Handler:    _ArticlesService_InsertArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticlesService_DeleteArticle_Handler,
		},
		{
			MethodName: "EditArticle",
			Handler:    _ArticlesService_EditArticle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArticles",
			Handler:       _ArticlesService_GetArticles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchArticles",
			Handler:       _ArticlesService_SearchArticles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "article_db/articlepb/article.proto",
}
