package models

import "time"

type AuditLog struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	ActorID     string    `gorm:"type:uuid;not null" json:"actor_id"`
	Action      string    `gorm:"type:varchar(100);not null" json:"action"`
	EntityType  string    `gorm:"type:varchar(100);not null" json:"entity_type"`
	EntityID    string    `gorm:"type:uuid;not null" json:"entity_id"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
