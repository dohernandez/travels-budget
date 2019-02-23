package domain

import (
	"context"
	"time"
)

// NewCallbackActivityFinderMock creates a callback mock for tests
// nolint:unused
func NewCallbackActivityFinderMock(finderFunc func() ([]Activity, error)) ActivityFinder {
	return &activityFinderMock{
		finderFunc: finderFunc,
	}
}

// nolint:unused
type activityFinderMock struct {
	finderFunc func() ([]Activity, error)
}

func (m *activityFinderMock) FindAll() ([]Activity, error) {
	return m.finderFunc()
}

// NewActivitiesMock returns a collection of Activity mocked for test purpose
// nolint:unused
func NewActivitiesMock() []Activity {
	return []Activity{
		{
			ID:       7495,
			Duration: 120 * time.Minute,
			Price:    423,
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
		{
			ID:       6748,
			Duration: 120 * time.Minute,
			Price:    15,
		},
		{
			ID:       6779,
			Duration: 90 * time.Minute,
			Price:    49,
		},
		{
			ID:       6659,
			Duration: 150 * time.Minute,
			Price:    39,
		},
		{
			ID:       6647,
			Duration: 30 * time.Minute,
			Price:    59,
		},
		{
			ID:       6741,
			Duration: 120 * time.Minute,
			Price:    94,
		},
		{
			ID:       6768,
			Duration: 30 * time.Minute,
			Price:    8,
		},
		{
			ID:       6558,
			Duration: 30 * time.Minute,
			Price:    5,
		},
	}
}

// NewItineraryDayMock returns an ItineraryDay mocked for test purpose
// nolint:unused
func NewItineraryDayMock() ItineraryDay {
	now := time.Now()
	nyear, nmonth, nday := now.Date()
	t := time.Date(nyear, nmonth, nday, 0, 0, 0, 0, now.Location())

	durationFirstActivity := 120 * time.Minute
	durationSecondActivity := 120 * time.Minute
	durationThirdActivity := 90 * time.Minute
	durationFourthActivity := 30 * time.Minute

	startFirstActivity := t.Add(dayStartTimeActivity)
	startSecondActivity := startFirstActivity.Add(durationFirstActivity).Add(dayActivityInterval)
	startThirdActivity := startSecondActivity.Add(durationSecondActivity).Add(dayActivityInterval)
	startFourthActivity := startThirdActivity.Add(durationThirdActivity).Add(dayActivityInterval)

	return ItineraryDay{
		Day: Day(1),
		Itineraries: []Itinerary{
			{
				Start: startFirstActivity,
				Activity: Activity{
					ID:       7495,
					Duration: durationFirstActivity,
					Price:    423,
				},
			},
			{
				Start: startSecondActivity,
				Activity: Activity{
					ID:       6748,
					Duration: durationSecondActivity,
					Price:    15,
				},
			},
			{
				Start: startThirdActivity,
				Activity: Activity{
					ID:       6588,
					Duration: durationThirdActivity,
					Price:    81,
				},
			},
			{
				Start: startFourthActivity,
				Activity: Activity{
					ID:       7046,
					Duration: durationFourthActivity,
					Price:    215,
				},
			},
		},
	}
}

// NewCallbackRandomPersonalizeScheduleUseCaseMock creates a callback mock for tests
// nolint:unused
func NewCallbackRandomPersonalizeScheduleUseCaseMock(useCaseFunc func(budget Budget, days Day) (Schedule, error)) RandomPersonalizeScheduleUseCase {
	return &randomPersonalizeScheduleMock{
		useCaseFunc: useCaseFunc,
	}
}

// nolint:unused
type randomPersonalizeScheduleMock struct {
	useCaseFunc func(budget Budget, days Day) (Schedule, error)
}

func (m *randomPersonalizeScheduleMock) Do(_ context.Context, budget Budget, days Day) (Schedule, error) {
	return m.useCaseFunc(budget, days)
}
