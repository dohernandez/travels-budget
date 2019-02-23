package bootstrap

import (
	"os"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/dohernandez/travels-budget/internal/platform"
	"github.com/dohernandez/travels-budget/internal/platform/config"
	"github.com/dohernandez/travels-budget/internal/platform/container"
	"github.com/dohernandez/travels-budget/internal/platform/output"
	"github.com/dohernandez/travels-budget/internal/platform/storage"
	"github.com/dohernandez/travels-budget/internal/platform/validator/form"
	"github.com/dohernandez/travels-budget/internal/platform/validator/input"
	"github.com/pkg/errors"
)

// NewContainer creates an instance of container.Container with the application resources
func NewContainer(conf *config.Specification) *container.Container {
	c := container.Container{}

	// Prepare activities source
	activitiesSource, err := os.Open(conf.ActivitiesFile)
	if err != nil {
		panic(errors.Wrap(platform.ErrInvalidInputFile, err.Error()))
	}

	// FINDERS
	activityFinder := storage.NewFileActivityFinder(activitiesSource)

	// USE CASES
	c.WithRandomPersonalizeScheduleUseCase(domain.NewRandomPersonalizeScheduleUseCase(activityFinder))

	// INPUT VALIDATORS
	daysValidator := input.NewDaysValidator()
	budgetValidator := input.NewBudgetValidator()

	// FORM VALIDATORS
	c.WithPersonalizeScheduleValidator(
		form.NewPersonalizeScheduleValidator(daysValidator, budgetValidator),
	)

	// OUTPUT
	switch conf.Output {
	default:
		c.WithRenderer(output.NewStandardOutput())
	}

	return &c
}
