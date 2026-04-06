package repositories

type QuestionRepository struct {
	BaseRepository
}

func NewQuestionRepository(base BaseRepository) *QuestionRepository {
	return &QuestionRepository{BaseRepository: base}
}
