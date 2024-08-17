package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fajarhidayad/thullo-trello-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type signInForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type signUpForm struct {
	Firstname string `json:"firstname" validate:"required,min=3,max=50"`
	Lastname  string `json:"lastname"  validate:"required,min=3,max=50"`
	Email     string `json:"email"  validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=6,max=32"`
	Image     string `json:"image"`
}

var validate = validator.New()
var secretKey = os.Getenv("JWT_SECRET")

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(form interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(form)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func RegisterUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(signUpForm)
		if err := c.BodyParser(user); err != nil {
			return c.JSON(fiber.Map{
				"message": "Invalid body",
			})
		}
		errors := ValidateStruct(user)
		if errors != nil {
			return c.JSON(errors)
		}

		var exist models.User
		result := db.Model(models.User{}).Where("email = ?", user.Email).First(&exist)
		if result.RowsAffected > 0 {
			return c.JSON(fiber.Map{
				"message": "user already exist",
			})
		}

		hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		newUser := models.User{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Password:  string(hashPass),
			Image:     user.Image,
		}
		db.Select("Firstname", "Lastname", "Email", "Password", "Image").Create(&newUser)

		claims := jwt.MapClaims{
			"name": user.Email,
			"exp":  time.Now().Add(time.Hour * 2).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return c.JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.JSON(fiber.Map{
			"token": t,
		})
	}
}

func SignInUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(signInForm)
		if err := c.BodyParser(body); err != nil {
			return c.JSON(fiber.Map{
				"message": "Invalid body",
			})
		}
		errors := ValidateStruct(body)
		if errors != nil {
			return c.JSON(errors)
		}

		var user models.User
		result := db.Model(models.User{}).Where("email = ?", body.Email).First(&user)
		if result.Error != nil {
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
			return c.JSON(fiber.Map{
				"message": "credentials mismatched",
			})
		}

		claims := jwt.MapClaims{
			"name": body.Email,
			"exp":  time.Now().Add(time.Hour * 2).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return c.JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.JSON(fiber.Map{
			"token": t,
		})
	}
}
