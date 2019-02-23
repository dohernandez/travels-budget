package form_test

import (
	"testing"

	"context"

	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/dohernandez/travels-budget/internal/platform/validator/form"
	"github.com/dohernandez/travels-budget/internal/platform/validator/input"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeScheduleValidator(t *testing.T) {
	budget := 680
	days := 2

	testCases := []struct {
		scenario            string
		dayValidatorFunc    func(days int) (bool, error)
		budgetValidatorFunc func(budget, days int) (bool, error)
		form                interface{}
		result              error
	}{
		{
			scenario: "Valid budget, days and file inputs",
			dayValidatorFunc: func(d int) (bool, error) {
				assert.Equal(t, days, d)

				return true, nil
			},
			budgetValidatorFunc: func(b, d int) (bool, error) {
				assert.Equal(t, budget, b)
				assert.Equal(t, days, d)

				return true, nil
			},
			form: intput.PersonalizeScheduleForm{
				Days:   days,
				Budget: budget,
			},
		},

		{
			scenario: "Invalid days input",
			dayValidatorFunc: func(d int) (bool, error) {
				assert.Equal(t, days, d)

				return false, errors.New("Invalid days input")
			},
			budgetValidatorFunc: func(b, d int) (bool, error) {
				panic("should not be called")
			},
			form: intput.PersonalizeScheduleForm{
				Days:   days,
				Budget: budget,
			},
			result: errors.New("Invalid days input"),
		},

		{
			scenario: "Invalid budget input",
			dayValidatorFunc: func(d int) (bool, error) {
				assert.Equal(t, days, d)

				return true, nil
			},
			budgetValidatorFunc: func(b, d int) (bool, error) {
				assert.Equal(t, budget, b)
				assert.Equal(t, days, d)

				return false, errors.New("Invalid budget input")
			},
			form: intput.PersonalizeScheduleForm{
				Days:   days,
				Budget: budget,
			},
			result: errors.New("Invalid budget input"),
		},

		{
			scenario: "Invalid input personalize schedule form",
			dayValidatorFunc: func(d int) (bool, error) {
				panic("should not be called")
			},
			budgetValidatorFunc: func(b, d int) (bool, error) {
				panic("should not be called")
			},
			form:   "invalid form",
			result: form.ErrInvalidPersonalizeScheduleForm,
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			dayValidator := input.NewCallbackDaysInputValidatorMock(tc.dayValidatorFunc)
			budgetValidator := input.NewCallbackBudgetInputValidatorMock(tc.budgetValidatorFunc)

			err := form.NewPersonalizeScheduleValidator(dayValidator, budgetValidator).Validate(context.TODO(), tc.form)
			if tc.result != nil {
				assert.EqualError(t, err, tc.result.Error(), "error was expected")
			} else {
				assert.NoError(t, err, "error was not expected, given %s", err)
			}
		})
	}
}
