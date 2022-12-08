package portfolio

import (
	"geektrust/enum"
	"geektrust/errors"
	"sync"
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
	mutex               sync.RWMutex
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

	allocation, err := p.getPreviousMonthAllocation(month)
	if err != nil {
		return err
	}

	if month != enum.January || p.lastAllocatedYear > 0 {
		p.calculator.addSip(&allocation, p.sip)
	}

	p.calculator.calculateChange(&allocation, change)

	if month.IsRebalanceMonth() {
		p.calculator.rebalancePortfolio(&allocation, p.initialAllocation)
		p.lastRebalancedMonth = month
	}

	p.allocateForMonth(month, allocation)

	return nil
}

func (p *portfolio) GetBalance(month enum.Month) (ClasswiseAllocationMap, error) {

	p.mutex.RLock()
	defer p.mutex.RUnlock()
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
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	allocation := p.monthlyAllocation[p.lastAllocatedYear][p.lastRebalancedMonth]
	return allocation.toMap(), nil
}

func (p *portfolio) getPreviousMonthAllocation(month enum.Month) (lastAllocation ClasswiseAllocation, err error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if month == enum.January {
		p.monthlyAllocation = append(p.monthlyAllocation, make(MonthlyAllocation))
		if p.lastCalculatedMonth == enum.December {
			lastAllocation = p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth]
			p.lastAllocatedYear += 1
		} else {
			lastAllocation = p.initialAllocation
		}
	} else if int(month)-int(p.lastCalculatedMonth) != 1 {
		err = errors.ErrInvalidChangeMonth
		return
	} else {
		lastAllocation = p.monthlyAllocation[p.lastAllocatedYear][p.lastCalculatedMonth]
	}

	return
}

func (p *portfolio) allocateForMonth(month enum.Month, allocation ClasswiseAllocation) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	monthlyAllocation := p.monthlyAllocation[p.lastAllocatedYear]
	monthlyAllocation[month] = allocation
	p.monthlyAllocation[p.lastAllocatedYear] = monthlyAllocation
	p.lastCalculatedMonth = month
}
