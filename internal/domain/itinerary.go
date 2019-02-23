package domain

import (
	"encoding/json"
	"time"
)

// Itinerary represents a single itinerary aggregator
type Itinerary struct {
	// Start start time of the activity
	Start    time.Time
	Activity Activity
}

// MarshalJSON implements the json.Marshaler interface.
// to enforces Start value
func (i Itinerary) MarshalJSON() ([]byte, error) {
	itinerary := struct {
		Start    string   `json:"start"`
		Activity Activity `json:"activity"`
	}{
		Start:    i.Start.Format("15:04"),
		Activity: i.Activity,
	}

	return json.Marshal(itinerary)
}

// ItineraryDay represents a day itinerary aggregator
type ItineraryDay struct {
	Day         Day         `json:"day"`
	Itineraries []Itinerary `json:"itinerary"`
}

// TimeSpent calculate the time spent in the day itinerary
func (id ItineraryDay) TimeSpent() int {
	var timeSpent int

	for _, itinerary := range id.Itineraries {
		timeSpent += int(itinerary.Activity.Duration.Minutes())
	}

	return timeSpent
}

// BudgetSpent calculate the budget spent in the day itinerary
func (id ItineraryDay) BudgetSpent() Budget {
	var budgetSpent int

	for _, itinerary := range id.Itineraries {
		budgetSpent += itinerary.Activity.Price
	}

	return Budget(budgetSpent)
}

// Activities returns all the activities for the day itinerary
func (id ItineraryDay) Activities() []Activity {
	activities := make([]Activity, len(id.Itineraries))

	for k, itinerary := range id.Itineraries {
		activity := itinerary.Activity

		activities[k] = activity
	}

	return activities
}

// EndingTime returns when the last activity ends
func (id ItineraryDay) EndingTime() time.Time {
	itinerary := id.Itineraries[len(id.Itineraries)-1]

	return itinerary.Start.Add(itinerary.Activity.Duration)
}
