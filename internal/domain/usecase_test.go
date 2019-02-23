package domain_test

import (
	"context"
	"testing"

	"time"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestRandomPersonalizeScheduleUseCase(t *testing.T) {
	activities := domain.NewActivitiesMock()

	testCases := []struct {
		scenario           string
		activityFinderFunc func() ([]domain.Activity, error)
		budget             domain.Budget
		days               domain.Day
		function           func(*testing.T, domain.ActivityFinder, domain.Budget, domain.Day)
	}{
		{
			scenario: "Suggested at least 3 activities per day",
			activityFinderFunc: func() ([]domain.Activity, error) {
				return activities, nil
			},
			budget:   domain.Budget(800),
			days:     domain.Day(2),
			function: testShouldSuggestAtLeastThreeActivitiesPerDay,
		},
		{
			scenario: "Error insufficient budget. Itinerary day less than 3",
			activityFinderFunc: func() ([]domain.Activity, error) {
				var i int
				nActivities := make([]domain.Activity, len(activities)-2)

				for _, activity := range activities {
					if activity.ID == 6768 || activity.ID == 6558 {
						continue
					}

					nActivities[i] = activity
					i++
				}

				return nActivities, nil
			},
			budget:   domain.Budget(50),
			days:     domain.Day(1),
			function: testShouldFailInsufficientBudget,
		},
		{
			scenario: "Error budget was not consumed at 60%",
			activityFinderFunc: func() ([]domain.Activity, error) {
				return activities, nil
			},
			budget:   domain.Budget(50),
			days:     domain.Day(1),
			function: testShouldFailBudgetNotConsume,
		},
		{
			scenario: "Error budget, activityFinder.FindAll fails",
			activityFinderFunc: func() ([]domain.Activity, error) {
				return nil, errors.New("error activityFinder.FindAll")
			},
			budget:   domain.Budget(50),
			days:     domain.Day(1),
			function: testShouldFailActivityFinderFindAll,
		},
		{
			scenario: "Suggested 3 activities, but not 4 per day",
			activityFinderFunc: func() ([]domain.Activity, error) {
				return []domain.Activity{
					{
						ID:       7495,
						Duration: 180 * time.Minute,
						Price:    275,
					},
					{
						ID:       6588,
						Duration: 240 * time.Minute,
						Price:    430,
					},
					{
						ID:       7046,
						Duration: 180 * time.Minute,
						Price:    312,
					},
					{
						ID:       7026,
						Duration: 120 * time.Minute,
						Price:    268,
					},
				}, nil
			},
			budget:   domain.Budget(1400),
			days:     domain.Day(1),
			function: testShouldSuggestedThreeButNotFour,
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			finder := domain.NewCallbackActivityFinderMock(tc.activityFinderFunc)
			tc.function(t, finder, tc.budget, tc.days)
		})
	}
}

// testShouldSuggestAtLeastThreeActivitiesPerDay tests scenario where travel suggested fulfill all requirements
func testShouldSuggestAtLeastThreeActivitiesPerDay(
	t *testing.T,
	finder domain.ActivityFinder,
	budget domain.Budget,
	days domain.Day,
) {
	uc := domain.NewRandomPersonalizeScheduleUseCase(finder)

	suggested, err := uc.Do(context.TODO(), budget, days)
	assert.NoError(t, err, "error was not expected, given %s", err)
	if assert.Len(t, suggested.Days, 2, "suggested days was not expected, expected 2 given %s", len(suggested.Days)) {
		assert.True(t, len(suggested.Days[0].Itineraries) >= 3, "suggested first day itineraries were not expected, expected 3 or more, given %s", len(suggested.Days[0].Itineraries))
		assert.True(t, len(suggested.Days[1].Itineraries) >= 3, "suggested second day itineraries were not expected, expected 3 or more, given %s", len(suggested.Days[1].Itineraries))
	}
}

// testShouldFailInsufficientBudget tests scenario where budget is insufficient to suggest at least 3 activities in one day
func testShouldFailInsufficientBudget(
	t *testing.T,
	finder domain.ActivityFinder,
	budget domain.Budget,
	days domain.Day,
) {
	uc := domain.NewRandomPersonalizeScheduleUseCase(finder)

	_, err := uc.Do(context.TODO(), budget, days)
	assert.Error(t, err, "error was expected")
	assert.EqualError(t, err, domain.ErrInsufficientBudgetForItinerary.Error())
}

// testShouldFailInsufficientBudget tests scenario where 60% of budget should be consume is not fulfill
func testShouldFailBudgetNotConsume(
	t *testing.T,
	finder domain.ActivityFinder,
	budget domain.Budget,
	days domain.Day,
) {
	uc := domain.NewRandomPersonalizeScheduleUseCase(finder)

	_, err := uc.Do(context.TODO(), budget, days)
	assert.Error(t, err, "error was expected")
	assert.EqualError(t, err, domain.ErrBudgetNotConsumed.Error())
}

// testShouldFailActivityFinderFindAll tests scenario where fail to load activities
func testShouldFailActivityFinderFindAll(
	t *testing.T,
	finder domain.ActivityFinder,
	budget domain.Budget,
	days domain.Day,
) {
	uc := domain.NewRandomPersonalizeScheduleUseCase(finder)

	_, err := uc.Do(context.TODO(), budget, days)
	assert.Error(t, err, "error was expected")
	assert.EqualError(t, err, "error activityFinder.FindAll")
}

// testShouldSuggestedThreeButNotFour tests scenario
// a day has 12 hours. The first activity needs to start at 10:00 and the last needs to finish before 22:00 (inclusive)
func testShouldSuggestedThreeButNotFour(
	t *testing.T,
	finder domain.ActivityFinder,
	budget domain.Budget,
	days domain.Day,
) {
	uc := domain.NewRandomPersonalizeScheduleUseCase(finder)

	s, err := uc.Do(context.TODO(), budget, days)
	assert.NoError(t, err, "error was not expected, given %s", err)
	if assert.Len(t, s.Days, 1) {
		assert.True(t, len(s.Days[0].Activities()) == 3, "suggested day itinerary activities were not expected, expected 3, given %s", len(s.Days[0].Activities()))
	}
}
