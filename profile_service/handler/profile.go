package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/service"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	profileService *service.ProfileService
	followService *service.FollowService
}

func NewProfileHandler(profileService *service.ProfileService, followService *service.FollowService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		followService: followService,
	}
}

func(handler *ProfileHandler) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.CreateProfile(request)
}

func (handler *ProfileHandler) GetProfileById(ctx context.Context, request *pb.ProfileIdRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.FindById(request)
}

func (handler *ProfileHandler) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.UpdateProfile(request)
}

func (handler *ProfileHandler) FollowUserByUsername(ctx context.Context, request *pb.FollowRequest) (*pb.ProfileResponse, error) {
	return handler.followService.FollowUserByUsername(ctx, request)
}