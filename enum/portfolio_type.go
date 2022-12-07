package enum

import (
	"geektrust/errors"
	"strings"
)

type PortfolioType int

const (
	Unsupported PortfolioType = iota
	Equity
	Debt
	Gold
)

var portfolios = map[string]PortfolioType{
	"equity": Equity,
	"debt":   Debt,
	"gold":   Gold,
}

func (p PortfolioType) String() string {
	return [...]string{"Equity", "Debt", "Gold"}[p-1]
}

func GetPortfolioTypeFromString(s string) (PortfolioType, error) {
	if s, ok := portfolios[strings.ToLower(s)]; ok {
		return s, nil
	}
	return Unsupported, errors.ErrInvalidChangeMonth
}
