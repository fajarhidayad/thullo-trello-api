package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fajarhidayad/thullo-trello-api/models"
	"github.com/fajarhidayad/thullo-trello-api/routes"
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

	var DB_USER string = os.Getenv("DB_USER")
	var DB_PASS string = os.Getenv("DB_PASSWORD")
	var DB_HOST string = os.Getenv("DB_HOST")
	var DB_PORT string = os.Getenv("DB_PORT")
	var DB_NAME string = os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect db: ", err.Error())
	}

	models.Migrate(DB)
}

// @title Thullo-Trello API
// @version 1.0
// @description Trello clone api using go-fiber.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://www.suryadev.my.id
// @contact.email fajarsuryahidayad@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1234
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// HealthCheck godoc
	// @Summary Show the status of server.
	// @Description get the status of server.
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router / [get]
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	routes.Register(v1, DB)

	PORT := os.Getenv("PORT")
	if err := app.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		panic(err)
	}
}
