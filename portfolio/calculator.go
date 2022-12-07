package portfolio

import "geektrust/enum"

type calculator struct{}

func (c *calculator) rebalanceAllocation(allocation *ClasswiseAllocation, initialAllocation ClasswiseAllocation) {
	currentTotal := allocation.Equity + allocation.Debt + allocation.Gold
	initialTotal := initialAllocation.Equity + initialAllocation.Debt + initialAllocation.Gold

	allocation.Equity = currentTotal * initialAllocation.Equity / initialTotal
	allocation.Debt = currentTotal * initialAllocation.Debt / initialTotal
	allocation.Gold = currentTotal * initialAllocation.Gold / initialTotal
}

func (c *calculator) calculateChange(allocation *ClasswiseAllocation, change Change) {

	allocation.Equity = int(float64(allocation.Equity) * (1 + change[enum.Equity]/100))
	allocation.Debt = int(float64(allocation.Debt) * (1 + change[enum.Debt]/100))
	allocation.Gold = int(float64(allocation.Gold) * (1 + change[enum.Gold]/100))
}

func (c *calculator) addSip(allocation *ClasswiseAllocation, sip ClasswiseAllocation) {
	if sip.isEmpty() {
		return
	}
	allocation.Equity = allocation.Equity + sip.Equity
	allocation.Gold = allocation.Gold + sip.Gold
	allocation.Debt = allocation.Debt + sip.Debt
}
