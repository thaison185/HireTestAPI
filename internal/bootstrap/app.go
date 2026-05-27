package bootstrap

import (
	"fmt"
	"log"

	"hiretest-api/configs"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, using system environment variables")
	}
	cfg := configs.LoadAppConfig()
	db := NewDatabase(cfg)
	if err := RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	if cfg.Env.RunSeeder {
		if err := RunSeeders(db); err != nil {
			log.Fatalf("failed to run seeders: %v", err)
		}
	}

	app := NewFiberApp(cfg)
	RegisterRoutes(app, db, cfg)
	addr := fmt.Sprintf("%s:%s", cfg.Env.AppHost, cfg.Env.AppPort)
	log.Fatal(app.Listen(addr))
}
