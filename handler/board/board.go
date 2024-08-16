package board

import (
	"fmt"
	"net/http"

	"github.com/fajarhidayad/thullo-trello-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetAllBoards(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var boards []models.Board

		db.Find(&boards)

		return c.JSON(Response{
			Data: boards,
		})
	}
}

func GetBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var board models.Board
		var err error
		var id int

		id, err = c.ParamsInt("id")
		if err != nil {
			return c.JSON(ErrorResponse{
				Message: "ID must be a number",
			})
		}

		err = db.First(&board, id).Error
		if err != nil {
			c.SendStatus(http.StatusNotFound)
			return c.JSON(ErrorResponse{
				Message: "Record not found",
			})
		}

		return c.JSON(Response{
			Data: board,
		})
	}
}

func CreateBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "create board",
		})
	}
}

func UpdateBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id int
		var err error

		id, err = c.ParamsInt("id")
		if err != nil {
			return c.JSON(ErrorResponse{
				Message: "ID must be a number",
			})
		}

		return c.JSON(Response{
			Data: fmt.Sprintf("update board %v", id),
		})
	}
}

func DeleteBoard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id int
		var err error

		id, err = c.ParamsInt("id")
		if err != nil {
			return c.JSON(ErrorResponse{
				Message: "ID must be a number",
			})
		}

		return c.JSON(Response{
			Data: fmt.Sprintf("delete board %v", id),
		})
	}
}

func GetAllBoardMembers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get all board members",
		})
	}
}

func AddBoardMember(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Add board member",
		})
	}
}

func DeleteBoardMember(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Delete board member",
		})
	}
}
