package bootstrap

import (
	"fmt"
	"log"

	"hiretest-api/configs"
)

func Run() {
	cfg := configs.LoadAppConfig()
	db := NewDatabase(cfg)
	app := NewFiberApp(cfg)
	RegisterRoutes(app, db, cfg)
	addr := fmt.Sprintf("%s:%s", cfg.Env.AppHost, cfg.Env.AppPort)
	log.Fatal(app.Listen(addr))
}
