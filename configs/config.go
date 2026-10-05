package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DB_DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
