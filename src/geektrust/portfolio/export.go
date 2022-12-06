package portfolio

import "geektrust/enum"

type Portfolio interface {
	StartSip(sip allocation)
	Change(month enum.Month, change change) error
	GetBalance(month enum.Month) (allocation, error)
	IsRebalanced() bool
	GetLastRebalance() (allocation, error)
	Allocate(allocation allocation) error
}

func GetEmptyPortfolio() Portfolio {
	return &portfolio{
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: 0,
	}
}
