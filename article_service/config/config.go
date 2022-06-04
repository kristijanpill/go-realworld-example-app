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
		ProfileServiceHost:      os.Getenv("PROFILE_SERVICE_HOST"),
		PublicKey:               os.Getenv("PUBLIC_KEY"),
		ProfileServicePort:      os.Getenv("PROFILE_SERVICE_PORT"),
		ArticleDatabaseHost:     os.Getenv("PROFILE_DATABASE_HOST"),
		ArticleDatabasePort:     os.Getenv("PROFILE_DATABASE_PORT"),
		ArticleDatabaseUser:     os.Getenv("PROFILE_DATABASE_USER"),
		ArticleDatabasePassword: os.Getenv("PROFILE_DATABASE_PASSWORD"),
		ArticleDatabaseName:     os.Getenv("PROFILE_DATABASE_NAME"),
		RestrictedPaths:         restrictedPaths(),
	}
}

func restrictedPaths() map[string]bool {
	const articleService = "/article.ArticleService/"

	return map[string]bool{
		articleService + "GetArticles":           false,
		articleService + "CreateArticle":         true,
		articleService + "GetArticle":            false,
		articleService + "GetArticlesFeed":       true,
		articleService + "UpdateArticle":         true,
		articleService + "DeleteArticle":         true,
		articleService + "GetArticleComments":    false,
		articleService + "CreateArticleComment":  true,
		articleService + "DeleteArticleComment":  true,
		articleService + "CreateArticleFavorite": true,
		articleService + "DeleteArticleFavorite": true,
	}
}