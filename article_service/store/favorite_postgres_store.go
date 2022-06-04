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

func (store *FavoritePostgresStore) FindByUserIdAndArticleId(userId, articleId string) (*model.Favorite, error) {
	var favorite model.Favorite
	result := store.db.Preload("Article.Tags").Where("user_id = ? AND article_id = ?", userId, articleId).First(&favorite)

	return &favorite, result.Error
}

func (store *FavoritePostgresStore) Delete(favorite *model.Favorite) error {
	result := store.db.Delete(favorite)

	return result.Error
}

func (store *FavoritePostgresStore) IsArticleFavoritedByUserId(slug, userId string) bool {
	var favorite model.Favorite
	
	return store.db.Where("user_id = ? AND article_id = ?", slug, userId).First(&favorite).RowsAffected == 1
}

func (store *FavoritePostgresStore) CountFavoritesByArticleId(articleId string) (int64, error) {
	var count int64
	result := store.db.Model(&model.Favorite{}).Where("article_id = ?", articleId).Count(&count)

	return count, result.Error
}