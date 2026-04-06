package bootstrap

import (
	"log"

	"hiretest-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg configs.AppConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.DSN(cfg.Env)), &gorm.Config{})
	if err != nil {
		log.Printf("database connection failed: %v", err)
		return nil
	}
	return db
}
