package main

import (
	"encoding/json"
	"fmt"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

func NewMinio(cfg *Config) (minioClient *minio.Client, err error) {
	minioClient, err = minio.NewWithOptions(cfg.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MINIO_ACCESS_KEY_ID, cfg.MINIO_SECRET_ACCESS_KEY, ""),
		Secure: cfg.MINIO_USE_SSL,
	})

	if err != nil {
		return
	}

	fmt.Printf("\nconnect minio success\n\n")

	return
}

func PrintJSON(a any) {
	json, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println("json marshal error")
	}

	fmt.Println(string(json))
}
