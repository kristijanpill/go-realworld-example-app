package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/service"
	"google.golang.org/protobuf/types/known/emptypb"
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
	return handler.profileService.GetProfileByUsername(ctx, request)
}

func (handler *ProfileHandler) FollowUserByUsername(ctx context.Context, request *pb.FollowRequest) (*pb.ProfileResponse, error) {
	return handler.followService.FollowUserByUsername(ctx, request)
}

func (handler *ProfileHandler) UnfollowUserByUsername(ctx context.Context, request *pb.UnfollowRequest) (*pb.ProfileResponse, error) {
	return handler.followService.UnfollowUserByUsername(ctx, request)
}

func (handler *ProfileHandler) GetProfileById(ctx context.Context, request *pb.ProfileIdRequest) (*pb.ProfileResponse, error) {
	return handler.profileService.GetProfileById(ctx, request)
}

func(handler *ProfileHandler) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.CreateProfile(request)
}

func (handler *ProfileHandler) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.profileService.UpdateProfile(request)
}

func (handler *ProfileHandler) GetProfileIdByUsername(ctx context.Context, request *pb.ProfileIdUsernameRequest) (*pb.ProfileIdResponse, error) {
	return handler.profileService.GetProfileIdByUsername(request)
}

func (handler *ProfileHandler) GetFollowedProfileIds(ctx context.Context, request *emptypb.Empty) (*pb.FollowedIds, error) {
	return handler.followService.GetFollowedProfileIds(ctx)
}
