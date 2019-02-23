package input

import (
	"context"

	"github.com/dohernandez/travels-budget/internal/platform"
	"github.com/dohernandez/travels-budget/internal/platform/validator"
	"github.com/pkg/errors"
)

const (
	// budgetMinimalPerDay defines the budget minimal requires per day
	budgetMinimalPerDay = 50
	// budgetMinimal defines the budget minimal
	budgetMinimal = 100
	// budgetMaximal defines the budget maximal
	budgetMaximal = 2000
)

type budgetInputValidator struct{}

// NewBudgetValidator creates an instance of validator.InputBudgetValidator
func NewBudgetValidator() validator.InputBudgetValidator {
	return &budgetInputValidator{}
}

// Validate validates budget
// Constraints:
// 		* budget is an integer between 100 and 2000.
// 		* The average budget to be allocated per day needs to be 50.
// 			Which means for budget 100 you can only have 1 or 2 days.
// 			Which also means the minimum budget for 5 days needs to be 250.
//
// Returns an error if the budget is invalid otherwise true.
func (v *budgetInputValidator) IsValid(_ context.Context, budget, days int) (bool, error) {
	if budget < budgetMinimal || budget > budgetMaximal {
		return false, errors.Wrap(platform.ErrInvalidInputBudget, "value must be between 100 and 2000")
	}

	if budget/days < budgetMinimalPerDay {
		return false, errors.Wrap(platform.ErrInvalidInputBudget, "value allocated per day needs to be 50")
	}

	return true, nil
}
