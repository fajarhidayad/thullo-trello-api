package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Register new user",
		})
	}
}

func SignInUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Sign in user",
		})
	}
}
