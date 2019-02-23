package domain

type (
	// Summary represents the information about how much of the budget entered you spent
	// and the time spent in activities in minutes
	Summary struct {
		BudgetSpent Budget `json:"budget_spent"`
		TimeSpent   int    `json:"time_spent"`
	}

	// Schedule represents the personalize itinerary with summary information
	// and the details of which activities to do which day including the start time of the activity
	Schedule struct {
		Summary Summary        `json:"summary"`
		Days    []ItineraryDay `json:"days"`
	}
)
