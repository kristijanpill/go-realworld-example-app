package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type FavoriteStore interface {
	Create(*model.Favorite) (*model.Favorite, error)
	FindByUserIdAndSlug(userId, slug string) (*model.Favorite, error)
	Delete(*model.Favorite) error
}