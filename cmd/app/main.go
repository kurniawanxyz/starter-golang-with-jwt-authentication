package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/config"
	"github.com/kurniawanxzy/backend-olshop/helper"
)

func main() {
	app := fiber.New()
	config.Load()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return helper.HandleResponse(c, 404, "Route is not found", nil)
	})

	APP_LISTEN := fmt.Sprintf(":%d",config.ENV.AppPort)
	log.Fatal(app.Listen(APP_LISTEN))
}