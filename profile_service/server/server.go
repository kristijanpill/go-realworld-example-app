package server

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/common/db"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/config"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/handler"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/service"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
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
	profileStore := server.initProfilePostgresStore(db)
	profileService := service.NewProfileService(profileStore)
	followStore := server.initFollowPostgresStore(db)
	followService := service.NewFollowService(followStore, profileStore)
	profileHandler := handler.NewProfileHandler(profileService, followService)

	publicKey := server.initPublicKey()
	authInterceptor := interceptor.NewAuthInterceptor("Token", server.config.RestrictedPaths, publicKey)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()))
	pb.RegisterProfileServiceServer(grpcServer, profileHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func (server *Server) initDbConnection() *gorm.DB {
	db, err := db.NewPostgresConnection(server.config.ProfileDatabaseHost, server.config.ProfileDatabasePort, server.config.ProfileDatabaseName, server.config.ProfileDatabaseUser, server.config.ProfileDatabasePassword)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return db
}

func (server *Server) initProfilePostgresStore(db *gorm.DB) *store.ProfilePostgresStore {
	profileStore, err := store.NewProfilePostgresStore(db)
	if err != nil {
		log.Fatal("failed to init profile store: ", err)
	}

	return profileStore
}

func (server *Server) initFollowPostgresStore(db *gorm.DB) *store.FollowPostgresStore {
	followStore, err := store.NewFollowPostgresStore(db)
	if err != nil {
		log.Fatal("failed to init follow store: ", err)
	}

	return followStore
}

func (server *Server) initPublicKey() *rsa.PublicKey {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatal("cannot parse public key: ", err)
	}

	return publicKey
}