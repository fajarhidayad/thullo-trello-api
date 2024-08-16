package label

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNewLabel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create new Label",
		})
	}
}

func GetAllLabelsInBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all labels in board",
		})
	}
}

func DeleteLabel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete Label",
		})
	}
}
