package bootstrap

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/models"

	"errors"

	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) error {
	adminPasswordHash, err := utils.HashPassword("Admin@123") // bcrypt hash for "admin123"
	if err != nil {
		return err
	}

	reviewerPasswordHash, err := utils.HashPassword("Reviewer@123") // bcrypt hash for "reviewer123"
	if err != nil {
		return err
	}

	candidatePasswordHash, err := utils.HashPassword("Candidate@123") // bcrypt hash for "candidate123"
	if err != nil {
		return err
	}

	seedUsers := []models.User{
		{
			Email:        "admin@hiretest.com",
			PasswordHash: adminPasswordHash,
			FullName:     "Admin User",
			Role:         "admin",
			IsActive:     true,
		},
		{
			Email:        "reviewer@hiretest.com",
			PasswordHash: reviewerPasswordHash,
			FullName:     "Reviewer User",
			Role:         "reviewer",
			IsActive:     true,
		},
		{
			Email:        "candidate@hiretest.com",
			PasswordHash: candidatePasswordHash,
			FullName:     "Candidate User",
			Role:         "candidate",
			IsActive:     true,
		},
	}

	for _, user := range seedUsers {
		var existingUser models.User
		err := db.Where("email = ?", user.Email).First(&existingUser).Error
		if err == nil {
			update := map[string]interface{}{
				"full_name":     user.FullName,
				"password_hash": user.PasswordHash,
				"role":          user.Role,
				"is_active":     user.IsActive,
			}
			if err := db.Model(&existingUser).Updates(update).Error; err != nil {
				return err
			}
			continue
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			continue
		}

		return err
	}

	return nil
}
