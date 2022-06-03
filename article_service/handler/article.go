package handler

import (
	"context"
	"log"

	"github.com/kristijanpill/go-realworld-example-app/article_service/service"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
)

type ArticleHandler struct {
	pb.UnimplementedArticleServiceServer
	articleService *service.ArticleService
	tagService *service.TagService
}

func NewArticleHandler(articleService *service.ArticleService, tagService *service.TagService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
		tagService: tagService,
	}
}

func (handler *ArticleHandler) CreateArticle(ctx context.Context, request *pb.NewArticleRequest) (*pb.SingleArticleResponse, error) {
	log.Println("ARTICLE HANDLER")
	log.Println(request)
	log.Println(request.Article)
	return handler.articleService.CreateArticle(ctx, request);
}