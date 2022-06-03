package service

import (
	"context"
	"errors"

	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc/metadata"
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

func (service *ArticleService) GetArticles(ctx context.Context, request *pb.GetArticlesRequest) (*pb.MultipleArticlesResponse, error) {
	limit := request.Limit
	if limit <= 0 {
		limit = 20
	}
	offset := request.Offset
	if offset < 0 {
		offset = 0
	}

	var articles []*model.Article
	var err error

	if request.Tag != "" {
		articles, err = service.findArticlesByTag(request.Tag, offset, limit)
	} else if request.Author != "" {
		articles, err = service.findArticlesByAuthor(request.Author, offset, limit)
	} else if request.Favorited != "" {
		articles, err = service.findArticlesByFavoritedUsername(request.Favorited, offset, limit)
	} else {
		articles, err = service.findArticles(offset, limit)
	}
	if err != nil {
		return nil, err
	}

	response := &pb.MultipleArticlesResponse{
		Articles: []*pb.Article{},
		ArticlesCount: int32(len(articles)),
	}

	for _, articleModel := range articles {
		var tagList []string
		for _, tag := range articleModel.Tags {
			tagList = append(tagList, tag.Name)
		}

		author, err := service.getProfileById(ctx, articleModel.UserID.String())
		if err != nil {
			return nil, err
		}

		article := &pb.Article{
			Slug: articleModel.Slug,
			Title: articleModel.Title,
			Description: articleModel.Description,
			Body: articleModel.Body,
			TagList: tagList,
			CreatedAt: articleModel.CreatedAt.UTC().String(),
			UpdatedAt: articleModel.UpdatedAt.UTC().String(),
			Favorited: service.isFavoritedByUser("TODO"),
			FavoritesCount: 0,
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		}
		response.Articles = append(response.Articles, article)
	}

	return response, nil
}

func (service *ArticleService) CreateArticle (ctx context.Context, request *pb.NewArticleRequest) (*pb.SingleArticleResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	profile, err := service.getProfileById(ctx, currentUserIdString)
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
				Username: profile.Profile.Username,
				Bio: profile.Profile.Bio,
				Image: profile.Profile.Image,
				Following: profile.Profile.Following,
			},
		},
	}, nil
}

func (service *ArticleService) getProfileById(ctx context.Context, id string) (*pb.ProfileResponse, error) {
	if ctx.Value(interceptor.TokenKey{}) != nil {
		md := metadata.New(map[string]string{"Authorization": "Token " + ctx.Value(interceptor.TokenKey{}).(string)})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	
	return service.profileServiceClient.GetProfileById(ctx, &pb.ProfileIdRequest{Id: id})
}

func (service *ArticleService) findArticlesByTag(tag string, offset, limit int32) ([]*model.Article, error) {
	return nil, nil
}

func (service *ArticleService) findArticlesByAuthor(tag string, offset, limit int32) ([]*model.Article, error) {
	return nil, nil
}

func (service *ArticleService) findArticlesByFavoritedUsername(tag string, offset, limit int32) ([]*model.Article, error) {
	return nil, nil
}

func (service *ArticleService) findArticles(offset, limit int32) ([]*model.Article, error) {
	return service.articleStore.Find(offset, limit)
}

func (service *ArticleService) isFavoritedByUser(userId string) bool {
	return false
}