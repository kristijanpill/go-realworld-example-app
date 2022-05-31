package config

import "os"

type Config struct {
	Port                 string
	ProfileServiceHost   string
	ProfileServicePort   string
	UserDatabaseHost     string
	UserDatabasePort     string
	UserDatabaseUser     string
	UserDatabasePassword string
	UserDatabaseName     string
	PrivateKey           string
	PublicKey            string
	RestrictedPaths      map[string]bool
}

func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("API_GATEWAY_PORT"),
		ProfileServiceHost:   os.Getenv("PROFILE_SERVICE_HOST"),
		ProfileServicePort:   os.Getenv("PROFILE_SERVICE_PORT"),
		PrivateKey:           os.Getenv("PRIVATE_KEY"),
		PublicKey:            os.Getenv("PUBLIC_KEY"),
		UserDatabaseHost:     os.Getenv("USER_DATABASE_HOST"),
		UserDatabasePort:     os.Getenv("USER_DATABASE_PORT"),
		UserDatabaseUser:     os.Getenv("USER_DATABASE_USER"),
		UserDatabasePassword: os.Getenv("USER_DATABASE_PASSWORD"),
		UserDatabaseName:     os.Getenv("USER_DATABASE_Name"),
		RestrictedPaths:      restrictedPaths(),
	}
}

func restrictedPaths() map[string]bool {
	const userService = "/user.UserService/"

	return map[string]bool{
		userService + "GetCurrentUser":    true,
		userService + "UpdateCurrentUser": true,
	}
}