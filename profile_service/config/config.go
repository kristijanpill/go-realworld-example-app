package config

import "os"

type Config struct {
	Port                    string
	ProfileDatabaseHost     string
	ProfileDatabasePort     string
	ProfileDatabaseUser     string
	ProfileDatabasePassword string
	ProfileDatabaseName     string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("API_GATEWAY_PORT"),
		ProfileDatabaseHost: os.Getenv("PROFILE_DATABASE_HOST"),
		ProfileDatabasePort: os.Getenv("PROFILE_DATABASE_PORT"),
		ProfileDatabaseUser: os.Getenv("PROFILE_DATABASE_USER"),
		ProfileDatabasePassword: os.Getenv("PROFILE_DATABASE_PASSWORD"),
		ProfileDatabaseName: os.Getenv("PROFILE_DATABASE_NAME"),
	}
}