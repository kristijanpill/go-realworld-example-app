package config

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
		// Port:                    os.Getenv("API_GATEWAY_PORT"),
		// ProfileDatabaseHost:     os.Getenv("PROFILE_DATABASE_HOST"),
		// ProfileDatabasePort:     os.Getenv("PROFILE_DATABASE_PORT"),
		// ProfileDatabaseUser:     os.Getenv("PROFILE_DATABASE_USER"),
		// ProfileDatabasePassword: os.Getenv("PROFILE_DATABASE_PASSWORD"),
		// ProfileDatabaseName:     os.Getenv("PROFILE_DATABASE_NAME"),
		// PublicKey:               os.Getenv("PUBLIC_KEY"),
		Port:                    "8083",
		ProfileServiceHost:      "localhost",
		ProfileServicePort:      "8082",
		PublicKey:               "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0AzWYJTc9jiPn+RMNjMJ\nhscn8hg/Mt0U22efM6IvM83CyQCiFHP1Z8rs2HFqRbid/hQxW23HrXQzKx5hGPdU\n14ncF8oN7utDQxdq6ivTsF1tMQtHWb2jnYmpKwTyelbMMGKLHj3yy2j59Y/X94EX\nPNtQtgAO9FF5gKzjkaBu6KzLU2RJC9bADVd5sotM/JP/Ce5D/97XV7i1KStTUDiV\nfDBWCkDylBTQTmI1rO9MdayVduuAzNdWXRfyqKcWI2i4pA1aaskiaViVsIhF3ksm\nYW4Bu0RxK5SP2byHj7pv93XsabA+QXZ37QRhYzBxx6nS0x/dNtAxIltIBZaeSTN0\ngQIDAQAB\n-----END PUBLIC KEY-----",
		ArticleDatabaseHost:     "localhost",
		ArticleDatabasePort:     "5432",
		ArticleDatabaseUser:     "postgres",
		ArticleDatabasePassword: "root",
		ArticleDatabaseName:     "articles",
		RestrictedPaths:         restrictedPaths(),
	}
}

func restrictedPaths() map[string]bool {
	const articleService = "/article.ArticleService/"

	return map[string]bool{
		articleService + "GetArticles":           false,
		articleService + "CreateArticle":         true,
		articleService + "GetArticle":            false,
		articleService + "UpdateArticle":         true,
		articleService + "DeleteArticle":         true,
		articleService + "GetArticleComments":    false,
		articleService + "CreateArticleComment":  true,
		articleService + "DeleteArticleComment":  true,
		articleService + "CreateArticleFavorite": true,
		articleService + "DeleteArticleFavorite": true,
	}
}