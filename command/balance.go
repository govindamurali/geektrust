package command

import (
	"geektrust/enum"
	"geektrust/output"
	"geektrust/portfolio"
)

type balance struct {
	display output.Display
	month   enum.Month
}

const messageBalanceUnavailable = "BALANCE UNAVAILABLE"

func (b balance) Execute(portfolio portfolio.Portfolio) error {
	bal, err := portfolio.GetBalance(b.month)
	if err != nil {
		return err
	}

	if bal == nil {
		b.display.Output(messageBalanceUnavailable)
		return nil
	}

	b.display.Output(bal.ToString())
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
