package routes

import (
	"github.com/gofiber/fiber/v2"
)

func WelcomePage(c *fiber.Ctx) error {
	return c.SendString(`Hello Mohanraj`)
}
func NewPage(c *fiber.Ctx) error {
	return c.SendString("This is new Page for test")

}

func ConfigureRoutes(app *fiber.App) {
	app.Get("/", WelcomePage)
	app.Get("/new", NewPage)
}
