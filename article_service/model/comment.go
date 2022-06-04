package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID int32 `gorm:"primaryKey;autoIncrement"`
	UserID uuid.UUID `gorm:"type:uuid"`
	ArticleID uuid.UUID `gorm:"type:uuid"`
	Article *Article 
	Body string
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

func NewComment(userIdString string, article *Article, body string) (*Comment, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, err
	}
	
	return &Comment{
		UserID: userId,
		ArticleID: article.ID,
		Article: article,
		Body: body,
	}, nil
}