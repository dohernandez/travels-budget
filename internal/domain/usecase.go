package domain

import (
	"context"
	"math/rand"
	"sort"
	"time"
)

// RandomPersonalizeScheduleUseCase suggests a personalize schedule
type RandomPersonalizeScheduleUseCase interface {
	Do(ctx context.Context, budget Budget, days Day) (Schedule, error)
}

type randomPersonalizeSchedule struct {
	activityFinder ActivityFinder
}

// NewRandomPersonalizeScheduleUseCase creates an instance of RandomPersonalizeScheduleUseCase
func NewRandomPersonalizeScheduleUseCase(activityFinder ActivityFinder) RandomPersonalizeScheduleUseCase {
	return &randomPersonalizeSchedule{
		activityFinder: activityFinder,
	}
}

// Do executes the use case logic.
//
// The use case algorithm is pretty simple:
//
// 		* suggests a random personalize schedule for a given budget and days
// 		* suggests at least 3 activities per given days
// 		* the first activity suggested starts at 10:00 and the last finishes before 22:00 (inclusive).
// 		* activities are separated by 30 minutes
// 		* the given budget is consumed more than 60%
//
// ErrInsufficientBudgetForDay is returned when the average budget per day is less than 50
// ErrInsufficientBudgetForItinerary is returned when the average budget per day is not enough to
// have 3 activities a day
// ErrBudgetNotConsumed is returned when the total budget spent is less than 60% of the total to spend
//
// Returns Schedule if error not happens
func (uc *randomPersonalizeSchedule) Do(_ context.Context, budget Budget, days Day) (Schedule, error) {
	budgetDay := Budget(int(budget) / int(days))

	activities, err := uc.getActivitiesOrdered(budget)
	if err != nil {
		return Schedule{}, err
	}

	now := time.Now()
	nyear, nmonth, nday := now.Date()
	t := time.Date(nyear, nmonth, nday, 0, 0, 0, 0, now.Location())

	startTime := t.Add(dayStartTimeActivity)
	endTime := t.Add(dayEndTimeActivity)

	var (
		budgetSpent Budget
		timeSpent   int
	)

	itineraryDays := make([]ItineraryDay, days)
	day := 1
	activitiesInBudget := uc.filterOutActivitiesNotInBudget(budgetDay, activities)

	for Day(day) <= days {
		itineraryDay, err := uc.suggestItineraryDay(startTime, endTime, budgetDay, day, dayMinimalActivities, activitiesInBudget)
		if err != nil {
			return Schedule{}, err
		}

		itineraryDays[day-1] = itineraryDay

		budgetSpent += itineraryDay.BudgetSpent()
		timeSpent += itineraryDay.TimeSpent()

		activities = uc.filterOutActivitiesSuggested(activities, itineraryDay.Activities()...)

		day++
	}

	var canNotAddMore int
	day = 1

	for true {
		budgetLeft := budget - budgetSpent
		activities = uc.filterOutActivitiesNotInBudget(budgetLeft, activities)
		if len(activities) == 0 {
			break
		}

		if Day(day) > days {
			day = 1
		}

		itineraryDay := itineraryDays[day-1]
		startTime = itineraryDay.EndingTime().Add(dayActivityInterval)

		extraItineraryDay, err := uc.suggestItineraryDay(startTime, endTime, budgetLeft, day, 1, activities)
		if err != nil {
			if err == ErrActivityCanNotStart || err == errActivityCanNotEnd {
				if Day(canNotAddMore) == days {
					break
				}

				canNotAddMore++
				continue
			}

			return Schedule{}, err
		}

		budgetSpent += extraItineraryDay.BudgetSpent()
		timeSpent += extraItineraryDay.TimeSpent()

		activities = uc.filterOutActivitiesSuggested(activities, extraItineraryDay.Activities()...)

		itineraryDay.Itineraries = append(itineraryDay.Itineraries, extraItineraryDay.Itineraries...)
		itineraryDays[day-1] = itineraryDay

		day++
	}

	budgetPercentageConsume := budget.PercentageConsume(budgetSpent)
	if budgetPercentageConsume < budgetMinimalPercentageConsume {
		return Schedule{}, ErrBudgetNotConsumed
	}

	return Schedule{
		Summary: Summary{
			BudgetSpent: budgetSpent,
			TimeSpent:   timeSpent,
		},
		Days: itineraryDays,
	}, nil
}

