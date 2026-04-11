package services

import (
	"errors"
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
		return nil, errors.New("unauthorized")
	}
	user, err := s.Repo.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.IsActive == false {
		return nil, errors.New("account is inactive")
	}

	return user, nil
}
