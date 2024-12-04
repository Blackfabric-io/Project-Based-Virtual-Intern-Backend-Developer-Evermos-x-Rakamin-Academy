package http

import (
	"evermos-project/configs"
	"evermos-project/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.Password = string(hashedPassword)

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	store := domain.Store{
		UserID: user.ID,
		Name: user.Name + "'s Store",
	}
	if err := configs.DB.Create(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create store",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Registration successful",
		"data":    fiber.Map{
			"user":  user,
			"store": store,
		}
	})
}

func Login(c *fiber.Ctx) error {
	input := new(domain.User)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := new(domain.User)
	configs.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	token, err := generateJwt(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var userData domain.User
	if err := configs.DB.First(&userData, userID).Error; err != nil (
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	)

	userData.Password = ""

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": userData,
		}
	})
}