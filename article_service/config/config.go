package config

import "os"

type Config struct {
	Port                    string
	ProfileServiceHost      string
	ProfileServicePort      string
	PublicKey               string
	ArticleDatabaseHost     string
	ArticleDatabasePort     string
	ArticleDatabaseUser     string
	ArticleDatabasePassword string
	ArticleDatabaseName     string
	RestrictedPaths         map[string]bool
}

func NewConfig() *Config {
	return &Config{
		Port:                    os.Getenv("API_GATEWAY_PORT"),
		ProfileServiceHost: os.Getenv("PROFILE_SERVICE_HOST"),
		ProfileServicePort: os.Getenv("PROFILE_SERVICE_PORT"),
		PublicKey: os.Getenv("PUBLIC_KEY"),
		ArticleDatabaseHost:     os.Getenv("ARTICLE_DATABASE_HOST"),
		ArticleDatabasePort:     os.Getenv("ARTICLE_DATABASE_PORT"),
		ArticleDatabaseUser:     os.Getenv("ARTICLE_DATABASE_USER"),
		ArticleDatabasePassword: os.Getenv("ARTICLE_DATABASE_PASSWORD"),
		ArticleDatabaseName:     os.Getenv("ARTICLE_DATABASE_NAME"),
		RestrictedPaths:         restrictedPaths(),
	}
}

func restrictedPaths() map[string]bool {
	const articleService = "/article.ArticleService/"

	return map[string]bool{
		articleService + "CreateArticle": true,
	}
}