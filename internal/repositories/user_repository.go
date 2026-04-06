package repositories

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(base BaseRepository) *UserRepository {
	return &UserRepository{BaseRepository: base}
}
