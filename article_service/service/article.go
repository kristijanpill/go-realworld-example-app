package service

import (
	"context"
	"errors"
	"log"

	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"gorm.io/gorm"
)

type ArticleService struct {
	articleStore store.ArticleStore
	tagStore store.TagStore
	profileServiceClient pb.ProfileServiceClient
}

func NewArticleService(articleStore store.ArticleStore, tagStore store.TagStore, profileServiceClient pb.ProfileServiceClient) *ArticleService {
	return &ArticleService{
		articleStore: articleStore,
		tagStore: tagStore,
		profileServiceClient: profileServiceClient,
	}
}

func (service *ArticleService) CreateArticle (ctx context.Context, request *pb.NewArticleRequest) (*pb.SingleArticleResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	profile, err := service.getProfileById(currentUserIdString)
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	for _, tagName := range request.Article.TagList {
		tag, err := service.tagStore.FindByName(tagName)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newTag := model.NewTag(tagName)
				newTag, err := service.tagStore.Create(newTag)
				log.Println(newTag.ID)
				if err != nil {
					return nil, err
				}
				tags = append(tags, newTag)
			} else {
				return nil, err
			}
		} else {
			tags = append(tags, tag)
		}
	}

	article, err := model.NewArticle(currentUserIdString, request.Article.Title, request.Article.Description, request.Article.Body, tags)
	if err != nil {
		return nil, err
	}

	article, err = service.articleStore.Create(article)
	if err != nil {
		return nil, err
	}

	return &pb.SingleArticleResponse{
		Article: &pb.Article{
			Slug: article.Slug,
			Title: article.Title,
			Description: article.Description,
			Body: article.Body,
			TagList: request.Article.TagList,
			CreatedAt: article.CreatedAt.UTC().String(),
			UpdatedAt: article.UpdatedAt.UTC().String(),
			Favorited: false,
			FavoritesCount: 0,
			Author: &pb.Profile{
				Username: profile.Username,
				Bio: profile.Bio,
				Image: profile.Image,
				Following: false,
			},
		},
	}, nil
}

func (service *ArticleService) getProfileById(id string) (*pb.ProfileInfo, error) {
	return service.profileServiceClient.GetProfileById(context.Background(), &pb.ProfileIdRequest{Id: id})
}