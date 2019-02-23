package input

import (
	"context"

	"github.com/dohernandez/travels-budget/internal/platform/validator"
)

// NewCallbackDaysInputValidatorMock creates a callback mock for tests
// nolint:unused
func NewCallbackDaysInputValidatorMock(validateFunc func(days int) (bool, error)) validator.InputDaysValidator {
	return &daysValidatorMock{
		validateFunc: validateFunc,
	}
}

// nolint:unused
type daysValidatorMock struct {
	validateFunc func(days int) (bool, error)
}

func (m *daysValidatorMock) IsValid(_ context.Context, days int) (bool, error) {
	return m.validateFunc(days)
}

// NewCallbackBudgetInputValidatorMock creates a callback mock for tests
// nolint:unused
func NewCallbackBudgetInputValidatorMock(validateFunc func(budget, days int) (bool, error)) validator.InputBudgetValidator {
	return &budgetValidatorMock{
		validateFunc: validateFunc,
	}
}

// nolint:unused
type budgetValidatorMock struct {
	validateFunc func(budget, days int) (bool, error)
}

func (m *budgetValidatorMock) IsValid(_ context.Context, budget, days int) (bool, error) {
	return m.validateFunc(budget, days)
}
