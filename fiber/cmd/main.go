package main

import (
	"go-pet-projects/fiber/config"
	"go-pet-projects/fiber/internal/home"
	"go-pet-projects/fiber/internal/vacancy"
	"go-pet-projects/fiber/pkg/database"
	"go-pet-projects/fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	store := session.New()

	//REPOSITORIES
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	//HANDLERS
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)

	app.Listen(":3000")
}
