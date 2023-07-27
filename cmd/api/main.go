package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/smochii/go-clean-architecture/config"
	"github.com/smochii/go-clean-architecture/custom_error"
	"github.com/smochii/go-clean-architecture/di"
)

func main() {
	// timezone setting
	time.Local = time.UTC

	// start server
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return custom_error.NewInternalError(*c)
			},
		},
	)
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Conf.App.AllowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	di.Route().Register(app)

	log.Fatal(app.Listen(":3000"))
}
