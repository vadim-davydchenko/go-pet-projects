package main

import (
	"go-pet-projects/fiber/config"
	"go-pet-projects/fiber/internal/home"
	"go-pet-projects/fiber/internal/vacancy"
	"go-pet-projects/fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("public", "./public")

	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger)

	app.Listen(":3000")
}
