package config

import "os"

type Config struct {
	Port               string
	UserServiceHost    string
	UserServicePort    string
	ProfileServiceHost string
	ProfileServicePort string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("API_GATEWAY_PORT"),
		UserServiceHost:    os.Getenv("USER_SERVICE_HOST"),
		UserServicePort:    os.Getenv("USER_SERVICE_PORT"),
		ProfileServiceHost: os.Getenv("PROFILE_SERVICE_HOST"),
		ProfileServicePort: os.Getenv("PROFILE_SERVICE_PORT"),
	}
}