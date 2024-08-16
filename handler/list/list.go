package list

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "create list",
		})
	}
}

func GetListsFromBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all list from board id",
		})
	}
}

func GetListDetails(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get list details",
		})
	}
}

func UpdateList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Update list",
		})
	}
}

func DeleteList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete list",
		})
	}
}
