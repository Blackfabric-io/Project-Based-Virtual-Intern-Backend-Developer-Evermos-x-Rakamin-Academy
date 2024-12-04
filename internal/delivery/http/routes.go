package http

import (
	"evermos-project/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)

	api := app.Group("/api")
	api.Use(middleware.JWTProtected())
	api.Get("/profile", GetProfile)

	regionHandler := NewRegionHandler()
	region := app.Group("/api/regions")
	region.Get("/provinces", regionHandler.GetProvinces)
	region.Get("/provinces/:provinceID/cities", regionHandler.GetCities)
}
