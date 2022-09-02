package main

import (
	"bus-api/config"
	"bus-api/database"
	"bus-api/router"
	"bus-api/utils"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fiber_config := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	app := fiber.New(fiber_config)
	app.Use(cors.New())

	app.Static("/static", config.Config("STATIC_ROOT"))
	app.Static("/media", config.Config("MEDIA_ROOT"))
	utils.CreateMediaDirectories()

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
