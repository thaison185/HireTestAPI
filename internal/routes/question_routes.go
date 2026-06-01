package routes

import (
	"hiretest-api/internal/common/constants"
	"hiretest-api/internal/common/middleware"
	"hiretest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionRoutes(r fiber.Router, h *handlers.Registry) {
	r.Post("/", middleware.RequireRole(
		constants.RoleAdmin,
		constants.RoleRecruiter), h.Question.Create)
	r.Get("/", middleware.RequireRole(
		constants.RoleAdmin,
		constants.RoleRecruiter,
		constants.RoleReviewer), h.Question.List)
	r.Get("/:id", middleware.RequireRole(
		constants.RoleAdmin,
		constants.RoleRecruiter,
		constants.RoleReviewer), h.Question.Detail)
	r.Patch("/:id", middleware.RequireRole(
		constants.RoleAdmin,
		constants.RoleRecruiter), h.Question.Update)
	r.Delete("/:id", middleware.RequireRole(
		constants.RoleAdmin,
		constants.RoleRecruiter), h.Question.Delete)
	r.Patch("/:id/restore", middleware.RequireRole(
		constants.RoleAdmin), h.Question.Restore)
	// r.Post("/:id/publish", h.Question.Publish)
}
