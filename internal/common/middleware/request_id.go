package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-Request-ID", uuid.NewString())
		return c.Next()
	}
}
