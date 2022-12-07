package portfolio

import (
	"fmt"
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

func (p *portfolio) StartSip(sip ClasswiseAllocationMap) {
	p.sip = sip.toStruct()
}

func (p *portfolio) Change(month enum.Month, change Change) error {

	//todo check december
	// todo race conditions
	var lastAllocation ClasswiseAllocation
	var updatedAllocation ClasswiseAllocation
	if month == enum.January {
		p.monthlyAllocation = append(p.monthlyAllocation, make(MonthlyAllocation))
		p.lastAllocatedYear += 1
		lastAllocation = p.initialAllocation
		updatedAllocation = lastAllocation
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		return errors.ErrInvalidChangeMonth
	} else {
		lastAllocation = p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth]
		updatedAllocation = addSip(lastAllocation, p.sip)
	}

	fmt.Print(month, updatedAllocation)
	//todo validation checks
	updatedAllocation = calculateChange(change, updatedAllocation)
	fmt.Print(month, lastAllocation)

	if month.IsRebalanceRequired() {
		updatedAllocation = p.getRebalancedAllocation(updatedAllocation)
		p.lastRebalancedMonth = month
		fmt.Print(month, updatedAllocation)
	}

	fmt.Println()
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
	updatedAllocation = allocation

	updatedAllocation.Equity = int(float64(updatedAllocation.Equity) * (1 + change[enum.Equity]/100))
	updatedAllocation.Debt = int(float64(updatedAllocation.Debt) * (1 + change[enum.Debt]/100))
	updatedAllocation.Gold = int(float64(updatedAllocation.Gold) * (1 + change[enum.Gold]/100))

	return updatedAllocation
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
