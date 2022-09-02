package middleware

import (
	"bus-api/database"
	"bus-api/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {
	db := database.DB
	var profiles []model.BusOperatorProfile
	db.Model(&model.BusOperatorProfile{}).Preload("User").Find(&profiles)
	return c.JSON(fiber.Map{"status": "success", "message": "All Operators", "data": profiles})
}

func UnauthorizedHandler(c *fiber.Ctx) error {
	return c.Status(401).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
}

// // Protected protect routes
// func IsAdmin(handler fiber.Handler) fiber.Handler {
// 	if false {
// 		return handler
// 	} else {
// 		return UnauthorizedHandler
// 	}
// }

// Protected protect routes
func IsAdmin() bool {
	return true

}

func IsCustomer() bool {
	return true
}

// Protected role
func RoleMiddleware(rules ...func() bool) fiber.Handler {
	result := true
	for i := 0; i < len(rules); i++ {
		fmt.Println(rules[i]())
		result = rules[i]() && result
	}
	// fmt.Println(rules)
	if result {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	} else {
		return UnauthorizedHandler
	}
}
