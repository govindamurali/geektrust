package portfolio

import (
	"geektrust/enum"
	"github.com/stretchr/testify/assert"
	"sync"
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

func TestPortfolio_StartSip(t *testing.T) {
	folio := GetFreshPortfolio()
	folio.StartSip(ClasswiseAllocationMap{
		enum.Equity: 10,
		enum.Debt:   22,
		enum.Gold:   30,
	})

	expectedFolio := portfolio{
		initialAllocation: ClasswiseAllocation{},
		allocated:         false,
		monthlyAllocation: make([]MonthlyAllocation, 0),
		sip: ClasswiseAllocation{
			Equity: 10,
			Debt:   22,
			Gold:   30,
		},
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: 0,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	assert.Equal(t, folio, &expectedFolio)
}

func TestPortfolio_GetBalance(t *testing.T) {
	folio := portfolio{
		initialAllocation: ClasswiseAllocation{},
		allocated:         false,
		monthlyAllocation: []MonthlyAllocation{
			{
				enum.January: ClasswiseAllocation{
					Equity: 2,
					Debt:   3,
					Gold:   4,
				},
				enum.February: ClasswiseAllocation{
					Equity: 5,
					Debt:   6,
					Gold:   7,
				}, enum.March: ClasswiseAllocation{
					Equity: 7,
					Debt:   8,
					Gold:   9,
				},
			},
		},
		sip:                 ClasswiseAllocation{},
		lastCalculatedMonth: 0,
		lastAllocatedYear:   0,
		lastRebalancedMonth: 0,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	bal, err := folio.GetBalance(enum.February)
	assert.Nil(t, err)
	assert.Equal(t, bal, ClasswiseAllocationMap{enum.Equity: 5, enum.Debt: 6, enum.Gold: 7})
}
