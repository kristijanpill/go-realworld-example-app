package model

type Tag struct {
	ID       int32 `gorm:"primaryKey;autoIncrement"`
	Name     string
	Articles []*Article `gorm:"many2many:article_tags"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}