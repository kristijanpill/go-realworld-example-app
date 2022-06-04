package store

import (
	"github.com/gosimple/slug"
	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"gorm.io/gorm"
)

type ArticlePostgresStore struct {
	db *gorm.DB
}

func NewArticlePostgresStore(db *gorm.DB) (*ArticlePostgresStore, error) {
	err := db.AutoMigrate(&model.Article{})	
	if err != nil {
		return nil, err
	}

	return &ArticlePostgresStore{
		db: db,
	}, nil
}

func (store *ArticlePostgresStore) Create(article *model.Article) (*model.Article, error) {
	result := store.db.Create(article)

	return article, result.Error
}

func (store *ArticlePostgresStore) Find(offset, limit int32) ([]*model.Article, error) {
	var articles []*model.Article
	result := store.db.Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Preload("Tags").Find(&articles)

	return articles, result.Error
}

func (store *ArticlePostgresStore) FindBySlug(slug string) (*model.Article, error) {
	var article model.Article
	result := store.db.Where("slug = ?", slug).Preload("Tags").First(&article)

	return &article, result.Error
}

func (store *ArticlePostgresStore) FindByTag(offset, limit int32, tag string) ([]*model.Article, error) {
	var tagModel *model.Tag
	store.db.Where("name = ?", tag).First(&tagModel)
	var articles []*model.Article
	var err error
	if tagModel != nil {
		err = store.db.Model(tagModel).Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Preload("Tags").Association("Articles").Find(&articles)
	}

	return articles, err
}

func (store *ArticlePostgresStore) FindByAuthorId(offset, limit int32, userId string) ([]*model.Article, error) {
	var articles []*model.Article
	result := store.db.Where("user_id = ?", userId).Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Preload("Tags").Find(&articles)

	return articles, result.Error
}

func (store *ArticlePostgresStore) FindFavoritedByUserId(offset, limit int32, userId string) ([]*model.Article, error) {
	var articles []*model.Article
	result := store.db.Joins("JOIN favorites ON favorites.article_id = articles.id").Where("favorites.user_id = ?", userId).Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Find(&articles)

	return articles, result.Error
}

func (store *ArticlePostgresStore) FindByUserIds(offset, limit int32, userIds []string) ([]*model.Article, error) {
	var articles []*model.Article
	result := store.db.Where("user_id IN ?", userIds).Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Preload("Tags").Find(&articles)

	return articles, result.Error
}

func (store *ArticlePostgresStore) Update(article *model.Article) (*model.Article, error) {
	result := store.db.Model(article).Updates(model.Article{Slug: slug.Make(article.Title), Title: article.Title, Description: article.Description, Body: article.Body})
	
	return article, result.Error
}

func (store *ArticlePostgresStore) Delete(article *model.Article) error {
	result := store.db.Select("Tags").Delete(article)

	return result.Error
}