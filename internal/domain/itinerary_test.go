package domain_test

import (
	"encoding/json"
	"regexp"
	"testing"
	"time"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestTimeSpentItinerary(t *testing.T) {
	itineraryDay := domain.NewItineraryDayMock()

	assert.Equal(t, itineraryDay.TimeSpent(), 360)
}

func TestBudgetSpentItinerary(t *testing.T) {
	itineraryDay := domain.NewItineraryDayMock()

	assert.Equal(t, domain.Budget(734), itineraryDay.BudgetSpent())
}

func TestActivitiesItinerary(t *testing.T) {
	itineraryDay := domain.NewItineraryDayMock()
	activities := []domain.Activity{
		{
			ID:       7495,
			Duration: 120 * time.Minute,
			Price:    423,
		},
		{
			ID:       6748,
			Duration: 120 * time.Minute,
			Price:    15,
		},
		{
			ID:       6588,
			Duration: 90 * time.Minute,
			Price:    81,
		},
		{
			ID:       7046,
			Duration: 30 * time.Minute,
			Price:    215,
		},
	}

	assert.Equal(t, activities, itineraryDay.Activities())
}

func TestEndingTimeItinerary(t *testing.T) {
	itineraryDay := domain.NewItineraryDayMock()

	now := time.Now()
	nyear, nmonth, nday := now.Date()
	td := time.Date(nyear, nmonth, nday, 0, 0, 0, 0, now.Location())

	endingTime := td.
		// the start time of the activities in the itinerary
		Add(time.Hour * 10).
		// duration activity
		Add(120 * time.Minute).
		// the interval time between activities in the itinerary
		Add(time.Minute * 30).
		// duration activity
		Add(120 * time.Minute).
		// the interval time between activities in the itinerary
		Add(time.Minute * 30).
		// duration activity
		Add(90 * time.Minute).
		// the interval time between activities in the itinerary
		Add(time.Minute * 30).
		// duration activity
		Add(30 * time.Minute)

	assert.Equal(t, endingTime, itineraryDay.EndingTime())
}

func TestMarshalItinerary(t *testing.T) {
	itineraryDay := domain.NewItineraryDayMock()

	// Reduce to 1 the amount of itinerary, there is no need to run the test with the whole collection
	// of itineraries with 1 itinerary is enough
	itineraryDay = domain.ItineraryDay{
		Day:         itineraryDay.Day,
		Itineraries: itineraryDay.Itineraries[0:1],
	}

	itineraryDayJson, err := json.Marshal(itineraryDay)
	assert.NoError(t, err, "error was not expected, given %s")
	assert.Equal(t, regexp.MustCompile(`\s+`).ReplaceAllString(`{
		"day": 1,
		"itinerary": [{
            "start": "10:00",
			"activity": {
                "id": 7495,
                "duration": 120,
				"price": 423
			}
		}]
	}`, ""), string(itineraryDayJson))
}
