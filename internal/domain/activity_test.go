package domain_test

import (
	"encoding/json"
	"regexp"
	"sort"
	"testing"
	"time"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSortActivityByPrice(t *testing.T) {
	// Reduce to 3 the amount of activity, there is no need to run the test with the whole collection
	// with 3 activities is enough
	activities := domain.NewActivitiesMock()[:3]

	sorted := []domain.Activity{
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
			ID:       7495,
			Duration: 120 * time.Minute,
			Price:    423,
		},
	}

	sort.Sort(domain.ActivityByPrice(activities))

	assert.Equal(t, sorted, activities)
}

func TestUnmarshalActivity(t *testing.T) {
	activityJson := `{
        "id": 7413,
        "duration": 180,
        "price": 411
    }`

	var activity domain.Activity

	err := json.Unmarshal([]byte(activityJson), &activity)
	assert.NoError(t, err, "error was not expected, given %s", err)
	assert.Equal(t, domain.Activity{
		ID:       7413,
		Duration: 180 * time.Minute,
		Price:    411,
	}, activity)
}

func TestMarshalActivity(t *testing.T) {
	activity := domain.Activity{
		ID:       7413,
		Duration: 180 * time.Minute,
		Price:    411,
	}

	activityJson, err := json.Marshal(activity)
	assert.NoError(t, err, "error was not expected, given %s", err)
	assert.Equal(t, regexp.MustCompile(`\s+`).ReplaceAllString(`{
        "id": 7413,
        "duration": 180,
        "price": 411
    }`, ""), string(activityJson))
}
