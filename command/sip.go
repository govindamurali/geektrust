package command

import (
	"geektrust/enum"
	"geektrust/errors"
	"geektrust/portfolio"
	"geektrust/util"
)

type sip struct {
	amount []int
}

func (s *sip) Execute(portfolio portfolio.Portfolio) error {
	sip := map[enum.PortfolioType]int{
		enum.Equity: s.amount[enum.Equity-1],
		enum.Debt:   s.amount[enum.Debt-1],
		enum.Gold:   s.amount[enum.Gold-1],
	}
	portfolio.StartSip(sip)
	return nil
}

func getSipCommand(params []string) (sipCommand *sip, err error) {

	sipValues, err := util.GetSlicesStringToInt(params[1:])
	if err != nil {
		return sipCommand, errors.GetAppendedErrors(errors.ErrInvalidCommandArguments, err)
	}
	return &sip{amount: sipValues}, nil
}
