package handlers

import "github.com/gofiber/fiber/v2"

type OrganizationHandler struct{}

func (h *OrganizationHandler) ListMembers(c *fiber.Ctx) error { return ok(c, "member list", []string{}) }
func (h *OrganizationHandler) UpdateMemberRoles(c *fiber.Ctx) error { return ok(c, "member roles updated", nil) }
