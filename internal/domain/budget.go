package domain

const (
	// budgetMinimalPercentageConsume defines the minimal budget that has to be consume
	budgetMinimalPercentageConsume = float64(60)
)

// Budget represents the budget value
type Budget int

// PercentageConsume calculate how much percentage represent
func (b Budget) PercentageConsume(budget Budget) float64 {
	return float64(int(budget) * 100 / int(b))
}
