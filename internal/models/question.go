package models

import "time"

type Question struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Type        string    `gorm:"type:varchar(50);not null" json:"type"`
	Level       string    `gorm:"type:varchar(50);not null" json:"level"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedBy   string    `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
