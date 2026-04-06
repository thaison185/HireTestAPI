package handlers

import "github.com/gofiber/fiber/v2"

type ReportHandler struct{}

func (h *ReportHandler) CampaignSummary(c *fiber.Ctx) error { return ok(c, "campaign summary", nil) }
func (h *ReportHandler) TestSummary(c *fiber.Ctx) error { return ok(c, "test summary", nil) }
