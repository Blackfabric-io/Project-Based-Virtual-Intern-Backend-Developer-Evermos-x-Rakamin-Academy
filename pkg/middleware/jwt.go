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

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest.JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
	}
}

