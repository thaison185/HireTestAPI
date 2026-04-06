package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCampaignRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/campaigns")
	r.Post("/", h.Campaign.Create)
	r.Get("/", h.Campaign.List)
	r.Get("/:campaignId", h.Campaign.Detail)
	r.Patch("/:campaignId", h.Campaign.Update)
	r.Delete("/:campaignId", h.Campaign.Delete)
}
