package server

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/article_service/config"
	"github.com/kristijanpill/go-realworld-example-app/article_service/handler"
	"github.com/kristijanpill/go-realworld-example-app/article_service/service"
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/db"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux: runtime.NewServeMux(),
	}

	return server;
}

func (server *Server) Start() {
	db := server.initDbConnection()
	articleStore := server.initArticlePostgresStore(db)
	tagStore := server.initTagPostgresStore(db)
	favoriteStore := server.initFavoritePostgresStore(db)
	commentStore := server.initCommentPostgresStore(db)
	profileServiceClient := server.initProfileServiceClient()
	articleService := service.NewArticleService(articleStore, tagStore, favoriteStore, profileServiceClient)
	tagService := service.NewTagService(tagStore)
	favoriteService := service.NewFavoriteService(favoriteStore, articleStore, profileServiceClient)
	commentService := service.NewCommentService(commentStore, articleStore, profileServiceClient)
	articleHandler := handler.NewArticleHandler(articleService, tagService, favoriteService, commentService)

	publicKey := server.initPublicKey()
	authInterceptor := interceptor.NewAuthInterceptor("Token", server.config.RestrictedPaths, publicKey)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()))
	pb.RegisterArticleServiceServer(grpcServer, articleHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func (server *Server) initDbConnection() *gorm.DB {
	db, err := db.NewPostgresConnection(server.config.ArticleDatabaseHost, server.config.ArticleDatabasePort, server.config.ArticleDatabaseName, server.config.ArticleDatabaseUser, server.config.ArticleDatabasePassword)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return db
}

func (server *Server) initArticlePostgresStore(db *gorm.DB) *store.ArticlePostgresStore {
	articleStore, err := store.NewArticlePostgresStore(db)
	if err != nil {
		log.Fatal("failed to init article store: ", err)
	}

	return articleStore
}

func (server *Server) initTagPostgresStore(db *gorm.DB) *store.TagPostgresStore {
	tagStore, err := store.NewTagPostgresStore(db)
	if err != nil {
		log.Fatal("failed to init tag store: ", err)
	}

	return tagStore
}

func (server *Server) initFavoritePostgresStore(db *gorm.DB) *store.FavoritePostgresStore {
	favoriteStore, err := store.NewFavoritePostgresStore(db)
	if err != nil {
		log.Fatal("failed to init favorite store: ", err)
	}

	return favoriteStore
}

func (server *Server) initCommentPostgresStore(db *gorm.DB) *store.CommentPostgresStore {
	commentStore, err := store.NewCommentPostgresStore(db)
	if err != nil {
		log.Fatal("failed to init comment store: ", err)
	}

	return commentStore
}

func (server *Server) initProfileServiceClient() pb.ProfileServiceClient {
	address := fmt.Sprintf("%s:%s", server.config.ProfileServiceHost, server.config.ProfileServicePort)
	profileServiceClient, err := service.NewProfileServiceClient(address)
	if err != nil {
		log.Fatal("cannot create profile service client: ", err)
	}

	return profileServiceClient
}

func (server *Server) initPublicKey() *rsa.PublicKey {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatal("cannot parse public key: ", err)
	}

	return publicKey
}