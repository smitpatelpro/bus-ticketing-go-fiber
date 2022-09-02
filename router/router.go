package router

import (
	"bus-api/handler"
	operator_handler "bus-api/handler/bus_operator"
	"bus-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/refresh_token", handler.RefreshToken)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	product := api.Group("/product")
	product.Get("/", middleware.Protected(), handler.GetAllProducts)
	// product.Get("/", middleware.Protected(), middleware.RoleMiddleware(middleware.IsAdmin, middleware.IsCustomer), handler.GetAllProducts)
	product.Get("/:id", middleware.Protected(), handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)

	// Bus Operator
	opertor := api.Group("/bus_operator")
	opertor.Get("/", middleware.Protected(), operator_handler.GetAllBusOperatorProfile)
	opertor.Post("/", middleware.Protected(), operator_handler.CreateBusOperatorProfile)
	opertor.Get("/profile", middleware.Protected(), operator_handler.GetCurrentBusOperatorProfile)
	opertor.Get("/profile/logo", middleware.Protected(), operator_handler.GetProfileBusOperatorProfileLogo)
	opertor.Patch("/profile/logo", middleware.Protected(), operator_handler.SetProfileBusOperatorProfileLogo)

}
