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
		updatedAllocation = addSip(lastAllocation, p.sip)
	}

	updatedAllocation = calculateChange(change, updatedAllocation)

	if month.IsRebalanceRequired() {
		updatedAllocation = p.getRebalancedAllocation(updatedAllocation)
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

func (p *portfolio) getRebalancedAllocation(allocation ClasswiseAllocation) ClasswiseAllocation {
	currentTotal := allocation.Equity + allocation.Debt + allocation.Gold
	initialTotal := p.initialAllocation.Equity + p.initialAllocation.Debt + p.initialAllocation.Gold

	allocation.Equity = currentTotal * p.initialAllocation.Equity / initialTotal
	allocation.Debt = currentTotal * p.initialAllocation.Debt / initialTotal
	allocation.Gold = currentTotal * p.initialAllocation.Gold / initialTotal

	return allocation
}

func calculateChange(change Change, allocation ClasswiseAllocation) (updatedAllocation ClasswiseAllocation) {

	allocation.Equity = int(float64(allocation.Equity) * (1 + change[enum.Equity]/100))
	allocation.Debt = int(float64(allocation.Debt) * (1 + change[enum.Debt]/100))
	allocation.Gold = int(float64(allocation.Gold) * (1 + change[enum.Gold]/100))

	return allocation
}

func addSip(allocation ClasswiseAllocation, sip ClasswiseAllocation) ClasswiseAllocation {
	if sip.isEmpty() {
		return allocation
	}
	updatedAllocation := allocation
	updatedAllocation.Equity = updatedAllocation.Equity + sip.Equity
	updatedAllocation.Gold = updatedAllocation.Gold + sip.Gold
	updatedAllocation.Debt = updatedAllocation.Debt + sip.Debt
	return updatedAllocation
}
