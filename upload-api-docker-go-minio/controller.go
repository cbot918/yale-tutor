package main

import (
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go"
)

type Controller struct {
	Minio *minio.Client
}

func NewController(minio *minio.Client) *Controller {
	return &Controller{
		Minio: minio,
	}
}

func (ctlr *Controller) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func (ctlr *Controller) Post(c *fiber.Ctx) error {
	fmt.Println("in post controller")
	return c.SendString("post")
}

func (ctlr *Controller) Upload(c *fiber.Ctx) error {

	fmt.Println("in upload")

	// Retrieve file from posted form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Upload request is missing a file.")
	}

	// Open file for reading
	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read the file.")
	}
	defer fileData.Close()

	// Read file content to buffer
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(fileData); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read the file data.")
	}

	// Upload file to MinIO
	bucketName := "testbucket" // Specify your bucket name here
	objectName := file.Filename
	contentType := "application/octet-stream"

	// Upload the file
	info, err := ctlr.Minio.PutObject(bucketName, objectName, buf, int64(buf.Len()), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to upload: %v", err))
	}

	fmt.Println("Successfully uploaded: ", info)

	return c.SendString(fmt.Sprintln("Successfully uploaded: ", info))

}
