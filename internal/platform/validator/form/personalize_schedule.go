package form

import (
	"context"

	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
	"github.com/pkg/errors"
)

// ErrInvalidPersonalizeScheduleForm is a common invalid input budget error
var ErrInvalidPersonalizeScheduleForm = errors.New("invalid input personalize schedule form")

type (
	personalizeScheduleValidator struct {
		daysValidator   validator.InputDaysValidator
		budgetValidator validator.InputBudgetValidator
	}
)

// NewPersonalizeScheduleValidator creates an instance of validator.FormValidator for intput.PersonalizeScheduleForm
func NewPersonalizeScheduleValidator(
	daysValidator validator.InputDaysValidator,
	budgetValidator validator.InputBudgetValidator,
) validator.FormValidator {
	return &personalizeScheduleValidator{
		daysValidator:   daysValidator,
		budgetValidator: budgetValidator,
	}
}

// Validate validates intput.PersonalizeScheduleForm
// Constraints:
// 		* budget is an integer between 100 and 2000.
// 		* The average budget to be allocated per day needs to be 50.
// 			Which means for budget 100 you can only have 1 or 2 days.
// 			Which also means the minimum budget for 5 days needs to be 250.
//		* days is an integer between 1 and 5.
//
// Returns an error if the intput.PersonalizeScheduleForm is invalid.
func (v *personalizeScheduleValidator) Validate(ctx context.Context, form interface{}) error {
	f, ok := form.(intput.PersonalizeScheduleForm)
	if !ok {
		return ErrInvalidPersonalizeScheduleForm
	}

	_, err := v.daysValidator.IsValid(ctx, f.Days)
	if err != nil {
		return err
	}

	_, err = v.budgetValidator.IsValid(ctx, f.Budget, f.Days)
	if err != nil {
		return err
	}

	return nil
}
