package repositories

type ReviewRepository struct {
	BaseRepository
}

func NewReviewRepository(base BaseRepository) *ReviewRepository {
	return &ReviewRepository{BaseRepository: base}
}
