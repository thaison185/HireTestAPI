package handlers

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProfileHandler struct {
	Service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{Service: service}
}

func (h *ProfileHandler) Get(c *fiber.Ctx) error {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return utils.Fail(c, fiber.StatusUnauthorized, "unauthorized")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return utils.Fail(c, fiber.StatusUnauthorized, "invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return utils.Fail(c, fiber.StatusUnauthorized, "invalid token subject")
	}

	profile, err := h.Service.GetCurrentProfile(userID)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "profile retrieved successfully", profile)
}
