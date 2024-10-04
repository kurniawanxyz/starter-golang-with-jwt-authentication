package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/helper"
	"time"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := helper.GetUser(c)
		if err != nil {
			return helper.HandleResponse(c, fiber.StatusUnauthorized, "Unauthorized",fiber.Map{"error": err.Error()})
		}
		
		// Check if the token is expired
		if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
			return helper.HandleResponse(c, fiber.StatusUnauthorized, "Unauthorized",fiber.Map{"error": "token is expired"})
		}

		c.Locals("user", claims.User)
		return c.Next()
	}
}
