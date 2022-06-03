package service

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc/metadata"
)

type FavoriteService struct {
	favoriteStore store.FavoriteStore
	articleStore store.ArticleStore
	profileServiceClient pb.ProfileServiceClient
}

func NewFavoriteService(favoriteStore store.FavoriteStore, articleStore store.ArticleStore, profileServiceClient pb.ProfileServiceClient) *FavoriteService {
	return &FavoriteService{
		favoriteStore: favoriteStore,
		articleStore: articleStore,
		profileServiceClient: profileServiceClient,
	}
}

func (service *FavoriteService) CreateArticleFavorite(ctx context.Context, request *pb.CreateArticleFavoriteRequest) (*pb.SingleArticleResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	article, err := service.articleStore.FindBySlug(request.Slug)
	if err != nil {
		return nil, err
	}

	favorite, err := model.NewFavorite(currentUserIdString, article)
	if err != nil {
		return nil, err
	}

	_, err = service.favoriteStore.Create(favorite)
	if err != nil {
		return nil, err
	}

	author, err := service.getProfileById(ctx, article.UserID.String())
	if err != nil {
		return nil, err
	}

	var tagList []string
	for _, tag := range article.Tags {
		tagList = append(tagList, tag.Name)
	}

	return &pb.SingleArticleResponse{
		Article: &pb.Article{
			Slug: article.Slug,
			Title: article.Title,
			Description: article.Description,
			Body: article.Body,
			TagList: tagList,
			CreatedAt: article.CreatedAt.UTC().String(),
			UpdatedAt: article.UpdatedAt.UTC().String(),
			Favorited: true,
			FavoritesCount: 1,
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		},
	}, nil
}

func (service *FavoriteService) DeleteArticleFavorite(ctx context.Context, request *pb.DeleteArticleFavoriteRequest) (*pb.SingleArticleResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	favorite, err := service.favoriteStore.FindByUserIdAndSlug(currentUserIdString, request.Slug)
	if err != nil {
		return nil, err
	}

	err = service.favoriteStore.Delete(favorite)
	if err != nil {
		return nil, err
	}

	author, err := service.getProfileById(ctx, favorite.Article.UserID.String())
	if err != nil {
		return nil, err
	}

	var tagList []string
	for _, tag := range favorite.Article.Tags {
		tagList = append(tagList, tag.Name)
	}

	return &pb.SingleArticleResponse{
		Article: &pb.Article{
			Slug: favorite.Article.Slug,
			Title: favorite.Article.Title,
			Description: favorite.Article.Description,
			Body: favorite.Article.Body,
			TagList: tagList,
			CreatedAt: favorite.Article.CreatedAt.UTC().String(),
			UpdatedAt: favorite.Article.UpdatedAt.UTC().String(),
			Favorited: false,
			FavoritesCount: 0,
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		},
	}, nil
}

func (service *FavoriteService) getProfileById(ctx context.Context, id string) (*pb.ProfileResponse, error) {
	if ctx.Value(interceptor.TokenKey{}) != nil {
		md := metadata.New(map[string]string{"Authorization": "Token " + ctx.Value(interceptor.TokenKey{}).(string)})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return service.profileServiceClient.GetProfileById(ctx, &pb.ProfileIdRequest{Id: id})
}