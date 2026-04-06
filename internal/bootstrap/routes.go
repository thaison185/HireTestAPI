package bootstrap

import (
	"hiretest-api/configs"
	"hiretest-api/internal/handlers"
	"hiretest-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg configs.AppConfig) {
	h := handlers.NewRegistry(db, cfg)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	routes.RegisterHealthRoutes(app)
	routes.RegisterAuthRoutes(v1, h)
	routes.RegisterProfileRoutes(v1, h)
	routes.RegisterOrganizationRoutes(v1, h)
	routes.RegisterQuestionRoutes(v1, h)
	routes.RegisterTestRoutes(v1, h)
	routes.RegisterCampaignRoutes(v1, h)
	routes.RegisterCandidateRoutes(v1, h)
	routes.RegisterInvitationRoutes(v1, h)
	routes.RegisterReportRoutes(v1, h)
	routes.RegisterAuditRoutes(v1, h)
	routes.RegisterPublicRoutes(v1, h)
	routes.RegisterReviewerRoutes(v1, h)
}
