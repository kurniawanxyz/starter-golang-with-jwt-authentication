package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/domain/usecases"
	"github.com/kurniawanxzy/backend-olshop/helper"
)

type AuthRoute struct {
	AuthUseCase *usecases.AuthUseCase
	Validate *validator.Validate
}

func NewAuthRoute(authUseCase *usecases.AuthUseCase) *AuthRoute {
	return &AuthRoute{
		AuthUseCase: authUseCase,
		Validate: validator.New(), 
	}
}

func (ar *AuthRoute) RegisterUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(&user); err != nil {
		return helper.HandleResponse(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	

	if err := ar.Validate.Struct(user); err != nil {

		return helper.HandleValidationMessage(c,err)
	}

	if err := ar.AuthUseCase.RegisterUser(user); err != nil {
		return helper.HandleResponse(c, 500, "Failed to register user", err.Error())

	}
	return helper.HandleResponse(c, fiber.StatusCreated, "User registered successfully", nil)
}

func SetupAuthRoute(r fiber.Router, authUseCase *AuthRoute) {
	r.Post("/register", authUseCase.RegisterUser)
}