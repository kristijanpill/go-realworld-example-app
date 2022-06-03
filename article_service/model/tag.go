package model

type Tag struct {
	ID      int32 `gorm:"primaryKey;autoIncrement"`
	Name    string
	Aricles []*Article `gorm:"many2many:article_tags"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}