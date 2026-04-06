package repositories

type TestRepository struct {
	BaseRepository
}

func NewTestRepository(base BaseRepository) *TestRepository {
	return &TestRepository{BaseRepository: base}
}
