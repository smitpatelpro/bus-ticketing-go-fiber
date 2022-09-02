package handler

import (
	"bus-api/database"
	"bus-api/handler"
	"bus-api/model"
	"bus-api/utils"
	"fmt"
	"math/rand"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all BusOperatorProfile
func GetAllBusOperatorProfile(c *fiber.Ctx) error {
	if handler.GetUserRoleFromCtx(c) != model.ROLE_BUS_OPERATOR {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

	db := database.DB
	var profiles []model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("User").Find(&profiles)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profiles})
}

// GetAllProducts query all BusOperatorProfile
func GetCurrentBusOperatorProfile(c *fiber.Ctx) error {
	id := handler.GetRequestUserID(c)
	user, _ := handler.FetchUserById(id)
	if user.Role != model.ROLE_ADMIN {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

	db := database.DB
	var profile model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("User").Where("user_id = ?", id).Find(&profile)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profile})
}

// Create new BusOperatorProfile
func CreateBusOperatorProfile(c *fiber.Ctx) error {
	if handler.GetUserRoleFromCtx(c) != model.ROLE_BUS_OPERATOR {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

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

// Get Logo of BusOperator
func GetProfileBusOperatorProfileLogo(c *fiber.Ctx) error {
	id := handler.GetRequestUserID(c)
	user, _ := handler.FetchUserById(id)
	if user.Role != model.ROLE_ADMIN {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

	db := database.DB
	var profile model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("BusinessLogo").Where("user_id = ?", id).Find(&profile)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profile.BusinessLogo})
}

// Get Logo of BusOperator
func SetProfileBusOperatorProfileLogo(c *fiber.Ctx) error {
	id := handler.GetRequestUserID(c)
	user, _ := handler.FetchUserById(id)
	if user.Role != model.ROLE_ADMIN {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}
	file_path := ""
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["file"]

		rand.Seed(time.Now().UnixNano())
		for _, file := range files {
			path, e := utils.SaveStaticFile(c, file, "operator_logos")
			if e != nil {
				return c.Status(401).JSON(fiber.Map{"status": "error", "message": e, "data": nil})
			}
			file_path = path
		}
	} else {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "no file supplied", "data": nil})
	}

	db := database.DB
	media := model.Media{Path: file_path}

	db.Create(&media) // pass pointer of data to Create

	fmt.Println(media)

	var profile model.BusOperatorProfile
	// profile.BusinessLogo = media
	// profile.BusinessLogoId = media.ID
	// db.Save(&profile)
	db.Model(&profile).Where("user_id = ?", id).Update("business_logo_id", media.ID)
	// db.Model(&model.BusOperatorProfile{}).Preload("BusinessLogo").Where("user_id = ?", id).Find(&profile)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profile.BusinessLogo})
}
