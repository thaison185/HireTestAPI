package repositories

type AuditRepository struct {
	BaseRepository
}

func NewAuditRepository(base BaseRepository) *AuditRepository {
	return &AuditRepository{BaseRepository: base}
}
