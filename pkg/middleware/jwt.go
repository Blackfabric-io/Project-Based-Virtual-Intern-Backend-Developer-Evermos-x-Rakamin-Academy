package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

func AuthRequired() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey: []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError