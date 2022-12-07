package portfolio

import "geektrust/enum"

type Portfolio interface {
	StartSip(sip Allocation)
	Change(month enum.Month, change Change) error
	GetBalance(month enum.Month) (Allocation, error)
	IsRebalanced() bool
	GetLastRebalance() (Allocation, error)
	Allocate(allocation Allocation) error
}
type Allocation map[enum.PortfolioType]int

type MonthlyAllocation map[enum.Month]Allocation

type Change map[enum.PortfolioType]float64

func GetEmptyPortfolio() Portfolio {
	return &portfolio{
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: 0,
	}
}
