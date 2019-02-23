package cli_test

import (
	"context"
	"testing"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/dohernandez/travels-budget/internal/platform/container"
	cliContext "github.com/dohernandez/travels-budget/internal/platform/context"
	handler "github.com/dohernandez/travels-budget/internal/platform/handler/cli"
	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/dohernandez/travels-budget/internal/platform/output"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeScheduleHandler(t *testing.T) {
	budget := 680
	days := 2

	schedule := domain.Schedule{
		Summary: domain.Summary{
			BudgetSpent: 530,
		},
	}

	testCases := []struct {
		scenario        string
		validatorFunc   func(form interface{}) error
		useCaseFunc     func(b domain.Budget, d domain.Day) (domain.Schedule, error)
		rendererOutFunc func(output interface{}) error
		rendererErrFunc func(err error) error
	}{
		{
			scenario: "Print schedule to output successfully",
			validatorFunc: func(form interface{}) error {
				f, ok := form.(intput.PersonalizeScheduleForm)
				assert.True(t, ok)

				assert.Equal(t, budget, f.Budget)
				assert.Equal(t, days, f.Days)

				return nil
			},
			useCaseFunc: func(b domain.Budget, d domain.Day) (domain.Schedule, error) {
				assert.Equal(t, domain.Budget(budget), b)
				assert.Equal(t, domain.Day(days), d)

				return schedule, nil
			},
			rendererOutFunc: func(output interface{}) error {
				out := struct {
					Schedule domain.Schedule `json:"schedule"`
				}{
					Schedule: schedule,
				}

				assert.Equal(t, out, output)

				return nil
			},
			rendererErrFunc: func(err error) error {
				panic("should not be called")
			},
		},
		{
			scenario: "Fails schedule due to validation fails",
			validatorFunc: func(form interface{}) error {
				f, ok := form.(intput.PersonalizeScheduleForm)
				assert.True(t, ok)

				assert.Equal(t, budget, f.Budget)
				assert.Equal(t, days, f.Days)

				return errors.New("validation fails")
			},
			useCaseFunc: func(b domain.Budget, d domain.Day) (domain.Schedule, error) {
				panic("should not be called")
			},
			rendererOutFunc: func(output interface{}) error {
				panic("should not be called")
			},
			rendererErrFunc: func(err error) error {
				assert.EqualError(t, err, "validation fails")

				return nil
			},
		},
		{
			scenario: "Fails schedule due to use case fails",
			validatorFunc: func(form interface{}) error {
				f, ok := form.(intput.PersonalizeScheduleForm)
				assert.True(t, ok)

				assert.Equal(t, budget, f.Budget)
				assert.Equal(t, days, f.Days)

				return nil
			},
			useCaseFunc: func(b domain.Budget, d domain.Day) (domain.Schedule, error) {
				assert.Equal(t, domain.Budget(budget), b)
				assert.Equal(t, domain.Day(days), d)

				return domain.Schedule{}, errors.New("use case fails")
			},
			rendererOutFunc: func(output interface{}) error {
				panic("should not be called")
			},
			rendererErrFunc: func(err error) error {
				assert.EqualError(t, err, "use case fails")

				return nil
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			c := container.Container{}

			c.WithPersonalizeScheduleValidator(validator.NewCallbackFormValidatorMock(tc.validatorFunc))
			c.WithRandomPersonalizeScheduleUseCase(domain.NewCallbackRandomPersonalizeScheduleUseCaseMock(tc.useCaseFunc))
			c.WithRenderer(output.NewCallbackRendererMock(tc.rendererOutFunc, tc.rendererErrFunc))

			cliCtx := cliContext.NewPersonalizeScheduleCliContextMock(days, budget)

			handler := handler.NewPersonalizeScheduleHandler(context.TODO(), &c)

			err := handler(cliCtx)
			assert.NoError(t, err, "error was not expected, given %s", err)
		})
	}
}
