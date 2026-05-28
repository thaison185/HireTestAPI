package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CurrentUserID(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		return "", errors.New("unauthorized")
	}
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid user claims")
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return "", errors.New("invalid user ID")
	}
	return userID, nil
}

func CurrentUserRole(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		return "", errors.New("unauthorized")
	}
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid user claims")
	}
	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return "", errors.New("invalid user role")
	}
	return role, nil
}
