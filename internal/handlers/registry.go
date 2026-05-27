package handlers

import (
	"hiretest-api/configs"
	"hiretest-api/internal/repositories"
	"hiretest-api/internal/services"

	"gorm.io/gorm"
)

type Registry struct {
	Auth       *AuthHandler
	Profile    *ProfileHandler
	Org        *OrganizationHandler
	Question   *QuestionHandler
	Test       *TestHandler
	Campaign   *CampaignHandler
	Candidate  *CandidateHandler
	Invitation *InvitationHandler
	PublicInv  *PublicInvitationHandler
	PublicSub  *PublicSubmissionHandler
	Reviewer   *ReviewerSubmissionHandler
	Report     *ReportHandler
	Audit      *AuditHandler
}

func NewRegistry(db *gorm.DB, cfg configs.AppConfig) *Registry {
	baseRepo := repositories.BaseRepository{DB: db}
	authRepo := repositories.NewAuthRepository(baseRepo)
	questionRepo := repositories.NewQuestionRepository(baseRepo)

	authService := services.NewAuthService(authRepo, cfg)
	profileService := services.NewProfileService(authRepo)
	questionService := services.NewQuestionService(questionRepo)

	return &Registry{
		Auth:       NewAuthHandler(authService),
		Profile:    NewProfileHandler(profileService),
		Org:        &OrganizationHandler{},
		Question:   NewQuestionHandler(questionService),
		Test:       &TestHandler{},
		Campaign:   &CampaignHandler{},
		Candidate:  &CandidateHandler{},
		Invitation: &InvitationHandler{},
		PublicInv:  &PublicInvitationHandler{},
		PublicSub:  &PublicSubmissionHandler{},
		Reviewer:   &ReviewerSubmissionHandler{},
		Report:     &ReportHandler{},
		Audit:      &AuditHandler{},
	}
}
