package command

import (
	"geektrust/enum"
	"geektrust/output"
	"geektrust/portfolio"
	"strconv"
)

type balance struct {
	display output.Display
	month   enum.Month
}

func (b balance) Execute(portfolio portfolio.Portfolio) error {
	allocation, err := portfolio.GetBalance(b.month)
	if err != nil {
		return err
	}

	if allocation == nil {
		//todo move to error
		b.display.Output("BALANCE UNAVAILABLE")
	}

	//todo change
	b.display.Output(strconv.Itoa(allocation[enum.Equity]) + " " + strconv.Itoa(allocation[enum.Debt]) + " " + strconv.Itoa(allocation[enum.Gold]))
	return nil
}

func getBalanceCommand(args []string, display output.Display) (*balance, error) {
	monthVal, err := enum.GetMonthFromString(args[1])
	if err != nil {
		return nil, err
	}
	return &balance{
		display: display,
		month:   monthVal,
	}, nil
}
