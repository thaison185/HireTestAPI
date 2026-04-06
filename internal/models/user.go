package models

import "time"

type User struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
	FullName string `gorm:"not null" json:"full_name"`
	Role string `json:"role"`
	IsActive bool `gorm:"default:true" json:"is_active"`
}
