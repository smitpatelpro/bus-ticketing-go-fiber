package handler

import (
	"bus-api/database"
	"bus-api/model"
	"bus-api/services"
	"errors"
	"fmt"
	// "strings"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt"
	// "github.com/gofiber/jwt/v3"

	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("err = ",err, "ref:", password, " - ", hash)
	return err == nil
}

// func getUserByEmail(e string) (*model.User, error) {
// 	db := database.DB
// 	var user model.User
// 	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

func getUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.Identity
	pass := input.Password

	// email, err := getUserByEmail(identity)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	// }
	// fmt.Println("email=", email.Username)
	user, err := getUserByUsername(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}
	fmt.Println("user=", user.Username)

	// if email == nil && user == nil {
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	ud = UserData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	// if user != nil {
	// 	fmt.Println("using user")
	// 	ud = UserData{
	// 		ID:       user.ID,
	// 		Username: user.Username,
	// 		Email:    user.Email,
	// 		Password: user.Password,
	// 	}
	// } else {
	// 	ud = UserData{
	// 		ID:       email.ID,
	// 		Username: email.Username,
	// 		Email:    email.Email,
	// 		Password: email.Password,
	// 	}
	// }
	fmt.Println(pass, " - ", ud.Password)
	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte(config.Config("SECRET")))
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	// Separate JWT Service Code
	jwt := services.Jwt{}
	t, err := jwt.CreateToken(*user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "access_token": t})
}

// Login get user and password
func RefreshToken(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	auth_token := auth[7:]

	fmt.Println(auth_token)

	type LoginInput struct {
		RefreshToken string `json:"refresh_token"`
	}
	// var input LoginInput
	input := LoginInput{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error in input data format", "data": err})
	}

	fmt.Println(input.RefreshToken)
	jwt := services.Jwt{}

	token := services.Token{}
	token.RefreshToken = input.RefreshToken
	token.AccessToken = auth_token

	user_id, err := jwt.ValidateRefreshToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid refresh token", "data": nil})
	}
	fmt.Println(user_id)
	db := database.DB

	user := model.User{}
	db.First(&user, user_id)

	t, err := jwt.CreateToken(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Token refreshed", "access_token": t})
}
