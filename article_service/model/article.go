package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type Article struct {
	UserID uuid.UUID `gorm:"type:uuid"`
	Slug   string    `gorm:"primaryKey"`
	Title string
	Description string
	Body string
	Tags []*Tag `gorm:"many2many:article_tags"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

func NewArticle(userIdString, title, description, body string, tags []*Tag) (*Article, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, err
	}

	return &Article{
		UserID: userId,
		Slug: slug.Make(title),
		Title: title,
		Description: description,
		Body: body,
		Tags: tags,
	}, nil
}