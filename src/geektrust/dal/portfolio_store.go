package dal

import (
	"geektrust/enum"
	"geektrust/errors"
)

type Portfolio struct {
	initialAllocation   Allocation
	monthlyAllocation   []MonthlyAllocation
	sip                 Allocation
	lastCalculatedMonth enum.Month
	lastAllocatedYear   int
	lastRebalancedMonth int
}

type Allocation map[enum.PortfolioType]int

type MonthlyAllocation map[enum.Month]Allocation

type Change map[enum.PortfolioType]float32

func GetEmptyPortfolio() *Portfolio {
	return &Portfolio{
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: -1,
	}
}

func (p *Portfolio) Allocate(allocation Allocation) error {
	if p.initialAllocation != nil {
		return errors.PortfolioAlreadyAllocatedError
	}

	p.initialAllocation = allocation
	return nil
}

func (p *Portfolio) StartSip(sip Allocation) {
	p.sip = sip
}

func (p *Portfolio) Change(month enum.Month, change Change) error {

	//check december
	if month == enum.January {
		p.lastAllocatedYear += 1
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		return errors.InvalidChangeMonth
	}

	p.lastCalculatedMonth = month

	monthlyAllocation := p.monthlyAllocation[p.lastAllocatedYear]

	calculateChange(change, p.a)

	monthlyAllocation[month] =
	return nil
}

func calculateChange(change Change, allocation Allocation) (updatedAllocation Allocation) {
	updatedAllocation = allocation
	for key, value := range change {
		updatedAllocation[key] = int(float32(updatedAllocation[key]) * (1 + value))
	}
	return updatedAllocation
}

func addSip(allocation Allocation, sip Allocation) Allocation {
	updatedAllocation := allocation
	for key := range sip {
		updatedAllocation[key] = updatedAllocation[key] + sip[key]
	}
	return updatedAllocation
}
