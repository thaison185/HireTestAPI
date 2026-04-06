package handlers

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/requests"
	"hiretest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func (h *AuthHandler) Login(c *fiber.Ctx) error   { return utils.Success(c, 200, "login success", fiber.Map{"token": "TODO"}) }
func (h *AuthHandler) Refresh(c *fiber.Ctx) error { return utils.Success(c, 200, "refresh success", fiber.Map{"token": "TODO"}) }
func (h *AuthHandler) Logout(c *fiber.Ctx) error  { return utils.Success(c, 200, "logout success", nil) }
