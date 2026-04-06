package repositories

type AuthRepository struct {
	BaseRepository
}

func NewAuthRepository(base BaseRepository) *AuthRepository {
	return &AuthRepository{BaseRepository: base}
}
