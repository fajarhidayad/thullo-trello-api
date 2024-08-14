package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fajarhidayad/thullo-trello-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "127.0.0.1", "3306", "thullo")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect db: ", err.Error())
	}

	models.Migrate(DB)
}

func main() {
	app := fiber.New()
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})

	PORT := os.Getenv("PORT")

	if err := app.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		panic(err)
	}
}
