package domain_test

import (
	"testing"

	"github.com/dohernandez/travels-budget/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPercentageConsumeBudget(t *testing.T) {
	testCases := []struct {
		scenario    string
		budget      domain.Budget
		budgetSpent domain.Budget
		assert      bool
	}{
		{
			scenario:    "percentage consume bigger than 60%",
			budget:      345,
			budgetSpent: 280,
			assert:      true,
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			assert.Equal(t, tc.assert, tc.budget.PercentageConsume(tc.budgetSpent) > 60)
		})
	}
}
