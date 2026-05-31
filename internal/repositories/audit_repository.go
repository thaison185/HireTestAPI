package repositories

import (
	"hiretest-api/internal/models"
)

type AuditRepository struct {
	BaseRepository
}

func NewAuditRepository(base BaseRepository) *AuditRepository {
	return &AuditRepository{BaseRepository: base}
}

func (r *AuditRepository) Create(log *models.AuditLog) error {
	return r.DB.Create(log).Error
}
