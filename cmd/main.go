package main

import (
	"evermos-project/configs"
	"evermos-project/internal/delivery/http"
	"evermos-project/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitDB()
	app := fiber.NEW()
	http.InitRoutes(app)
	app.Listen(":3000")
	configs.Migrate()

	if err := utils.CreateUploadDir(); err != nil {
		log.Fatal("Failed to create upload directory", err)
	}
}
