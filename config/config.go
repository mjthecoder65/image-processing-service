package config

import (
	"fmt"
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
	BucketName    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		ServerPort:    getEnvWithDefaultValue("SERVER_PORT", ":8080"),
		JWTSecret:     getEnvWithDefaultValue("JWT_SECRET", "Bk1Rqg1Vl2oktA1pTpIbzbZAeWIbus"),
		StorageRegion: getEnvWithDefaultValue("STORAGE_REGION", "ap-northeast-2"),
		StorageKey:    getEnvWithDefaultValue("STORAGE_KEY", ""),
		StorageSecret: getEnvWithDefaultValue("STORAGE_SECRET", ""),
		DatabaseURL:   getEnvWithDefaultValue("DATABASE_URL", "postgres://localhost:5234"),
		BucketName:    getEnvRequired("BUCKET_NAME"),
	}, nil
}

func getEnvWithDefaultValue(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	if value, exists := os.LookupEnv(key); !exists {
		message := fmt.Sprintf("failed to load value for key: %v", key)
		panic(message)
	} else {
		return value
	}
}
