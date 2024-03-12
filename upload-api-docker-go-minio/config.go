package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT string

	MINIO_ENDPOINT          string
	MINIO_ACCESS_KEY_ID     string
	MINIO_SECRET_ACCESS_KEY string
	MINIO_USE_SSL           bool
}

func NewConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// 寫死 MINIO_USE_SSL 簡化程式碼
	return &Config{
		HTTP_PORT:               os.Getenv("HTTP_PORT"),
		MINIO_ENDPOINT:          os.Getenv("MINIO_ENDPOINT"),
		MINIO_ACCESS_KEY_ID:     os.Getenv("MINIO_ACCESS_KEY_ID"),
		MINIO_SECRET_ACCESS_KEY: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		MINIO_USE_SSL:           false,
	}, nil
}
