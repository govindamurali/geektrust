package portfolio

import (
	"geektrust/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator_rebalance(t *testing.T) {
	allocation := ClasswiseAllocation{Equity: 20, Debt: 30, Gold: 50}

	calc := calculator{}
	calc.rebalancePortfolio(&allocation, ClasswiseAllocation{Equity: 15, Debt: 15, Gold: 20})

	assert.Equal(t, ClasswiseAllocation{Equity: 30, Debt: 30, Gold: 40}, allocation)
}

func TestCalculator_change(t *testing.T) {
	allocation := ClasswiseAllocation{Equity: 20, Debt: 30, Gold: 50}

	calc := calculator{}
	calc.calculateChange(&allocation, Change{enum.Equity: 12, enum.Debt: 16.6, enum.Gold: 25.7})

	assert.Equal(t, ClasswiseAllocation{Equity: 22, Debt: 34, Gold: 62}, allocation)
}

func TestCalculator_addSip(t *testing.T) {
	allocation := ClasswiseAllocation{Equity: 20, Debt: 30, Gold: 50}

	calc := calculator{}
	calc.addSip(&allocation, ClasswiseAllocation{Equity: 15, Debt: 15, Gold: 20})

	assert.Equal(t, ClasswiseAllocation{Equity: 35, Debt: 45, Gold: 70}, allocation)
}
