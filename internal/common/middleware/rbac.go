package middleware

import (
	"hiretest-api/internal/common/utils"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(allowedRoles ...string) fiber.Handler {
	allowedRolesSet := make(map[string]struct{}, len(allowedRoles))
	for _, role := range allowedRoles {
		allowedRolesSet[role] = struct{}{}
	}

	return func(c *fiber.Ctx) error {
		userRole, err := utils.CurrentUserRole(c)
		if err != nil {
			return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
		}

		if _, allowed := allowedRolesSet[userRole]; allowed {
			return c.Next()
		}

		return utils.Fail(c, fiber.StatusForbidden, "forbidden: insufficient permissions")
	}
}
