package storage_test

import (
	"io"
	"strings"
	"testing"

	"time"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/dohernandez/travels-budget/internal/platform/storage"
	"github.com/stretchr/testify/assert"
)

func TestFileActivityFinder(t *testing.T) {
	testCases := []struct {
		scenario string
		file     io.Reader
		result   []domain.Activity
	}{
		{
			scenario: "Load activities successfully",
			file: strings.NewReader(`[
  				{
    				"id": 7413,
    				"duration": 180,
    				"price": 411
  				},
  				{
    				"id": 7099,
    				"duration": 180,
    				"price": 267
				},
				{
    				"id": 7048,
    				"duration": 240,
    				"price": 264
				}
			]`),
			result: []domain.Activity{
				{
					ID:       7413,
					Duration: 180 * time.Minute,
					Price:    411,
				},
				{
					ID:       7099,
					Duration: 180 * time.Minute,
					Price:    267,
				},
				{
					ID:       7048,
					Duration: 240 * time.Minute,
					Price:    264,
				},
			},
		},
		{
			scenario: "Error parsing activities json",
			file: strings.NewReader(`[
  				{
    				"id": 7413,
    				"duration": 180,
    				"price": 411,
  				},
			]`),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {

			finder := storage.NewFileActivityFinder(tc.file)

			activities, err := finder.FindAll()
			if tc.result != nil {
				assert.NoError(t, err, "error was not expected, given %s", err)
				assert.Equal(t, tc.result, activities)
			} else {
				assert.Error(t, err, "error was expected")
			}
		})
	}
}
