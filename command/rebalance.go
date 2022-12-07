package command

import (
	"geektrust/output"
	"geektrust/portfolio"
)

type rebalance struct {
	display output.Display
}

const messageCannotRebalance = "CANNOT_REBALANCE"

func (r rebalance) Execute(portfolio portfolio.Portfolio) error {

	if !portfolio.IsRebalanced() {
		r.display.Output(messageCannotRebalance)
		return nil
	}

	rebalanceVal, err := portfolio.GetLastRebalance()
	if err != nil {
		return err
	}

	r.display.Output(rebalanceVal.ToString())
	return nil
}

func getRebalanceCommand(display output.Display) (*rebalance, error) {
	return &rebalance{
		display: display,
	}, nil
}
