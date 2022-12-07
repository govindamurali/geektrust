package portfolio

import (
	"geektrust/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPortfolio_Allocate(t *testing.T) {

	allocation := ClasswiseAllocationMap{
		enum.Equity: 1,
		enum.Debt:   2,
		enum.Gold:   3,
	}

	folio := GetFreshPortfolio()
	folio.Allocate(allocation)

	expectedFolio := portfolio{
		initialAllocation: ClasswiseAllocation{
			Equity: 1,
			Debt:   2,
			Gold:   3,
		},
		allocated:           true,
		monthlyAllocation:   make([]MonthlyAllocation, 0),
		sip:                 ClasswiseAllocation{},
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: 0,
		calculator:          calculator{},
	}

	assert.Equal(t, folio, &expectedFolio)
}
