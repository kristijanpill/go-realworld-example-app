package model

import (
	"github.com/google/uuid"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
)

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

func (profile *Profile) ProfileResponse(following bool) *pb.ProfileResponse {
	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: profile.Username,
			Bio: profile.Bio,
			Image: profile.Image,
			Following: following,
		},
	}
}