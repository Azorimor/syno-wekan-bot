package controllers

import "github.com/gofiber/fiber/v2"

func GetRoot(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": 200,
		"msg":    "Server working",
	})
}
