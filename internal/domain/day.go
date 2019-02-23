package domain

import "time"

const (
	// dayStartTimeActivity defines the start time of the activities in the itinerary
	dayStartTimeActivity = time.Hour * 10
	// dayEndTimeActivity defines the end time of the activities in the itinerary
	dayEndTimeActivity = time.Hour * 22
	// dayActivityInterval defines the interval time between activities in the itinerary
	dayActivityInterval = time.Minute * 30
	// dayMinimalActivities defines the minimal activities in the itinerary per day
	dayMinimalActivities = 3
)

// Day represents the day value
type Day int
