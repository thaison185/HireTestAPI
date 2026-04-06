package repositories

type OrganizationRepository struct {
	BaseRepository
}

func NewOrganizationRepository(base BaseRepository) *OrganizationRepository {
	return &OrganizationRepository{BaseRepository: base}
}
