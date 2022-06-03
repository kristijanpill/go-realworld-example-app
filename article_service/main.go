package main

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/config"
	"github.com/kristijanpill/go-realworld-example-app/article_service/server"
)

func main() {
	config := config.NewConfig()
	server := server.NewServer(config)
	server.Start()
}