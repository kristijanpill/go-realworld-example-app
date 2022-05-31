package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/common/db"
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
	profileHandler := handler.NewProfileHandler(profileService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
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
