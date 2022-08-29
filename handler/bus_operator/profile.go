package handler

import (
	"bus-api/database"
	"bus-api/model"
	"fmt"
	"net/mail"

	// "bus-api/handler/bus_operator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetAllProducts query all BusOperatorProfile
func GetAllBusOperatorProfile(c *fiber.Ctx) error {
	// user := c.Locals("user").(*jwt.Token)
	// fmt.Println(user)
	// claims := user.Claims.(jwt.MapClaims)
	// fmt.Println(claims)
	// id := claims["user_id"]
	// fmt.Println(id)
	// fmt.Printf("t1: %T\n", id)

	db := database.DB
	var profiles []model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("User").Find(&profiles)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profiles})
}

// GetAllProducts query all BusOperatorProfile
func GetCurrentBusOperatorProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := claims["user_id"]
	fmt.Println(id)
	fmt.Printf("t1: %T\n", id)

	db := database.DB
	var profile model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("User").Where("user_id = ?", id).Find(&profile)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profile})
}

// Create new BusOperatorProfile
func CreateBusOperatorProfile(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create User for BusOperatorProfile", "data": err})
	}
	user.Role = model.ROLE_CUSTOMER
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "invalid email supplied", "data": err})
	}

	result := db.Create(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create BusOperatorProfile", "data": result.Error})
	}

	profile := new(model.BusOperatorProfile)
	profile.User = *user
	if err := c.BodyParser(profile); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create BusOperatorProfile", "data": err})
	}
	profile.ApprovalStatus = model.STATUS_PENDING_APPROVAL
	profile.RejectionComment = ""
	profile.Ratings = -1

	result = db.Omit("BusinessLogoId").Create(&profile)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create BusOperatorProfile", "data": result.Error})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created BusOperatorProfile", "data": profile})
}
