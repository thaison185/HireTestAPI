package repositories

type RbacRepository struct {
	BaseRepository
}

func NewRbacRepository(base BaseRepository) *RbacRepository {
	return &RbacRepository{BaseRepository: base}
}
