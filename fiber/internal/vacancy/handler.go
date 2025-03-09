package vacancy

import (
	"go-pet-projects/fiber/pkg/tadapter"
	"go-pet-projects/fiber/pkg/validator"
	"go-pet-projects/fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email not present"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatError(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
