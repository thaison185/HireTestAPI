package repositories

type CandidateRepository struct {
	BaseRepository
}

func NewCandidateRepository(base BaseRepository) *CandidateRepository {
	return &CandidateRepository{BaseRepository: base}
}
