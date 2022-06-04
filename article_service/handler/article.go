package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/article_service/service"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleHandler struct {
	pb.UnimplementedArticleServiceServer
	articleService *service.ArticleService
	tagService *service.TagService
	favoriteService *service.FavoriteService
	commentService *service.CommentService
}

func NewArticleHandler(articleService *service.ArticleService, tagService *service.TagService, favoriteService *service.FavoriteService, commentService *service.CommentService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
		tagService: tagService,
		favoriteService: favoriteService,
		commentService: commentService,
	}
}

func (handler* ArticleHandler) GetArticles(ctx context.Context, request *pb.GetArticlesRequest) (*pb.MultipleArticlesResponse, error) {
	return handler.articleService.GetArticles(ctx, request);
}

func (handler *ArticleHandler) CreateArticle(ctx context.Context, request *pb.NewArticleRequest) (*pb.SingleArticleResponse, error) {
	return handler.articleService.CreateArticle(ctx, request);
}

func (handler *ArticleHandler) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.SingleArticleResponse, error) {
	return handler.articleService.GetArticle(ctx, request)
}

func (handler *ArticleHandler) UpdateArticle(ctx context.Context, request *pb.UpdateArticleRequest) (*pb.SingleArticleResponse, error) {
	return handler.articleService.UpdateArticle(ctx, request)
}

func (handler *ArticleHandler) DeleteArticle(ctx context.Context, request *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	return handler.articleService.DeleteArticle(ctx, request)
}

func (handler *ArticleHandler) CreateArticleFavorite(ctx context.Context, request *pb.CreateArticleFavoriteRequest) (*pb.SingleArticleResponse, error) {
	return handler.favoriteService.CreateArticleFavorite(ctx, request)
}

func (handler *ArticleHandler) DeleteArticleFavorite(ctx context.Context, request *pb.DeleteArticleFavoriteRequest) (*pb.SingleArticleResponse, error) {
	return handler.favoriteService.DeleteArticleFavorite(ctx, request)
}

func (handler *ArticleHandler) CreateArticleComment(ctx context.Context, request *pb.NewCommentRequest) (*pb.SingleCommentResponse, error) {
	return handler.commentService.CreateArticleComment(ctx, request)
}