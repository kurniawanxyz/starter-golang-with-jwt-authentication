package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kurniawanxzy/backend-olshop/config"
	"github.com/kurniawanxzy/backend-olshop/domain/usecases"
	"github.com/kurniawanxzy/backend-olshop/helper"
	"github.com/kurniawanxzy/backend-olshop/infrastructure/api/middleware"
	"github.com/kurniawanxzy/backend-olshop/infrastructure/api/routes"
	"github.com/kurniawanxzy/backend-olshop/infrastructure/database"
	"github.com/kurniawanxzy/backend-olshop/repository"
	"github.com/kurniawanxzy/backend-olshop/service"
)

func main() {
	app := fiber.New()
	config.Load()
	database.Load()

	app.Use(logger.New())
	app.Use(middleware.ApiKeyMiddleware)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})


	// repository
	userRepo  := repository.NewUserRepository(database.DB)
	tokenRepo := repository.NewTokenVerificationRepository(database.DB)

	// service
	authService  := service.NewAuthService(database.DB, userRepo, tokenRepo)


	// usecase
	authUseCase := usecases.NewAuthUseCase(authService)

	// handler & route
	api := app.Group("/api")

	authHandler := routes.NewAuthRoute(authUseCase)
	routes.SetupAuthRoute(api, authHandler)


	app.Use(func(c *fiber.Ctx) error {
		return helper.HandleResponse(c, 404, "Route is not found", nil)
	})


	APP_LISTEN := fmt.Sprintf(":%s",config.ENV.AppPort)
	log.Fatal(app.Listen(APP_LISTEN))
}