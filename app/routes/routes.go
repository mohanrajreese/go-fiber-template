package routes

import (
	"github.com/gofiber/fiber/v2"
)

func WelcomePage(c *fiber.Ctx) error {
	return c.SendString(`Hello Mohanraj`)
}

func ConfigureRoutes(app *fiber.App) {
	app.Get("/", WelcomePage)

}
