package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GINMode   string `mapstructure:"GIN_MODE"`
	CookieKey string `mapstructure:"COOKIE_KEY"`
	CsrfKey   string `mapstructure:"CSRF_KEY"`
	Secure    string `mapstructure:"SECURE"`
	Domain    string `mapstructure:"DOMAIN"`
	URLScheme string `mapstructure:"URLSCHEME"`
}

func LoadConfig() (config Config, err error) {
	godotenv.Load("app.env")

	config.GINMode = os.Getenv("GIN_MODE")
	config.CookieKey = os.Getenv("COOKIE_KEY")
	config.CsrfKey = os.Getenv("CSRF_KEY")
	config.Secure = os.Getenv("SECURE")
	config.Domain = os.Getenv("DOMAIN")
	config.URLScheme = os.Getenv("URLSCHEME")

	return config, err
}
