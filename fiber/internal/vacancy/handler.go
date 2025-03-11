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
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Salary:   c.FormValue("salary"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email not present"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Location not present"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Field company not present"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Post not present"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Name Company not present"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Salary not present"},
	)

	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatError(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
