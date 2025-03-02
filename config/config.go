package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	JWTSecret     string
	StorageRegion string
	StorageKey    string
	StorageSecret string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		ServerPort:    getEnv("SERVER_PORT", ":8080"),
		JWTSecret:     getEnv("JWT_SECRET", "Bk1Rqg1Vl2oktA1pTpIbzbZAeWIbus"),
		StorageRegion: getEnv("STORAGE_REGION", "ap-northeast-2"),
		StorageKey:    getEnv("STORAGE_KEY", ""),
		StorageSecret: getEnv("STORAGE_SECRET", ""),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://localhost:5234"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
