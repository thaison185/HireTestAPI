package repositories

type ProfileRepository struct {
	BaseRepository
}

func NewProfileRepository(base BaseRepository) *ProfileRepository {
	return &ProfileRepository{BaseRepository: base}
}
