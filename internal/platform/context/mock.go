package context

import (
	"flag"

	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/urfave/cli"
)

// NewPersonalizeScheduleCliContextMock returns a cli.NewContext mocked for test purpose
// nolint:unused
func NewPersonalizeScheduleCliContextMock(days, budget int) *cli.Context {
	set := flag.NewFlagSet("test", 0)
	set.Int(intput.DaysKey, days, "days")
	set.Int(intput.BudgetKey, budget, "budget")

	return cli.NewContext(nil, set, nil)
}
