package routes

import (
	"github.com/Azorimor/syno-wekan-bot/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RootRoutes(a *fiber.App) {
	a.Get("/", controllers.GetRoot)
}