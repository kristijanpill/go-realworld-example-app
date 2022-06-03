package store

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"gorm.io/gorm"
)

type FavoritePostgresStore struct {
	db *gorm.DB
}

func NewFavoritePostgresStore(db *gorm.DB) (*FavoritePostgresStore, error) {
	err := db.AutoMigrate(&model.Favorite{})	
	if err != nil {
		return nil, err
	}

	return &FavoritePostgresStore{
		db: db,
	}, nil
}

func (store *FavoritePostgresStore) Create(favorite *model.Favorite) (*model.Favorite, error) {
	result := store.db.Create(favorite)

	return favorite, result.Error
}