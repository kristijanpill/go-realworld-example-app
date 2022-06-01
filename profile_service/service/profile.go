package service

import (
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

func (service *ProfileService) GetProfileByUsername(request *pb.ProfileUsernameRequest) (*model.Profile, error) {
	profile, err := service.store.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (service *ProfileService) GetProfileById(request *pb.ProfileIdRequest) (*pb.ProfileInfo, error) {
	profile, err := service.store.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileInfo{
		Username: profile.Username,
		Bio: profile.Bio,
		Image: profile.Image,
	}, nil
}

func (service *ProfileService) CreateProfile(request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
	profile, err := model.NewProfile(request.Id, request.Profile.Username, request.Profile.Bio, request.Profile.Image)
	if err != nil {
		return nil, err
	}

	_, err = service.store.Create(profile)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileInfo{
		Username: profile.Username,
		Bio: profile.Bio,
		Image: profile.Image,
	}, nil
}

func (service *ProfileService) UpdateProfile(request *pb.UpdateProfileRequest) (*pb.ProfileInfo, error) {
	profile, err := service.store.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	profile.Username = request.Profile.Username
	profile.Bio = request.Profile.Bio
	profile.Image = request.Profile.Image

	_, err = service.store.Update(profile)
	if err != nil {
		return nil, err
	}

	return request.Profile, nil
}