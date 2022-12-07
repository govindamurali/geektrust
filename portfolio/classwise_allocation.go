package portfolio

import "geektrust/enum"

type ClasswiseAllocationMap map[enum.PortfolioType]int

type ClasswiseAllocation struct {
	Equity int
	Debt   int
	Gold   int
}

func (c *ClasswiseAllocation) isEmpty() bool {
	return c.Equity == 0 && c.Debt == 0 && c.Gold == 0
}

func (c *ClasswiseAllocation) toMap() ClasswiseAllocationMap {
	return map[enum.PortfolioType]int{
		enum.Equity: c.Equity,
		enum.Debt:   c.Debt,
		enum.Gold:   c.Gold,
	}
}

func (c ClasswiseAllocationMap) toStruct() ClasswiseAllocation {
	return ClasswiseAllocation{
		Equity: c[enum.Equity],
		Gold:   c[enum.Gold],
		Debt:   c[enum.Debt],
	}
}
