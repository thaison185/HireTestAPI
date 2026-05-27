package bootstrap

import (
	"hiretest-api/configs"
	"hiretest-api/internal/common/middleware"
	"hiretest-api/internal/handlers"
	"hiretest-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg configs.AppConfig) {
	h := handlers.NewRegistry(db, cfg)

	routes.RegisterHealthRoutes(app)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	public := v1.Group("/public")
	auth := v1.Group("/auth")
	profile := protected(v1, "/profile", cfg)
	question := protected(v1, "/questions", cfg)
	protected := protected(v1, "/protected", cfg)
	reviewer := v1.Group("/reviewer", middleware.JWT(cfg.Env.JWTSecret))

	routes.RegisterAuthRoutes(auth, h)
	routes.RegisterPublicRoutes(public, h)

	routes.RegisterProfileRoutes(profile, h)
	routes.RegisterOrganizationRoutes(protected, h)
	routes.RegisterQuestionRoutes(question, h)
	routes.RegisterTestRoutes(protected, h)
	routes.RegisterCampaignRoutes(protected, h)
	routes.RegisterCandidateRoutes(protected, h)
	routes.RegisterInvitationRoutes(protected, h)
	routes.RegisterReportRoutes(protected, h)
	routes.RegisterAuditRoutes(protected, h)

	routes.RegisterReviewerRoutes(reviewer, h)
}

func protected(v1 fiber.Router, prefix string, cfg configs.AppConfig) fiber.Router {
	return v1.Group(prefix, middleware.JWT(cfg.Env.JWTSecret))
}
