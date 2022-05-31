package model

import "github.com/google/uuid"

type Profile struct {
	ID       uuid.UUID `gorm:"primaryKey; unique; type:uuid"`
	Username string `gorm:"unique"`
	Bio      string
	Image    string
}

func NewProfile(id, username, bio, image string) (*Profile, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &Profile{
		ID: uuid,
		Username: username,
		Bio: bio,
		Image: image,
	}, nil
}