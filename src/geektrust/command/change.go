package command

import (
	"geektrust/enum"
	"geektrust/portfolio"
	"geektrust/util"
)

type change struct {
	month       enum.Month
	percentages []float64
}

func (a *change) Execute(portfolio portfolio.Portfolio) error {

	typewiseChanges := map[enum.PortfolioType]float64{
		enum.Equity: a.percentages[enum.Equity-1],
		enum.Debt:   a.percentages[enum.Debt-1],
		enum.Gold:   a.percentages[enum.Gold-1],
	}

	return portfolio.Change(a.month, typewiseChanges)
}

func getChangeCommand(params []string) (changeCommand *change, err error) {

	percentages, err := util.GetPercentagesFromString(params[1:3])
	if err != nil {
		return
	}
	month, err := enum.GetMonthFromString(params[4])
	if err != nil {
		return
	}

	return &change{month: month, percentages: percentages}, nil
}
