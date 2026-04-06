package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler struct{}

func (h *AuthHandler) Login(c *fiber.Ctx) error   { return ok(c, "login success", fiber.Map{"token": "TODO"}) }
func (h *AuthHandler) Refresh(c *fiber.Ctx) error { return ok(c, "refresh success", fiber.Map{"token": "TODO"}) }
func (h *AuthHandler) Logout(c *fiber.Ctx) error  { return ok(c, "logout success", nil) }
