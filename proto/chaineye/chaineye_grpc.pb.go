// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: savourrpc/chaineye.proto

package chaineye

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

const (
	ChaineyeService_GetArticleCat_FullMethodName    = "/savourrpc.chaineye.ChaineyeService/getArticleCat"
	ChaineyeService_GetArticleList_FullMethodName   = "/savourrpc.chaineye.ChaineyeService/getArticleList"
	ChaineyeService_GetArticleDetail_FullMethodName = "/savourrpc.chaineye.ChaineyeService/getArticleDetail"
	ChaineyeService_GetCommentList_FullMethodName   = "/savourrpc.chaineye.ChaineyeService/getCommentList"
	ChaineyeService_GetLikeAddress_FullMethodName   = "/savourrpc.chaineye.ChaineyeService/getLikeAddress"
	ChaineyeService_LikeArticle_FullMethodName      = "/savourrpc.chaineye.ChaineyeService/likeArticle"
)

// ChaineyeServiceClient is the client API for ChaineyeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaineyeServiceClient interface {
	GetArticleCat(ctx context.Context, in *ArticleCatReq, opts ...grpc.CallOption) (*ArticleCatRep, error)
	GetArticleList(ctx context.Context, in *ArticleListReq, opts ...grpc.CallOption) (*ArticleListRep, error)
	GetArticleDetail(ctx context.Context, in *ArticleDetailReq, opts ...grpc.CallOption) (*ArticleDetailRep, error)
	GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRep, error)
	GetLikeAddress(ctx context.Context, in *AddressReq, opts ...grpc.CallOption) (*AddressRep, error)
	LikeArticle(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRep, error)
}

type chaineyeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChaineyeServiceClient(cc grpc.ClientConnInterface) ChaineyeServiceClient {
	return &chaineyeServiceClient{cc}
}

func (c *chaineyeServiceClient) GetArticleCat(ctx context.Context, in *ArticleCatReq, opts ...grpc.CallOption) (*ArticleCatRep, error) {
	out := new(ArticleCatRep)
	err := c.cc.Invoke(ctx, ChaineyeService_GetArticleCat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaineyeServiceClient) GetArticleList(ctx context.Context, in *ArticleListReq, opts ...grpc.CallOption) (*ArticleListRep, error) {
	out := new(ArticleListRep)
	err := c.cc.Invoke(ctx, ChaineyeService_GetArticleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaineyeServiceClient) GetArticleDetail(ctx context.Context, in *ArticleDetailReq, opts ...grpc.CallOption) (*ArticleDetailRep, error) {
	out := new(ArticleDetailRep)
	err := c.cc.Invoke(ctx, ChaineyeService_GetArticleDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaineyeServiceClient) GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRep, error) {
	out := new(CommentListRep)
	err := c.cc.Invoke(ctx, ChaineyeService_GetCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaineyeServiceClient) GetLikeAddress(ctx context.Context, in *AddressReq, opts ...grpc.CallOption) (*AddressRep, error) {
	out := new(AddressRep)
	err := c.cc.Invoke(ctx, ChaineyeService_GetLikeAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaineyeServiceClient) LikeArticle(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRep, error) {
	out := new(LikeRep)
	err := c.cc.Invoke(ctx, ChaineyeService_LikeArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaineyeServiceServer is the server API for ChaineyeService service.
// All implementations must embed UnimplementedChaineyeServiceServer
// for forward compatibility
type ChaineyeServiceServer interface {
	GetArticleCat(context.Context, *ArticleCatReq) (*ArticleCatRep, error)
	GetArticleList(context.Context, *ArticleListReq) (*ArticleListRep, error)
	GetArticleDetail(context.Context, *ArticleDetailReq) (*ArticleDetailRep, error)
	GetCommentList(context.Context, *CommentListReq) (*CommentListRep, error)
	GetLikeAddress(context.Context, *AddressReq) (*AddressRep, error)
	LikeArticle(context.Context, *LikeReq) (*LikeRep, error)
	mustEmbedUnimplementedChaineyeServiceServer()
}

// UnimplementedChaineyeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChaineyeServiceServer struct {
}

func (UnimplementedChaineyeServiceServer) GetArticleCat(context.Context, *ArticleCatReq) (*ArticleCatRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleCat not implemented")
}
func (UnimplementedChaineyeServiceServer) GetArticleList(context.Context, *ArticleListReq) (*ArticleListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedChaineyeServiceServer) GetArticleDetail(context.Context, *ArticleDetailReq) (*ArticleDetailRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleDetail not implemented")
}
func (UnimplementedChaineyeServiceServer) GetCommentList(context.Context, *CommentListReq) (*CommentListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedChaineyeServiceServer) GetLikeAddress(context.Context, *AddressReq) (*AddressRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikeAddress not implemented")
}
func (UnimplementedChaineyeServiceServer) LikeArticle(context.Context, *LikeReq) (*LikeRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeArticle not implemented")
}
func (UnimplementedChaineyeServiceServer) mustEmbedUnimplementedChaineyeServiceServer() {}

// UnsafeChaineyeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaineyeServiceServer will
// result in compilation errors.
type UnsafeChaineyeServiceServer interface {
	mustEmbedUnimplementedChaineyeServiceServer()
}

func RegisterChaineyeServiceServer(s grpc.ServiceRegistrar, srv ChaineyeServiceServer) {
	s.RegisterService(&ChaineyeService_ServiceDesc, srv)
}

func _ChaineyeService_GetArticleCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).GetArticleCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_GetArticleCat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).GetArticleCat(ctx, req.(*ArticleCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaineyeService_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_GetArticleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).GetArticleList(ctx, req.(*ArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaineyeService_GetArticleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).GetArticleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_GetArticleDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).GetArticleDetail(ctx, req.(*ArticleDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaineyeService_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_GetCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).GetCommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaineyeService_GetLikeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).GetLikeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_GetLikeAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).GetLikeAddress(ctx, req.(*AddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaineyeService_LikeArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaineyeServiceServer).LikeArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaineyeService_LikeArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaineyeServiceServer).LikeArticle(ctx, req.(*LikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChaineyeService_ServiceDesc is the grpc.ServiceDesc for ChaineyeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaineyeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "savourrpc.chaineye.ChaineyeService",
	HandlerType: (*ChaineyeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getArticleCat",
			Handler:    _ChaineyeService_GetArticleCat_Handler,
		},
		{
			MethodName: "getArticleList",
			Handler:    _ChaineyeService_GetArticleList_Handler,
		},
		{
			MethodName: "getArticleDetail",
			Handler:    _ChaineyeService_GetArticleDetail_Handler,
		},
		{
			MethodName: "getCommentList",
			Handler:    _ChaineyeService_GetCommentList_Handler,
		},
		{
			MethodName: "getLikeAddress",
			Handler:    _ChaineyeService_GetLikeAddress_Handler,
		},
		{
			MethodName: "likeArticle",
			Handler:    _ChaineyeService_LikeArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "savourrpc/chaineye.proto",
}
