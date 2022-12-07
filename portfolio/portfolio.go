package portfolio

import (
	"geektrust/enum"
	"geektrust/errors"
)

type portfolio struct {
	initialAllocation   Allocation
	monthlyAllocation   []MonthlyAllocation
	sip                 Allocation
	lastCalculatedMonth enum.Month
	lastAllocatedYear   int
	lastRebalancedMonth enum.Month
}

func (p *portfolio) Allocate(allocation Allocation) error {
	if p.initialAllocation != nil {
		return errors.ErrPortfolioAlreadyAllocated
	}

	p.initialAllocation = allocation
	return nil
}

func (p *portfolio) StartSip(sip Allocation) {
	p.sip = sip
}

func (p *portfolio) Change(month enum.Month, change Change) error {

	//todo check december
	// todo race conditions
	if month == enum.January {
		p.lastAllocatedYear += 1
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		return errors.ErrInvalidChangeMonth
	}

	p.lastCalculatedMonth = month

	monthlyAllocation := p.monthlyAllocation[p.lastAllocatedYear]

	updatedSipAllocation := addSip(p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth], p.sip)

	//todo validation checks
	updatedAllocation := calculateChange(change, updatedSipAllocation)

	if month.IsRebalanceRequired() {
		updatedAllocation = p.getRebalancedAllocation(updatedAllocation)
	}

	monthlyAllocation[month] = updatedAllocation
	return nil
}

func (p *portfolio) GetBalance(month enum.Month) (Allocation, error) {
	return p.monthlyAllocation[p.lastAllocatedYear][month], nil
}

func (p *portfolio) IsRebalanced() bool {
	return p.lastRebalancedMonth > 0
}

func (p *portfolio) GetLastRebalance() (Allocation, error) {
	if !p.IsRebalanced() {
		return nil, errors.ErrInvalidCommandArguments
	}
	return p.monthlyAllocation[p.lastAllocatedYear][p.lastRebalancedMonth], nil
}

func (p *portfolio) getRebalancedAllocation(allocation Allocation) Allocation {
	currentTotal := 0
	initialTotal := 0
	for _, val := range p.initialAllocation {
		initialTotal += val
	}

	for _, val := range allocation {
		currentTotal += val
	}

	for i := range allocation {
		allocation[i] = currentTotal * p.initialAllocation[i] / initialTotal
	}
	return allocation
}

func calculateChange(change Change, allocation Allocation) (updatedAllocation Allocation) {
	updatedAllocation = allocation
	for key, value := range change {
		updatedAllocation[key] = int(float64(updatedAllocation[key]) * (1 + value))
	}
	return updatedAllocation
}

func addSip(allocation Allocation, sip Allocation) Allocation {
	if sip == nil {
		return allocation
	}
	updatedAllocation := allocation
	for key := range sip {
		updatedAllocation[key] = updatedAllocation[key] + sip[key]
	}
	return updatedAllocation
}
