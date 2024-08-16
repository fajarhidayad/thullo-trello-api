package card

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create card",
		})
	}
}

func GetAllCardsFromList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all cards from list",
		})
	}
}

func GetCardDetails(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get cards details",
		})
	}
}

func UpdateCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Update cards",
		})
	}
}

func DeleteCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete cards",
		})
	}
}

func GetAllMembersInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all card member",
		})
	}
}

func AddMemberInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Add card member",
		})
	}
}

func DeleteMemberInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete card member",
		})
	}
}

func GetAllAttachmentsInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all attachments in card",
		})
	}
}

func AddAttachmentInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Add attachment in card",
		})
	}
}

func DeleteAttachmentInCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete attachment in card",
		})
	}
}
