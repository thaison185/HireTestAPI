package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"hiretest-api/configs"
)

func NewFiberApp(cfg configs.AppConfig) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())
	return app
}
