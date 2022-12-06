package command

import (
	"geektrust/enum"
	"geektrust/output"
	"geektrust/portfolio"
	"strconv"
)

type rebalance struct {
	display output.IDisplay
}

func (r rebalance) Execute(portfolio portfolio.Portfolio) error {

	if !portfolio.IsRebalanced() {
		//todo rethink
		r.display.Display("CANNOT_REBALANCE")
	}

	//todo check order
	rebalanceVal, err := portfolio.GetLastRebalance()
	if err != nil {
		return err
	}
	r.display.Display(strconv.Itoa(rebalanceVal[enum.Equity]) + " " + strconv.Itoa(rebalanceVal[enum.Debt]) + " " + strconv.Itoa(rebalanceVal[enum.Gold]))
	return nil
}

func getRebalanceCommand(display output.IDisplay) (*rebalance, error) {
	return &rebalance{
		display: display,
	}, nil
}
