package container

import (
	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/dohernandez/travels-budget/internal/platform/output"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
)

// Container contains application resources
type Container struct {
	personalizeScheduleValidator validator.FormValidator

	randomPersonalizeScheduleUseCase domain.RandomPersonalizeScheduleUseCase

	renderer output.Renderer
}

// WithRandomPersonalizeScheduleUseCase sets a global instance of domain.RandomPersonalizeScheduleUseCase
func (c *Container) WithRandomPersonalizeScheduleUseCase(personalizeScheduleUseCase domain.RandomPersonalizeScheduleUseCase) *Container {
	c.randomPersonalizeScheduleUseCase = personalizeScheduleUseCase

	return c
}

// RandomPersonalizeScheduleUseCase get the global instance of domain.RandomPersonalizeScheduleUseCase
func (c *Container) RandomPersonalizeScheduleUseCase() domain.RandomPersonalizeScheduleUseCase {
	return c.randomPersonalizeScheduleUseCase
}

// WithPersonalizeScheduleValidator sets a global instance of validator.FormValidator
func (c *Container) WithPersonalizeScheduleValidator(personalizeScheduleValidator validator.FormValidator) *Container {
	c.personalizeScheduleValidator = personalizeScheduleValidator

	return c
}

// PersonalizeScheduleValidator get the global instance of validator.FormValidator
func (c *Container) PersonalizeScheduleValidator() validator.FormValidator {
	return c.personalizeScheduleValidator
}

// WithRenderer sets a global instance of output.renderer
func (c *Container) WithRenderer(renderer output.Renderer) *Container {
	c.renderer = renderer

	return c
}

// Renderer get the global instance of output.renderer
func (c *Container) Renderer() output.Renderer {
	return c.renderer
}
