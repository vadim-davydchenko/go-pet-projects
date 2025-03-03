package home

import (
	"go-pet-projects/fiber/pkg/tadapter"
	"go-pet-projects/fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Hello("John")

	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
