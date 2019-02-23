package intput_test

import (
	"testing"

	"github.com/dohernandez/travels-budget/internal/platform/context"
	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeScheduleFormFromCliContext(t *testing.T) {
	budget := 680
	days := 2

	cliCtx := context.NewPersonalizeScheduleCliContextMock(days, budget)

	p := intput.PersonalizeScheduleFormFromCliContext(cliCtx)
	assert.Equal(t, budget, p.Budget)
	assert.Equal(t, days, p.Days)
}
