package store

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"gorm.io/gorm"
)

type TagPostgresStore struct {
	db *gorm.DB
}


func NewTagPostgresStore(db *gorm.DB) (*TagPostgresStore, error) {
	err := db.AutoMigrate(&model.Tag{})	
	if err != nil {
		return nil, err
	}

	return &TagPostgresStore{
		db: db,
	}, nil
}

func (store *TagPostgresStore) Create(tag *model.Tag) (*model.Tag, error) {
	result := store.db.Create(tag)

	return tag, result.Error
}

func (store *TagPostgresStore) FindByName(name string) (*model.Tag, error) {
	var tag model.Tag
	result := store.db.Where("name = ?", name).First(&tag)

	return &tag, result.Error
}

func (store *TagPostgresStore) FindAll() ([]*model.Tag, error) {
	var tags []*model.Tag
	result := store.db.Find(&tags)

	return tags, result.Error
}