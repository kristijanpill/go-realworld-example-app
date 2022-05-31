package service

import (
	"github.com/google/uuid"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
)

type ProfileService struct {
	store store.ProfileStore
}

func NewProfileService(store store.ProfileStore) *ProfileService {
	return &ProfileService{
		store: store,
	}
}

func (service *ProfileService) CreateProfile(request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
 	profile, err := model.NewProfile(request.Id, request.Profile.Username, request.Profile.Bio, request.Profile.Image)
	if err != nil {
		return nil, err
	}

	_, err = service.store.Save(profile)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileInfo{
		Username: profile.Username,
		Bio: profile.Bio,
		Image: profile.Image,
	}, nil
}

func (service *ProfileService) FindById(request *pb.ProfileIdRequest) (*pb.ProfileResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	profile, err := service.store.FindById(id)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: profile.Username,
			Bio: profile.Bio,
			Image: profile.Image,
			Following: false,
		},
	}, nil
}