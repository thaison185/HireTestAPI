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

	switch query.Status {
	case "inactive":
		dbQuery = dbQuery.Where("is_active = ?", false)
	case "all":
		// no filter
	default:
		dbQuery = dbQuery.Where("is_active = ?", true)
	}

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
	order := getOrder(query.Sort)

	if err := dbQuery.
		Order(order).
		Offset(offset).
		Limit(query.PageSize).
		Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

func getOrder(sort string) string {
	switch sort {
	case "created_at.asc":
		return "created_at ASC"
	case "title.asc":
		return "title ASC"
	case "title.desc":
		return "title DESC"
	default:
		return "created_at DESC"
	}
}

func (r *QuestionRepository) FindByID(id string) (*models.Question, error) {
	var question models.Question
	if err := r.DB.Where("id = ? AND is_active = ?", id, true).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) FindOwnedByID(id string, ownerID string) (*models.Question, error) {
	var question models.Question
	if err := r.DB.Where("id = ? AND created_by = ? AND is_active = ?", id, ownerID, true).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) FindAnyByID(id string) (*models.Question, error) {
	var question models.Question
	if err := r.DB.Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) Create(question *models.Question) error {
	return r.DB.Create(question).Error
}

func (r *QuestionRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.Question{}).Where("id = ? AND is_active = ?", id, true).
		Updates(updates).Error
}

func (r *QuestionRepository) Delete(id string) error {
	return r.DB.Model(&models.Question{}).Where("id = ? AND is_active = ?", id, true).
		Updates(map[string]interface{}{"is_active": false}).Error
}

func (r *QuestionRepository) Restore(id string) error {
	return r.DB.Model(&models.Question{}).Where("id = ? AND is_active = ?", id, false).
		Updates(map[string]interface{}{"is_active": true}).Error
}
