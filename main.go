package main

import (
	"bus-api/database"
	"bus-api/router"
	"bus-api/utils"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	app := fiber.New(config)
	app.Use(cors.New())

	app.Static("/static", "./static")
	utils.CreateMediaDirectories()

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
