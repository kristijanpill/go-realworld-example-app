package service

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	commentStore store.CommentStore
	articleStore store.ArticleStore
	profileServiceClient pb.ProfileServiceClient
}

func NewCommentService(commentStore store.CommentStore, articleStore store.ArticleStore, profileServiceClient pb.ProfileServiceClient) *CommentService {
	return &CommentService{
		commentStore: commentStore,
		articleStore: articleStore,
		profileServiceClient: profileServiceClient,
	}
}

func (service *CommentService) CreateArticleComment(ctx context.Context, request *pb.NewCommentRequest) (*pb.SingleCommentResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	article, err := service.articleStore.FindBySlug(request.Comment.Slug)
	if err != nil {
		return nil, err
	}

	author, err := service.getProfileById(ctx, currentUserIdString)
	if err != nil {
		return nil, err
	}

	comment, err := model.NewComment(currentUserIdString, article, request.Comment.Body)
	if err != nil {
		return nil, err
	}

	comment, err = service.commentStore.Create(comment)
	if err != nil {
		return nil, err
	}
	
	return &pb.SingleCommentResponse{
		Comment: &pb.Comment{
			Id: comment.ID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: comment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Body: comment.Body,
			Author: &pb.Profile{
				Username: author.Profile.Username,
				Bio: author.Profile.Bio,
				Image: author.Profile.Image,
				Following: author.Profile.Following,
			},
		},
	}, nil
}

func (service *CommentService) DeleteArticleComment(ctx context.Context, request *pb.DeleteArticleCommentRequest) (*emptypb.Empty, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)
	article, err := service.articleStore.FindBySlug(request.Slug)
	if err != nil {
		return nil, err
	}

	comment, err := service.commentStore.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	if comment.ArticleID.String() != article.ID.String() {
		return nil, status.Error(codes.InvalidArgument, "comment does not belong to this article")
	}

	if comment.UserID.String() != currentUserIdString {
		return nil, status.Error(codes.Unauthenticated, "forbidden")
	}

	err = service.commentStore.Delete(comment)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (service *CommentService) getProfileById(ctx context.Context, id string) (*pb.ProfileResponse, error) {
	if ctx.Value(interceptor.TokenKey{}) != nil {
		md := metadata.New(map[string]string{"Authorization": "Token " + ctx.Value(interceptor.TokenKey{}).(string)})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return service.profileServiceClient.GetProfileById(ctx, &pb.ProfileIdRequest{Id: id})
}