func (uc *randomPersonalizeSchedule) getActivitiesOrdered(budget Budget) ([]Activity, error) {
	activities, err := uc.activityFinder.FindAll()
	if err != nil {
		return nil, err
	}

	sort.Sort(ActivityByPrice(activities))

	return activities, nil
}

func (uc *randomPersonalizeSchedule) filterOutActivitiesNotInBudget(budget Budget, activities []Activity) []Activity {
	var activitiesInBudget []Activity

	for i := 0; i < len(activities); i++ {
		activity := activities[i]
		if Budget(activity.Price) > budget {
			break
		}

		activitiesInBudget = append(activitiesInBudget, activity)
	}

	return activitiesInBudget
}

func (uc *randomPersonalizeSchedule) filterOutActivitiesSuggested(activities []Activity, activitiesSuggested ...Activity) []Activity {
	var (
		filterOut []Activity
		removed   int
		i         int
	)

	l := len(activities)

	for i < l {
		var found bool

		activity := activities[i]

		for _, activitySuggested := range activitiesSuggested {
			if activities[i].ID == activitySuggested.ID {
				found = true

				break
			}
		}

		if found {
			removed++

			if removed == len(activitiesSuggested) {
				break
			}

			i++
			continue
		}

		filterOut = append(filterOut, activity)

		i++
	}

	if i < l-1 {
		filterOut = append(filterOut, activities[i+1:]...)
	}

	return filterOut
}

// suggestItineraryDay ...
//
// Returns error
// 		ErrInsufficientBudgetForItinerary when there is not activity available for the budget intended for an activity
func (uc *randomPersonalizeSchedule) suggestItineraryDay(
	startTime,
	endTime time.Time,
	budget Budget,
	day int,
	numActivities int,
	activities []Activity,
) (ItineraryDay, error) {
	if startTime.After(endTime) {
		return ItineraryDay{}, ErrActivityCanNotStart
	}

	itineraries := make([]Itinerary, numActivities)
	numItinerary := 0

	for startTime.Before(endTime) && numItinerary < numActivities {
		activitiesInBudget := uc.filterOutActivitiesNotInBudget(Budget(int(budget)/(numActivities-numItinerary)), activities)
		if len(activitiesInBudget) == 0 {
			return ItineraryDay{}, ErrInsufficientBudgetForItinerary
		}

		itinerary := uc.suggestItinerary(budget, startTime, activitiesInBudget)

		activities = uc.filterOutActivitiesSuggested(activities, itinerary.Activity)

		if startTime.Add(itinerary.Activity.Duration).After(endTime) {
			if len(activities) == 0 {
				return ItineraryDay{}, errActivityCanNotEnd
			}

			continue
		}

		itineraries[numItinerary] = itinerary

		budget -= Budget(itinerary.Activity.Price)

		startTime = startTime.Add(itinerary.Activity.Duration).Add(dayActivityInterval)

		numItinerary++

	}

	return ItineraryDay{
		Day:         Day(day),
		Itineraries: itineraries,
	}, nil
}

func (uc *randomPersonalizeSchedule) suggestItinerary(
	budget Budget,
	activityTime time.Time,
	activities []Activity,
) Itinerary {
	rand.Seed(time.Now().UTC().UnixNano())

	activity := activities[rand.Intn(len(activities))]

	return Itinerary{
		Start:    activityTime,
		Activity: activity,
	}
}
