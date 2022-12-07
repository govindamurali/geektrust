package portfolio

import "geektrust/enum"

type Portfolio interface {
	StartSip(sip ClasswiseAllocationMap)
	Change(month enum.Month, change Change) error
	GetBalance(month enum.Month) (ClasswiseAllocationMap, error)
	IsRebalanced() bool
	GetLastRebalance() (ClasswiseAllocationMap, error)
	Allocate(allocation ClasswiseAllocationMap) error
}

type MonthlyAllocation map[enum.Month]ClasswiseAllocation

type Change map[enum.PortfolioType]float64

func GetFreshPortfolio() Portfolio {
	return &portfolio{
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: enum.InvalidMonth,
		monthlyAllocation:   make([]MonthlyAllocation, 0),
		calculator:          calculator{},
	}
}
