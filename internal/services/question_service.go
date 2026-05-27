package services

import (
	"errors"
	"hiretest-api/internal/models"
	"hiretest-api/internal/repositories"
	"hiretest-api/internal/requests"
)

type QuestionService struct {
	Repository *repositories.QuestionRepository
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
}

func NewQuestionService(repo *repositories.QuestionRepository) *QuestionService {
	return &QuestionService{
		Repository: repo,
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
