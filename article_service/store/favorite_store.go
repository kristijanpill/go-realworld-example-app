package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type FavoriteStore interface {
	Create(*model.Favorite) (*model.Favorite, error)
	FindByUserIdAndArticleId(userId, articleId string) (*model.Favorite, error)
	Delete(*model.Favorite) error
	IsArticleFavoritedByUserId(slug, userId string) bool
}