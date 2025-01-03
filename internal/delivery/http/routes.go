package http

import (
	"evermos-project/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepo)

	productHandler := NewProductHandler(productUsecase)

	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)

	api := app.Group("/api")
	api.Use(middleware.JWTProtected())
	api.Get("/profile", GetProfile)
	products := api.Group("/products")
	products.Get("/", productHandler.GetProducts)

	regionHandler := NewRegionHandler()
	region := app.Group("/api/regions")
	region.Get("/provinces", regionHandler.GetProvinces)
	region.Get("/provinces/:provinceID/cities", regionHandler.GetCities)

	uploadHandler := NewUploadHandler()
	api := app.Group("/api")
	api.Use(middleware.JWTProtected())

	api.Post("/upload/product", uploadHandler.UploadProductImage)
	api.Post("/upload/avatar", uploadHandler.UploadAvatar)

	app.Static("/uploads", "./uploads")
}
