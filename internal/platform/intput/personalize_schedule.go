package intput

import "github.com/urfave/cli"

// PersonalizeScheduleForm represents the data input
type PersonalizeScheduleForm struct {
	Budget int
	Days   int
}

// PersonalizeScheduleFormFromCliContext retrieves PersonalizeScheduleForm from cli.Context
func PersonalizeScheduleFormFromCliContext(cliCtx *cli.Context) PersonalizeScheduleForm {
	return PersonalizeScheduleForm{
		Budget: cliCtx.Int(BudgetKey),
		Days:   cliCtx.Int(DaysKey),
	}
}
