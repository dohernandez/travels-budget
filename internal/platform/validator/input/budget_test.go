package input_test

import (
	"testing"

	"context"

	"github.com/dohernandez/travels-budget/internal/platform"
	"github.com/dohernandez/travels-budget/internal/platform/validator/input"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBudgetValidator(t *testing.T) {
	testCases := []struct {
		scenario string
		budget   int
		days     int
		result   error
	}{
		{
			scenario: "Budget is valid, between 100 and 2000",
			budget:   150,
			days:     1,
		},
		{
			scenario: "Budget is invalid, is lower than 100",
			budget:   50,
			days:     1,
			result:   errors.Wrap(platform.ErrInvalidInputBudget, "value must be between 100 and 2000"),
		},
		{
			scenario: "Budget is invalid, is grater than 2000",
			budget:   2400,
			days:     1,
			result:   errors.Wrap(platform.ErrInvalidInputBudget, "value must be between 100 and 2000"),
		},
		{
			scenario: "Budget is invalid, is lower than 50 per day",
			budget:   100,
			days:     5,
			result:   errors.Wrap(platform.ErrInvalidInputBudget, "value allocated per day needs to be 50"),
		},
	}

	v := input.NewBudgetValidator()

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			valid, err := v.IsValid(context.TODO(), tc.budget, tc.days)
			if tc.result != nil {
				assert.False(t, valid, "false was expected")
				assert.EqualError(t, err, tc.result.Error(), "error was expected")
			} else {
				assert.True(t, valid, "true was expected")
				assert.NoError(t, err, "error was not expected, given %s", err)
			}
		})
	}
}
