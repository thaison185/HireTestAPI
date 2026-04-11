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

	public := v1.Group("/")
	protected := v1.Group("/", middleware.JWT(cfg.Env.JWTSecret))
	reviewer := v1.Group("/reviewer", middleware.JWT(cfg.Env.JWTSecret))

	routes.RegisterAuthRoutes(public, h)
	routes.RegisterPublicRoutes(public, h)

	routes.RegisterProfileRoutes(protected, h)
	routes.RegisterOrganizationRoutes(protected, h)
	routes.RegisterQuestionRoutes(protected, h)
	routes.RegisterTestRoutes(protected, h)
	routes.RegisterCampaignRoutes(protected, h)
	routes.RegisterCandidateRoutes(protected, h)
	routes.RegisterInvitationRoutes(protected, h)
	routes.RegisterReportRoutes(protected, h)
	routes.RegisterAuditRoutes(protected, h)

	routes.RegisterReviewerRoutes(reviewer, h)
}
