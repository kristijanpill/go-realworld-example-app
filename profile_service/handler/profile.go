package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
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

func (handler *ProfileHandler) GetProfileByUsername(ctx context.Context, request *pb.ProfileUsernameRequest) (*pb.ProfileResponse, error) {
	profile, err := handler.profileService.GetProfileByUsername(request)
	if err != nil {
		return nil, err
	}

	following := false
	if(ctx.Value(interceptor.CurrentUserKey{}) != nil) {
		following = handler.followService.ExistsByProfileIdAndTargetId(ctx.Value(interceptor.CurrentUserKey{}).(string), profile.ID.String())
	}

	return profile.ProfileResponse(following), nil
}

func (handler *ProfileHandler) FollowUserByUsername(ctx context.Context, request *pb.FollowRequest) (*pb.ProfileResponse, error) {
	return handler.followService.FollowUserByUsername(ctx, request)
}

func (handler *ProfileHandler) UnfollowUserByUsername(ctx context.Context, request *pb.UnfollowRequest) (*pb.ProfileResponse, error) {
	return handler.followService.UnfollowUserByUsername(ctx, request)
}

func (handler *ProfileHandler) GetProfileById(ctx context.Context, request *pb.ProfileIdRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.GetProfileById(request)
}

func(handler *ProfileHandler) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.CreateProfile(request)
}

func (handler *ProfileHandler) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.UpdateProfile(request)
}
