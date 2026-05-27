package repositories

import (
	"hiretest-api/internal/models"
	"hiretest-api/internal/requests"
)

type QuestionRepository struct {
	BaseRepository
}

func NewQuestionRepository(base BaseRepository) *QuestionRepository {
	return &QuestionRepository{BaseRepository: base}
}

func (r *QuestionRepository) List(query requests.ListQuestionRequest) ([]models.Question, int64, error) {
	var questions []models.Question
	var total int64

	dbQuery := r.DB.Model(&models.Question{})

	if query.Keyword != "" {
		likeQuery := "%" + query.Keyword + "%"
		dbQuery = dbQuery.Where("title ILIKE ? OR description ILIKE ?", likeQuery, likeQuery)
	}

	if query.Type != "" {
		dbQuery = dbQuery.Where("type = ?", query.Type)
	}

	if query.Level != "" {
		dbQuery = dbQuery.Where("level = ?", query.Level)
	}

	if query.Category != "" {
		dbQuery = dbQuery.Where("category = ?", query.Category)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize

	if err := dbQuery.
		Order("created_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

func (r *QuestionRepository) FindByID(id string) (*models.Question, error) {
	var question models.Question
	if err := r.DB.Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
