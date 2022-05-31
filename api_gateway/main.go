package main

import (
	"github.com/kristijanpill/go-realworld-example-app/api_gateway/config"
	"github.com/kristijanpill/go-realworld-example-app/api_gateway/server"
)

func main() {
	config := config.NewConfig()
	server := server.NewServer(config)
	server.Start()
}