package domain

import (
	"encoding/json"
	"time"
)

// Activity represents the activity entity
type Activity struct {
	ID       int
	Duration time.Duration
	Price    int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// to enforces Duration value
func (a *Activity) UnmarshalJSON(data []byte) error {
	var activity struct {
		ID       int `json:"id"`
		Duration int `json:"duration"`
		Price    int `json:"price"`
	}

	if err := json.Unmarshal(data, &activity); err != nil {
		return err
	}

	a.ID = activity.ID
	a.Duration = time.Duration(activity.Duration) * time.Minute
	a.Price = activity.Price

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// to enforces Duration value
func (a Activity) MarshalJSON() ([]byte, error) {
	activity := struct {
		ID       int `json:"id"`
		Duration int `json:"duration"`
		Price    int `json:"price"`
	}{
		ID:       a.ID,
		Duration: int(a.Duration.Minutes()),
		Price:    a.Price,
	}

	return json.Marshal(activity)
}

// ActivityFinder defines the way to find activity
type ActivityFinder interface {
	FindAll() ([]Activity, error)
}

// ActivityByPrice is a collection of Activity instances and implements the sort
// interface. See the sort package for more details.
// https://golang.org/pkg/sort/
type ActivityByPrice []Activity

// Len returns the length of activities. The number of Activity instances
// on the slice.
func (s ActivityByPrice) Len() int { return len(s) }

// Swap is needed for the sort interface to replace the Activity objects
// at two different positions in the slice.
func (s ActivityByPrice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less is needed for the sort interface to compare two Activity objects on the
// slice. If checks if one is less than the other.
func (s ActivityByPrice) Less(i, j int) bool {
	return s[i].Price < s[j].Price
}
