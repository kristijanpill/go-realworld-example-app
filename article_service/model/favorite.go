package model

import "github.com/google/uuid"

type Favorite struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	ArticleID string `gorm:"primaryKey"`
	Article *Article
}

func NewFavorite(userIdString string, article *Article) (*Favorite, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, err
	}

	return &Favorite{
		UserID: userId,
		ArticleID: article.Slug,
		Article: article,
	}, nil
}