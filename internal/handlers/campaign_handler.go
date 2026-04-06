package handlers

import "github.com/gofiber/fiber/v2"

type CampaignHandler struct{}

func (h *CampaignHandler) Create(c *fiber.Ctx) error { return ok(c, "campaign created", nil) }
func (h *CampaignHandler) List(c *fiber.Ctx) error { return ok(c, "campaign list", []string{}) }
func (h *CampaignHandler) Detail(c *fiber.Ctx) error { return ok(c, "campaign detail", fiber.Map{"id": c.Params("campaignId")}) }
func (h *CampaignHandler) Update(c *fiber.Ctx) error { return ok(c, "campaign updated", nil) }
func (h *CampaignHandler) Delete(c *fiber.Ctx) error { return ok(c, "campaign deleted", nil) }
