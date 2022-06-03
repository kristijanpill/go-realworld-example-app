// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0--rc1
// source: article_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	GetArticlesFeed(ctx context.Context, in *GetArticlesFeedRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error)
	CreateArticle(ctx context.Context, in *NewArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetArticleComments(ctx context.Context, in *GetArticleCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsResponse, error)
	CreateArticleComment(ctx context.Context, in *NewCommentRequest, opts ...grpc.CallOption) (*SingleCommentResponse, error)
	DeleteArticleComment(ctx context.Context, in *DeleteArticleCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateArticleFavorite(ctx context.Context, in *CreateArticleFavoriteRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	DeleteArticleFavorite(ctx context.Context, in *DeleteArticleFavoriteRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	GetTags(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TagsResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) GetArticlesFeed(ctx context.Context, in *GetArticlesFeedRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error) {
	out := new(MultipleArticlesResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/GetArticlesFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error) {
	out := new(MultipleArticlesResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/GetArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *NewArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/article.ArticleService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleComments(ctx context.Context, in *GetArticleCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsResponse, error) {
	out := new(MultipleCommentsResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/GetArticleComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) CreateArticleComment(ctx context.Context, in *NewCommentRequest, opts ...grpc.CallOption) (*SingleCommentResponse, error) {
	out := new(SingleCommentResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/CreateArticleComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticleComment(ctx context.Context, in *DeleteArticleCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/article.ArticleService/DeleteArticleComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) CreateArticleFavorite(ctx context.Context, in *CreateArticleFavoriteRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/CreateArticleFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticleFavorite(ctx context.Context, in *DeleteArticleFavoriteRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/DeleteArticleFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetTags(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TagsResponse, error) {
	out := new(TagsResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	GetArticlesFeed(context.Context, *GetArticlesFeedRequest) (*MultipleArticlesResponse, error)
	GetArticles(context.Context, *GetArticlesRequest) (*MultipleArticlesResponse, error)
	CreateArticle(context.Context, *NewArticleRequest) (*SingleArticleResponse, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error)
	GetArticleComments(context.Context, *GetArticleCommentsRequest) (*MultipleCommentsResponse, error)
	CreateArticleComment(context.Context, *NewCommentRequest) (*SingleCommentResponse, error)
	DeleteArticleComment(context.Context, *DeleteArticleCommentRequest) (*emptypb.Empty, error)
	CreateArticleFavorite(context.Context, *CreateArticleFavoriteRequest) (*SingleArticleResponse, error)
	DeleteArticleFavorite(context.Context, *DeleteArticleFavoriteRequest) (*SingleArticleResponse, error)
	GetTags(context.Context, *emptypb.Empty) (*TagsResponse, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) GetArticlesFeed(context.Context, *GetArticlesFeedRequest) (*MultipleArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesFeed not implemented")
}
func (UnimplementedArticleServiceServer) GetArticles(context.Context, *GetArticlesRequest) (*MultipleArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *NewArticleRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticle(context.Context, *GetArticleRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleComments(context.Context, *GetArticleCommentsRequest) (*MultipleCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleComments not implemented")
}
func (UnimplementedArticleServiceServer) CreateArticleComment(context.Context, *NewCommentRequest) (*SingleCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleComment not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticleComment(context.Context, *DeleteArticleCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticleComment not implemented")
}
func (UnimplementedArticleServiceServer) CreateArticleFavorite(context.Context, *CreateArticleFavoriteRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleFavorite not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticleFavorite(context.Context, *DeleteArticleFavoriteRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticleFavorite not implemented")
}
func (UnimplementedArticleServiceServer) GetTags(context.Context, *emptypb.Empty) (*TagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_GetArticlesFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticlesFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/GetArticlesFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticlesFeed(ctx, req.(*GetArticlesFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/GetArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticles(ctx, req.(*GetArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*NewArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/GetArticleComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleComments(ctx, req.(*GetArticleCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_CreateArticleComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticleComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/CreateArticleComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticleComment(ctx, req.(*NewCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticleComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticleComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/DeleteArticleComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticleComment(ctx, req.(*DeleteArticleCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_CreateArticleFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticleFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/CreateArticleFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticleFavorite(ctx, req.(*CreateArticleFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticleFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticleFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/DeleteArticleFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticleFavorite(ctx, req.(*DeleteArticleFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetTags(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticlesFeed",
			Handler:    _ArticleService_GetArticlesFeed_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _ArticleService_GetArticles_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ArticleService_GetArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
		{
			MethodName: "GetArticleComments",
			Handler:    _ArticleService_GetArticleComments_Handler,
		},
		{
			MethodName: "CreateArticleComment",
			Handler:    _ArticleService_CreateArticleComment_Handler,
		},
		{
			MethodName: "DeleteArticleComment",
			Handler:    _ArticleService_DeleteArticleComment_Handler,
		},
		{
			MethodName: "CreateArticleFavorite",
			Handler:    _ArticleService_CreateArticleFavorite_Handler,
		},
		{
			MethodName: "DeleteArticleFavorite",
			Handler:    _ArticleService_DeleteArticleFavorite_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _ArticleService_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article_service.proto",
}