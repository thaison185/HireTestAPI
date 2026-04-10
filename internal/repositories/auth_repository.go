package repositories

import "hiretest-api/internal/models"

type AuthRepository struct {
	BaseRepository
}

func NewAuthRepository(base BaseRepository) *AuthRepository {
	return &AuthRepository{BaseRepository: base}
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) FindUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) SaveRefreshToken(userID string, refreshToken string) error {
	// Implement logic to save refresh token in the database
	return nil
}

func (r *AuthRepository) FindRefreshToken(userID string) (string, error) {
	// Implement logic to find refresh token by user ID in the database
	return "", nil
}

func (r *AuthRepository) DeleteRefreshToken(userID string) error {
	// Implement logic to delete refresh token from the database
	return nil
}
