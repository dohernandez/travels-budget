package platform

import "github.com/pkg/errors"

// ErrInsufficientBudgetForDay is a common insufficient budget per day
var ErrInsufficientBudgetForDay = errors.New("insufficient budget. Budget must be at least 50 per day")

// ErrInvalidInputFile is a common invalid input file error
var ErrInvalidInputFile = errors.New("invalid input file")

// ErrInvalidInputDay is a common invalid input day error
var ErrInvalidInputDay = errors.New("invalid input day")

// ErrInvalidInputBudget is a common invalid input budget error
var ErrInvalidInputBudget = errors.New("invalid input budget")
