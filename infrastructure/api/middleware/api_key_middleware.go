package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kurniawanxzy/backend-olshop/config"
	"github.com/kurniawanxzy/backend-olshop/helper"
)

func ApiKeyMiddleware(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-key")
		if strings.Compare(apiKey, config.ENV.APIKey) != 0 {
			
			return helper.HandleResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
		}
		return c.Next()
}