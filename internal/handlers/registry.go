package handlers

import (
	"hiretest-api/configs"

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
	_ = db
	_ = cfg
	return &Registry{
		Auth:       &AuthHandler{},
		Profile:    &ProfileHandler{},
		Org:        &OrganizationHandler{},
		Question:   &QuestionHandler{},
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
