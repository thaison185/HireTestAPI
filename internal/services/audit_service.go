package services

import (
	"hiretest-api/internal/models"
	"hiretest-api/internal/repositories"

	"github.com/google/uuid"
)

type AuditService struct {
	Repository *repositories.AuditRepository
}

func NewAuditService(repo *repositories.AuditRepository) *AuditService {
	return &AuditService{Repository: repo}
}

func (s *AuditService) Log(actorID, action, entityType, entityID, description string) error {
	log := &models.AuditLog{
		ID:          uuid.NewString(),
		ActorID:     actorID,
		Action:      action,
		EntityType:  entityType,
		EntityID:    entityID,
		Description: description,
	}
	return s.Repository.Create(log)
}
