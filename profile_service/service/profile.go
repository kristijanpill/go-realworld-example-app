package service

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
)

type ProfileService struct {
	store store.ProfileStore
	followService *FollowService
}

func NewProfileService(store store.ProfileStore, followService *FollowService) *ProfileService {
	return &ProfileService{
		store: store,
		followService: followService,
	}
}

func (service *ProfileService) GetProfileByUsername(ctx context.Context, request *pb.ProfileUsernameRequest) (*pb.ProfileResponse, error) {
	profile, err := service.store.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	following := false
	if (ctx.Value(interceptor.CurrentUserKey{}) != nil) {
		following = service.followService.ExistsByProfileIdAndTargetId(ctx.Value(interceptor.CurrentUserKey{}).(string), profile.ID.String())
	}

	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: profile.Username,
			Bio: profile.Bio,
			Image: profile.Image,
			Following: following,
		},
	}, nil
}

func (service *ProfileService) GetProfileById(ctx context.Context, request *pb.ProfileIdRequest) (*pb.ProfileResponse, error) {
	profile, err := service.store.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	following := false
	if (ctx.Value(interceptor.CurrentUserKey{}) != nil) {
		following = service.followService.ExistsByProfileIdAndTargetId(ctx.Value(interceptor.CurrentUserKey{}).(string), profile.ID.String())
	}

	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: profile.Username,
			Bio: profile.Bio,
			Image: profile.Image,
			Following: following,
		},
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