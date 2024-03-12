package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	var err error

	// init config
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	PrintJSON(cfg)

	// init router
	router := fiber.New()

	// init minio
	minio, err := NewMinio(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init controller
	controller := NewController(minio)

	/* Setup HTTPServer*/
	router, err = setupHTTPServer(router)
	if err != nil {
		log.Fatal(err)
	}

	/* Setup Router */
	router, err = setupAPIRouter(router, controller)
	if err != nil {
		log.Fatal(err)
	}

	// server listen
	err = router.Listen(cfg.HTTP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(router *fiber.App) (*fiber.App, error) {

	router.Use(cors.New())
	router.Use(logger.New())
	router.Use(recover.New())

	router.Static("/", "./public")

	return router, nil
}

func setupAPIRouter(router *fiber.App, ctlr *Controller) (*fiber.App, error) {

	router.Get("/ping", ctlr.Ping)

	router.Post("/post", ctlr.Post)

	router.Post("/upload", ctlr.Upload)

	return router, nil
}
