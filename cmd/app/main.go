package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/config"
)

func main() {
	app := fiber.New()
	config.Load()

	
	APP_LISTEN := fmt.Sprintf(":%d",config.ENV.AppPort)
	log.Fatal(app.Listen(APP_LISTEN))
}