package utils

import (
	"errors"

	code_errors "hiretest-api/internal/common/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CurrentUserID(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		return "", errors.New(code_errors.CodeUnauthorized)
	}
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New(code_errors.CodeInvalidUserClaims)
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return "", errors.New(code_errors.CodeInvalidUserID)
	}
	return userID, nil
}

func CurrentUserRole(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		return "", errors.New(code_errors.CodeUnauthorized)
	}
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New(code_errors.CodeInvalidUserClaims)
	}
	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return "", errors.New(code_errors.CodeInvalidUserRole)
	}
	return role, nil
}
