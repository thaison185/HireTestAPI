package handlers

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/requests"
	"hiretest-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "invalid request body ")
	}

	if err := utils.Validate.Struct(req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "validation error "+err.Error())
	}

	result, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "login success", result)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req requests.RefreshTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "invalid request body ")
	}

	if err := utils.Validate.Struct(req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "validation error "+err.Error())
	}

	result, err := h.Service.Refresh(req.Token)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "token refreshed successfully", result)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "logout success", nil)
}
