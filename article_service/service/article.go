package service

import (
	"context"
	"errors"

	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ArticleService struct {
	articleStore store.ArticleStore
	tagStore store.TagStore
	favoriteStore store.FavoriteStore
	profileServiceClient pb.ProfileServiceClient
}

func NewArticleService(articleStore store.ArticleStore, tagStore store.TagStore, favoriteStore store.FavoriteStore, profileServiceClient pb.ProfileServiceClient) *ArticleService {
	return &ArticleService{
		articleStore: articleStore,
		tagStore: tagStore,
		favoriteStore: favoriteStore,
		profileServiceClient: profileServiceClient,
	}
}

func (service *ArticleService) GetArticlesFeed(ctx context.Context, request *pb.GetArticlesFeedRequest) (*pb.MultipleArticlesResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	limit := request.Limit
	if limit <= 0 {
		limit = 20
	}
	offset := request.Offset
	if offset < 0 {
		offset = 0
	}

	followedIds, err := service.getFollowedProfileIds(ctx, currentUserIdString)
	if err != nil {
		return nil, err
	}

	articles, err := service.articleStore.FindByUserIds(offset, limit, followedIds.Ids)
	if err != nil {
		return nil, err
	}

	response := &pb.MultipleArticlesResponse{
		Articles: []*pb.Article{},
		ArticlesCount: int32(len(articles)),
	}

	for _, articleModel := range articles {
		author, err := service.getProfileById(ctx, articleModel.UserID.String())
		if err != nil {
			return nil, err
		}

		isFavorited := false
		if ctx.Value(interceptor.CurrentUserKey{}) != nil {
			isFavorited = service.isFavoritedByUserId(ctx.Value(interceptor.CurrentUserKey{}).(string), articleModel.ID.String())
		}

		favoritesCount, err := service.favoriteStore.CountFavoritesByArticleId(articleModel.ID.String())
		if err != nil {
			return nil, err
		}

		article := &pb.Article{
			Slug: articleModel.Slug,
			Title: articleModel.Title,
			Description: articleModel.Description,
			Body: articleModel.Body,
			TagList: service.getTagList(articleModel),
			CreatedAt: articleModel.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: articleModel.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Favorited: isFavorited,
			FavoritesCount: int32(favoritesCount),
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
		articles, err = service.findArticlesByTag(offset, limit, request.Tag)
	} else if request.Author != "" {
		if userId, errr := service.getProfileIdByUsername(request.Author); errr == nil {
			articles, err = service.findArticlesByAuthor(offset, limit, userId.Id)
		}
	} else if request.Favorited != "" {
		if userId, errr := service.getProfileIdByUsername(request.Favorited); errr == nil {
			articles, err = service.findArticlesFavoritedByUserId(offset, limit, userId.Id)
		}
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
		author, err := service.getProfileById(ctx, articleModel.UserID.String())
		if err != nil {
			return nil, err
		}

		isFavorited := false
		if ctx.Value(interceptor.CurrentUserKey{}) != nil {
			isFavorited = service.isFavoritedByUserId(ctx.Value(interceptor.CurrentUserKey{}).(string), articleModel.ID.String())
		}

		favoritesCount, err := service.favoriteStore.CountFavoritesByArticleId(articleModel.ID.String())
		if err != nil {
			return nil, err
		}

		article := &pb.Article{
			Slug: articleModel.Slug,
			Title: articleModel.Title,
			Description: articleModel.Description,
			Body: articleModel.Body,
			TagList: service.getTagList(articleModel),
			CreatedAt: articleModel.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: articleModel.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Favorited: isFavorited,
			FavoritesCount: int32(favoritesCount),
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
			CreatedAt: article.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02T15:04:05Z"),
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

func (service *ArticleService) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.SingleArticleResponse, error) {
	article, err := service.findArticleBySlug(request.Slug)
	if err != nil {
		return nil, err
	}

	author, err := service.getProfileById(ctx, article.UserID.String())
	if err != nil {
		return nil, err
	}

	isFavorited := false
	if ctx.Value(interceptor.CurrentUserKey{}) != nil {
		isFavorited = service.isFavoritedByUserId(ctx.Value(interceptor.CurrentUserKey{}).(string), article.ID.String())
	}

	favoritesCount, err := service.favoriteStore.CountFavoritesByArticleId(article.ID.String())
	if err != nil {
		return nil, err
	}

	return &pb.SingleArticleResponse{
		Article: &pb.Article{
			Slug: article.Slug,
			Title: article.Title,
			Description: article.Description,
			Body: article.Body,
			TagList: service.getTagList(article),
			CreatedAt: article.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Favorited: isFavorited,
			FavoritesCount: int32(favoritesCount),
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		},
	}, nil
}

func (service *ArticleService) UpdateArticle(ctx context.Context, request *pb.UpdateArticleRequest) (*pb.SingleArticleResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	article, err := service.findArticleBySlug(request.Article.Slug)
	if err != nil {
		return nil, err
	}

	if article.UserID.String() != currentUserIdString {
		return nil, status.Error(codes.Unauthenticated, "forbidden")
	}

	article.Title = request.Article.Title
	article.Description = request.Article.Description
	article.Body = request.Article.Body

	article, err = service.updateArticle(article)
	if err != nil {
		return nil, err
	}

	author, err := service.getProfileById(ctx, article.UserID.String())
	if err != nil {
		return nil, err
	}

	isFavorited := false
	if ctx.Value(interceptor.CurrentUserKey{}) != nil {
		isFavorited = service.isFavoritedByUserId(ctx.Value(interceptor.CurrentUserKey{}).(string), article.ID.String())
	}

	favoritesCount, err := service.favoriteStore.CountFavoritesByArticleId(article.ID.String())
	if err != nil {
		return nil, err
	}

	return &pb.SingleArticleResponse{
		Article: &pb.Article{
			Slug: article.Slug,
			Title: article.Title,
			Description: article.Description,
			Body: article.Body,
			TagList: service.getTagList(article),
			CreatedAt: article.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Favorited: isFavorited,
			FavoritesCount: int32(favoritesCount),
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		},
	}, nil
}

func (service *ArticleService) DeleteArticle(ctx context.Context, request *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	article, err := service.findArticleBySlug(request.Slug)
	if err != nil {
		return nil, err
	}
	if article.UserID.String() != currentUserIdString {
		return nil, status.Error(codes.Unauthenticated, "forbidden")
	}

	err = service.deleteArticle(article)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (service *ArticleService) getFollowedProfileIds(ctx context.Context, id string) (*pb.FollowedIds, error) {
	if ctx.Value(interceptor.TokenKey{}) != nil {
		md := metadata.New(map[string]string{"Authorization": "Token " + ctx.Value(interceptor.TokenKey{}).(string)})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return service.profileServiceClient.GetFollowedProfileIds(ctx, &emptypb.Empty{})
}

func (service *ArticleService) getProfileById(ctx context.Context, id string) (*pb.ProfileResponse, error) {
	if ctx.Value(interceptor.TokenKey{}) != nil {
		md := metadata.New(map[string]string{"Authorization": "Token " + ctx.Value(interceptor.TokenKey{}).(string)})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return service.profileServiceClient.GetProfileById(ctx, &pb.ProfileIdRequest{Id: id})
}

func (service *ArticleService) getProfileIdByUsername(username string) (*pb.ProfileIdResponse, error) {
	return service.profileServiceClient.GetProfileIdByUsername(context.Background(), &pb.ProfileIdUsernameRequest{
		Username: username,
	})
}

func (service *ArticleService) findArticlesByTag(offset, limit int32, tag string) ([]*model.Article, error) {
	return service.articleStore.FindByTag(offset, limit, tag)
}

func (service *ArticleService) findArticlesByAuthor(offset, limit int32, userId string) ([]*model.Article, error) {
	return service.articleStore.FindByAuthorId(offset, limit, userId)
}

func (service *ArticleService) findArticlesFavoritedByUserId(offset, limit int32, userId string) ([]*model.Article, error) {
	return service.articleStore.FindFavoritedByUserId(offset, limit, userId)
}

func (service *ArticleService) findArticles(offset, limit int32) ([]*model.Article, error) {
	return service.articleStore.Find(offset, limit)
}

func (service *ArticleService) findArticleBySlug(slug string) (*model.Article, error) {
	return service.articleStore.FindBySlug(slug)
}

func (service *ArticleService) isFavoritedByUserId(slug, userId string) bool {
	return service.favoriteStore.IsArticleFavoritedByUserId(slug, userId)
}

func (service *ArticleService) updateArticle(article *model.Article) (*model.Article, error) {
	return service.articleStore.Update(article)
}

func (service *ArticleService) deleteArticle(article *model.Article) error {
	return service.articleStore.Delete(article)
}

func (service *ArticleService) getTagList(article *model.Article) []string {
	var tagList []string
	for _, tag := range article.Tags {
		tagList = append(tagList, tag.Name)
	}

	return tagList
}