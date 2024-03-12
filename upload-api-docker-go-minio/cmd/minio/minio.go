package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

const (
	endpoint        = "localhost:9000"
	accessKeyID     = "9OgSm2fFiPi4Q5JsWySs"
	secretAccessKey = "gg7mYVNM4bqN7WDj98vg3xtPqpLkuByjCjk7qjls"
	useSSL          = false
)

func main() {
	// Initialize minio client object.
	minioClient, err := minio.NewWithOptions(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "testbucket"
	objectName := "test.png"
	filePath := "cmd/minio/test.PNG" // Change this to the path of your test.png file
	contentType := "image/png"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(info)

}
