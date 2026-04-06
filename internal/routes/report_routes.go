package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterReportRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/reports")
	r.Get("/campaigns/:campaignId/summary", h.Report.CampaignSummary)
	r.Get("/tests/:testId/summary", h.Report.TestSummary)
}
