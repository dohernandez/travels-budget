package domain

import "github.com/pkg/errors"

// ErrInsufficientBudgetForItinerary is a common insufficient budget to have the minimal activities in the itinerary error
var ErrInsufficientBudgetForItinerary = errors.New("insufficient budget for itinerary")

// ErrActivityCanNotStart is a common activity can not start after end time in the itinerary error
var ErrActivityCanNotStart = errors.New("activity can not start after 22:00")

// errActivityCanNotEnd is a common activity can not end after end time in the itinerary error
var errActivityCanNotEnd = errors.New("activity can not end after 22:00")

// ErrBudgetNotConsumed is a common budget was not consumed at 60% error
var ErrBudgetNotConsumed = errors.New("budget was not consumed at 60%")

// ErrInsufficientBudgetForDay is a common insufficient budget per day
var ErrInsufficientBudgetForDay = errors.New("insufficient budget. Budget must be at least 50 per day")
