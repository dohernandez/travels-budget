package validator

import "context"

// NewCallbackFormValidatorMock creates a callback mock for tests
// nolint:unused
func NewCallbackFormValidatorMock(formValidatorFunc func(form interface{}) error) FormValidator {
	return &formValidatorMock{
		formValidatorFunc: formValidatorFunc,
	}
}

// nolint:unused
type formValidatorMock struct {
	formValidatorFunc func(form interface{}) error
}

func (m *formValidatorMock) Validate(_ context.Context, form interface{}) error {
	return m.formValidatorFunc(form)
}
