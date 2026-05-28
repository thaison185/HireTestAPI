package handlers

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	Service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{Service: service}
}

func (h *ProfileHandler) Get(c *fiber.Ctx) error {
	userID, err := utils.CurrentUserID(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	profile, err := h.Service.GetCurrentProfile(userID)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "profile retrieved successfully", profile)
}
