package services

import "hiretest-api/internal/repositories"

type Dependencies struct {
	Auth        *repositories.AuthRepository
	User        *repositories.UserRepository
	Organization *repositories.OrganizationRepository
	RBAC        *repositories.RbacRepository
	Question    *repositories.QuestionRepository
	Test        *repositories.TestRepository
	Campaign    *repositories.CampaignRepository
	Candidate   *repositories.CandidateRepository
	Invitation  *repositories.InvitationRepository
	Submission  *repositories.SubmissionRepository
	Review      *repositories.ReviewRepository
	Report      *repositories.ReportRepository
	Audit       *repositories.AuditRepository
}
