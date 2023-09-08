package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mohanrajreese.in/sample-gofiber/app/routes"
)

func Run() {
	// Creating New Fiber Instance
	app := fiber.New()

	// Creating a new middleware for logging
	app.Use(logger.New())

	// Allowing Cross Origin
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.ConfigureRoutes(app)
	port := fmt.Sprintf(":%s", "2105")
	err := app.Listen(port)

	if err != nil {
		return
	}
}
