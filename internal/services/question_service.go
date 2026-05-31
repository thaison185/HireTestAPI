package services

import (
	"errors"
	"strings"

	"hiretest-api/internal/common/constants"
	"hiretest-api/internal/models"
	"hiretest-api/internal/repositories"
	"hiretest-api/internal/requests"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionService struct {
	Repository   *repositories.QuestionRepository
	AuditService *AuditService
}

type ListQuestionResponse struct {
	Items []models.Question `json:"items"`
	Meta  PaginationMeta    `json:"meta"`
}

type PaginationMeta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

func NewQuestionService(repo *repositories.QuestionRepository, auditService *AuditService) *QuestionService {
	return &QuestionService{
		Repository:   repo,
		AuditService: auditService,
	}
}

func (s *QuestionService) List(query requests.ListQuestionRequest) (*ListQuestionResponse, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.PageSize > 100 {
		query.PageSize = 100
	}

	if query.Sort == "" {
		query.Sort = "created_at.desc"
	}

	if query.Status == "" {
		query.Status = "active"
	}

	questions, total, err := s.Repository.List(query)
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(query.PageSize) - 1) / int64(query.PageSize))
	response := &ListQuestionResponse{
		Items: questions,
		Meta: PaginationMeta{
			Total:      total,
			Page:       query.Page,
			PageSize:   query.PageSize,
			TotalPages: totalPages,
			HasNext:    query.Page < totalPages,
			HasPrev:    query.Page > 1,
		},
	}

	return response, nil
}

func (s *QuestionService) GetByID(id string) (*models.Question, error) {
	if id == "" {
		return nil, errors.New("question ID is required")
	}

	question, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, errors.New("question not found")
	}
	return question, nil
}

func (s *QuestionService) Create(req requests.CreateQuestionRequest, createdBy string) (*models.Question, error) {
	if createdBy == "" {
		return nil, errors.New("createdBy is required")
	}

	isActive := true
	question := &models.Question{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		Level:       req.Level,
		Category:    req.Category,
		IsActive:    isActive,
		CreatedBy:   createdBy,
	}

	if err := s.Repository.Create(question); err != nil {
		return nil, err
	}

	return question, nil
}

func (s *QuestionService) Update(id string, req requests.UpdateQuestionRequest, userID string, role string) (*models.Question, error) {
	if id == "" {
		return nil, errors.New("question ID is required")
	}
	if userID == "" {
		return nil, errors.New("unauthorized")
	}

	var question *models.Question
	var err error

	if role == constants.RoleAdmin {
		question, err = s.Repository.FindByID(id)
	} else {
		question, err = s.Repository.FindOwnedByID(id, userID)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if role == constants.RoleAdmin {
				return nil, errors.New("question not found")
			}
			return nil, errors.New("question not found or you are not the owner")
		}
		return nil, err
	}

	updates := make(map[string]interface{})

	if req.Title != nil {
		updates["title"] = strings.TrimSpace(*req.Title)
	}
	if req.Description != nil {
		updates["description"] = strings.TrimSpace(*req.Description)
	}
	if req.Type != nil {
		updates["type"] = *req.Type
	}
	if req.Level != nil {
		updates["level"] = *req.Level
	}
	if req.Category != nil {
		updates["category"] = strings.TrimSpace(*req.Category)
	}

	if len(updates) == 0 {
		return question, errors.New("no fields to update") // No updates, return existing question
	}

	if err := s.Repository.Update(id, updates); err != nil {
		return nil, err
	}

	// Fetch the updated question
	updatedQuestion, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	s.AuditService.Log(userID, "question.updated", "question", id, "Question updated by user")

	return updatedQuestion, nil

}

func (s *QuestionService) Delete(id string, userID string, role string) error {
	if id == "" {
		return errors.New("question ID is required")
	}
	if userID == "" {
		return errors.New("unauthorized")
	}

	if role == constants.RoleAdmin {
		_, err := s.Repository.FindByID(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("question not found")
			}
			return err
		}
		_ = s.AuditService.Log(userID, "question.deleted", "question", id, "Question deleted by admin")
		return s.Repository.Delete(id)
	}

	_, err := s.Repository.FindOwnedByID(id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("question not found or you are not the owner")
		}
		return err
	}

	_ = s.AuditService.Log(userID, "question.deleted", "question", id, "Question deleted by owner")

	return s.Repository.Delete(id)
}
