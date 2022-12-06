package portfolio

import (
	"geektrust/enum"
	"geektrust/errors"
)

type portfolio struct {
	initialAllocation   allocation
	monthlyAllocation   []monthlyAllocation
	sip                 allocation
	lastCalculatedMonth enum.Month
	lastAllocatedYear   int
	lastRebalancedMonth enum.Month
}

type allocation map[enum.PortfolioType]int

type monthlyAllocation map[enum.Month]allocation

type change map[enum.PortfolioType]float64

func (p *portfolio) Allocate(allocation allocation) error {
	if p.initialAllocation != nil {
		return errors.ErrPortfolioAlreadyAllocated
	}

	p.initialAllocation = allocation
	return nil
}

func (p *portfolio) StartSip(sip allocation) {
	p.sip = sip
}

func (p *portfolio) Change(month enum.Month, change change) error {

	//check december
	if month == enum.January {
		p.lastAllocatedYear += 1
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		return errors.ErrInvalidChangeMonth
	}

	p.lastCalculatedMonth = month

	monthlyAllocation := p.monthlyAllocation[p.lastAllocatedYear]

	//todo validation checks
	updatedAllocation := calculateChange(change, p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth])

	updatedSipAllocation := addSip(updatedAllocation, p.sip)

	monthlyAllocation[month] = updatedSipAllocation
	return nil
}

func (p *portfolio) GetBalance(month enum.Month) (allocation, error) {
	return p.monthlyAllocation[p.lastAllocatedYear][month], nil
}

func (p *portfolio) IsRebalanced() bool {
	return p.lastRebalancedMonth > 0
}

func (p *portfolio) GetLastRebalance() (allocation, error) {
	if !p.IsRebalanced() {
		return nil, errors.ErrInvalidCommandArguments
	}
	return p.monthlyAllocation[p.lastAllocatedYear][p.lastRebalancedMonth], nil
}

func calculateChange(change change, allocation allocation) (updatedAllocation allocation) {
	updatedAllocation = allocation
	for key, value := range change {
		updatedAllocation[key] = int(float64(updatedAllocation[key]) * (1 + value))
	}
	return updatedAllocation
}

func addSip(allocation allocation, sip allocation) allocation {
	updatedAllocation := allocation
	for key := range sip {
		updatedAllocation[key] = updatedAllocation[key] + sip[key]
	}
	return updatedAllocation
}
