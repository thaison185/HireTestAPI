package services

import (
	"errors"
	code_errors "hiretest-api/internal/common/errors"
	"hiretest-api/internal/models"
	"hiretest-api/internal/repositories"
)

type ProfileService struct {
	Repo *repositories.AuthRepository
}

func NewProfileService(repo *repositories.AuthRepository) *ProfileService {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) GetCurrentProfile(userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New(code_errors.CodeUnauthorized)
	}
	user, err := s.Repo.FindUserByID(userID)
	if err != nil {
		return nil, errors.New(code_errors.CodeUserNotFound)
	}

	if user.IsActive == false {
		return nil, errors.New(code_errors.CodeAccountInactive)
	}

	return user, nil
}
