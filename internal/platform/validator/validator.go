package validator

import "context"

// InputBudgetValidator validates budget input value
type InputBudgetValidator interface {
	IsValid(ctx context.Context, budget, days int) (bool, error)
}

// InputDaysValidator validates budget input value
type InputDaysValidator interface {
	IsValid(ctx context.Context, days int) (bool, error)
}

// FormValidator validates form (all the inputs)
type FormValidator interface {
	Validate(ctx context.Context, form interface{}) error
}
