package handler

import (
	"bus-api/database"
	"bus-api/model"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetRequestUserID(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := claims["user_id"].(float64)
	fmt.Println(id)
	fmt.Printf("t1: %T\n", id)
	return uint(id)
}

func FetchUserById(id uint) (model.User, error) {
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return user, errors.New("user not found")
	}
	return user, nil
}

func GetUserRoleFromCtx(c *fiber.Ctx) model.Role {
	id := GetRequestUserID(c)
	user, _ := FetchUserById(id)
	return user.Role
}
