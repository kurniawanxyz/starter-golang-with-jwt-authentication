package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/domain/usecases"
	"github.com/kurniawanxzy/backend-olshop/helper"
	"github.com/kurniawanxzy/backend-olshop/infrastructure/api/middleware"
	"github.com/kurniawanxzy/backend-olshop/requests"
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

func (ar *AuthRoute) VerifyUser(c *fiber.Ctx) error {
	token := new(requests.VerifyUserRequest)
	
	if err := c.BodyParser(&token); err != nil {
		return helper.HandleResponse(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	if err := ar.Validate.Struct(token); err != nil {
		return helper.HandleValidationMessage(c,err)
	}
	
	if err := ar.AuthUseCase.VerifyUser(token.Token, token.Email); err != nil {
		return helper.HandleResponse(c, 500, "Failed to verify user", err.Error())
	}
	return helper.HandleResponse(c, fiber.StatusOK, "User verified successfully", nil)
}

func (ar *AuthRoute) CreateToken(c *fiber.Ctx) error {
	
	data := new(requests.CreateTokenRequest)
	
	if err := c.BodyParser(&data); err != nil {
		return helper.HandleResponse(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	if err := ar.Validate.Struct(data); err != nil {
		return helper.HandleValidationMessage(c,err)
	}

	if err := ar.AuthUseCase.CreateToken(data.Email, data.Type); err != nil {
		return helper.HandleResponse(c, 400, "Failed to create token verification", err.Error())
	}
	return helper.HandleResponse(c, fiber.StatusOK, "Token verification created successfully", nil)
}

func (ar *AuthRoute) Login(c *fiber.Ctx) error {
	data := new(requests.LoginRequest)
	
	if err := c.BodyParser(&data); err != nil {
		return helper.HandleResponse(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	if err := ar.Validate.Struct(data); err != nil {
		return helper.HandleValidationMessage(c,err)
	}

	token, err := ar.AuthUseCase.Login(data)
	if err != nil {
		return helper.HandleResponse(c, 400, "Failed to login", err.Error())
	}
	return helper.HandleResponse(c, fiber.StatusOK, "Login success", fiber.Map{"token": token})
}

func (ar *AuthRoute) ResetPassword(c *fiber.Ctx) error {
	data := new(requests.ResetPasswordRequest)
	
	if err := c.BodyParser(&data); err != nil {
		return helper.HandleResponse(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	if err := ar.Validate.Struct(data); err != nil {
		return helper.HandleValidationMessage(c,err)
	}

	user, err := helper.GetUser(c)
	if err != nil {
		return helper.HandleResponse(c, fiber.StatusUnauthorized, "Unauthorized", err)
	}

	if err := ar.AuthUseCase.ResetPassword(data, user.User); err != nil {
		return helper.HandleResponse(c, 400, "Failed to reset password", err.Error())
	}
	return helper.HandleResponse(c, fiber.StatusOK, "Password reset success", nil)
}

func SetupAuthRoute(r fiber.Router, authUseCase *AuthRoute) {
	r.Post("/register", authUseCase.RegisterUser)
	r.Post("/verify", authUseCase.VerifyUser)
	r.Post("/request-token", authUseCase.CreateToken)
	r.Post("/login", authUseCase.Login)
	r.Use(middleware.AuthMiddleware()).Post("/reset-password", authUseCase.ResetPassword)
}