package http

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)
	//
}
