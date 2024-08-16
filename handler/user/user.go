package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUserInfo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get user info",
		})
	}
}

func CreateNewUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create new user",
		})
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Update user profile",
		})
	}
}
