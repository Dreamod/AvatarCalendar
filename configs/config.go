package configs

import (
	"fmt"
	"os"
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

type EnvConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
	AssetsDir  string
}

func loadEnv() *EnvConfig {
	return &EnvConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "avatarCalendar"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "avatarCalendar"),
		AppPort:    getEnv("APP_PORT", "8080"),
		AssetsDir:  getEnv("ASSETS_DIR", "./assets"),
	}
}

func LoadConfig() *Config {
	env := loadEnv()
	return &Config{
		Db: DbConfig{
			Dsn: makeDSN(env),
		},
		Auth: AuthConfig{
			Secret: getEnv("TOKEN", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func makeDSN(env *EnvConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
}
