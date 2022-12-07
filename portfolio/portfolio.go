package portfolio

import (
	"geektrust/enum"
	"geektrust/errors"
)

type portfolio struct {
	initialAllocation   ClasswiseAllocation
	allocated           bool
	monthlyAllocation   []MonthlyAllocation
	sip                 ClasswiseAllocation
	lastCalculatedMonth enum.Month
	lastAllocatedYear   int
	lastRebalancedMonth enum.Month
	calculator          calculator
}

func (p *portfolio) Allocate(allocationMap ClasswiseAllocationMap) error {
	allocation := allocationMap.toStruct()
	if p.allocated {
		return errors.ErrPortfolioAlreadyAllocated
	}
	p.initialAllocation = allocation
	p.allocated = true
	return nil
}

func (p *portfolio) StartSip(sipMap ClasswiseAllocationMap) {
	p.sip = sipMap.toStruct()
}

func (p *portfolio) Change(month enum.Month, change Change) error {

	if !p.allocated {
		return errors.ErrPortfolioNotAllocated
	}

	var lastAllocation ClasswiseAllocation
	var updatedAllocation ClasswiseAllocation
	if month == enum.January {
		p.monthlyAllocation = append(p.monthlyAllocation, make(MonthlyAllocation))
		if p.lastCalculatedMonth == enum.December {
			lastAllocation = p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth]
			p.lastAllocatedYear += 1
		} else {
			lastAllocation = p.initialAllocation
		}
		updatedAllocation = lastAllocation
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		return errors.ErrInvalidChangeMonth
	} else {
		lastAllocation = p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth]
		p.calculator.addSip(&lastAllocation, p.sip)
		updatedAllocation = lastAllocation
	}

	p.calculator.calculateChange(&updatedAllocation, change)

	if month.IsRebalanceRequired() {
		p.calculator.rebalanceAllocation(&updatedAllocation, p.initialAllocation)
		p.lastRebalancedMonth = month
	}

	monthlyAllocation := p.monthlyAllocation[p.lastAllocatedYear]
	monthlyAllocation[month] = updatedAllocation
	p.monthlyAllocation[p.lastAllocatedYear] = monthlyAllocation
	p.lastCalculatedMonth = month
	return nil
}

func (p *portfolio) GetBalance(month enum.Month) (ClasswiseAllocationMap, error) {

	allocation := p.monthlyAllocation[p.lastAllocatedYear][month]
	return allocation.toMap(), nil
}

func (p *portfolio) IsRebalanced() bool {
	return p.lastRebalancedMonth > 0
}

func (p *portfolio) GetLastRebalance() (classwiseAllocation ClasswiseAllocationMap, err error) {
	if !p.IsRebalanced() {
		err = errors.ErrInvalidCommandArguments
		return
	}
	allocation := p.monthlyAllocation[p.lastAllocatedYear][p.lastRebalancedMonth]
	return allocation.toMap(), nil
}
