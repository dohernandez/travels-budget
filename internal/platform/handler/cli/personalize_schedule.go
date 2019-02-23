package cli

import (
	"context"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/dohernandez/travels-budget/internal/platform/output"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
	"github.com/urfave/cli"
)

// NewPersonalizeScheduleHandler create a cli.ActionFunc for personalized command
//
// Handle the command itinerary-planner budget <budget> days <days>
//
func NewPersonalizeScheduleHandler(ctx context.Context, c interface {
	PersonalizeScheduleValidator() validator.FormValidator
	RandomPersonalizeScheduleUseCase() domain.RandomPersonalizeScheduleUseCase
	Renderer() output.Renderer
}) cli.ActionFunc {
	return func(cliCtx *cli.Context) error {
		personalizeScheduleForm := intput.PersonalizeScheduleFormFromCliContext(cliCtx)

		err := c.PersonalizeScheduleValidator().Validate(ctx, personalizeScheduleForm)
		if err != nil {
			return c.Renderer().Error(err)
		}

		budget := domain.Budget(personalizeScheduleForm.Budget)
		days := domain.Day(personalizeScheduleForm.Days)

		schedule, err := c.RandomPersonalizeScheduleUseCase().Do(ctx, budget, days)
		if err != nil {
			return c.Renderer().Error(err)
		}

		out := struct {
			Schedule domain.Schedule `json:"schedule"`
		}{
			Schedule: schedule,
		}

		return c.Renderer().Render(out)
	}
}
