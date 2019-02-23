package input_test

import (
	"context"
	"testing"

	"github.com/dohernandez/travels-budget/internal/platform"
	"github.com/dohernandez/travels-budget/internal/platform/validator/input"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDayValidator(t *testing.T) {
	testCases := []struct {
		scenario string
		days     int
		result   error
	}{
		{
			scenario: "Days are valid, between 1 and 5",
			days:     2,
		},
		{
			scenario: "Days are invalid, is lower than 1",
			result:   errors.Wrap(platform.ErrInvalidInputDay, "value must be between 1 and 5"),
		},
		{
			scenario: "Days are invalid, is grater than 5",
			days:     6,
			result:   errors.Wrap(platform.ErrInvalidInputDay, "value must be between 1 and 5"),
		},
	}

	v := input.NewDaysValidator()

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			valid, err := v.IsValid(context.TODO(), tc.days)
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
