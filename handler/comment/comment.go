package comment

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllCommentsFromCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all comments from card",
		})
	}
}

func CreateComment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create a comment",
		})
	}
}

func DeleteComment(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete a comment",
		})
	}
}
