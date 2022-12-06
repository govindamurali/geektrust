package command

import (
	"geektrust/enum"
	"geektrust/errors"
	"geektrust/portfolio"
	"geektrust/util"
)

type allocate struct {
	amount []int
}

func (a *allocate) Execute(portfolio portfolio.Portfolio) error {
	allocation := map[enum.PortfolioType]int{
		enum.Equity: a.amount[enum.Equity-1],
		enum.Debt:   a.amount[enum.Debt-1],
		enum.Gold:   a.amount[enum.Gold-1],
	}
	return portfolio.Allocate(allocation)
}

func getAllocateCommand(params []string) (allocation *allocate, err error) {

	allocationValues, err := util.GetSlicesStringToInt(params[1:])
	if err != nil {
		// TODO return inside error values
		return allocation, errors.ErrInvalidCommandArguments
	}
	return &allocate{amount: allocationValues}, nil
}
