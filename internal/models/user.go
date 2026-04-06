package models

import "time"

type User struct {
	ID        string    	gorm:\"type:uuid;primaryKey\" json:\"id\"	CreatedAt time.Time 	json:\"created_at\"	UpdatedAt time.Time 	json:\"updated_at\"}
