package repositories

type SubmissionRepository struct {
	BaseRepository
}

func NewSubmissionRepository(base BaseRepository) *SubmissionRepository {
	return &SubmissionRepository{BaseRepository: base}
}
