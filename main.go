package main

import (
	"github.com/Azorimor/syno-wekan-bot/pkg/configs"
	"github.com/Azorimor/syno-wekan-bot/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	// Middleware

	// Routes
	routes.RootRoutes(app)

	// Start
	app.Listen(":3003")
}