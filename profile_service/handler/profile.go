package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/service"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func(handler *ProfileHandler) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.ProfileInfo, error) {
	return handler.service.CreateProfile(request)
}

func (handler *ProfileHandler) GetProfileById(ctx context.Context, request *pb.ProfileIdRequest) (*pb.ProfileResponse, error) {
	return handler.service.FindById(request)
}