package input

import (
	"context"

	"github.com/dohernandez/travels-budget/internal/platform"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
	"github.com/pkg/errors"
)

const (
	// dayMinimal defines the day minimal
	dayMinimal = 1
	// dayMaximal defines the day maximal
	dayMaximal = 5
)

type daysValidator struct{}

// NewDaysValidator creates an validator.FormValidator for intput.PersonalizeScheduleForm
func NewDaysValidator() validator.InputDaysValidator {
	return &daysValidator{}
}

// Validate validates days
// Constraints:
//		* days is an integer between 1 and 5.
//
// Returns an false and error if the days is invalid otherwise true.
func (v *daysValidator) IsValid(_ context.Context, days int) (bool, error) {
	if days < dayMinimal || days > dayMaximal {
		return false, errors.Wrap(platform.ErrInvalidInputDay, "value must be between 1 and 5")
	}

	return true, nil
}
