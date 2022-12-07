package portfolio

import (
	"geektrust/enum"
	"geektrust/errors"
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

func TestPortfolio_GetLastRebalance_Success(t *testing.T) {
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
				},
				enum.March: ClasswiseAllocation{
					Equity: 7,
					Debt:   8,
					Gold:   9,
				},
				enum.April: ClasswiseAllocation{
					Equity: 10,
					Debt:   11,
					Gold:   12,
				},
				enum.June: ClasswiseAllocation{
					Equity: 13,
					Debt:   14,
					Gold:   15,
				},
				enum.July: ClasswiseAllocation{
					Equity: 16,
					Debt:   17,
					Gold:   18,
				},
			},
		},
		sip:                 ClasswiseAllocation{},
		lastCalculatedMonth: enum.July,
		lastAllocatedYear:   0,
		lastRebalancedMonth: enum.June,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	bal, err := folio.GetLastRebalance()
	assert.Nil(t, err)
	assert.Equal(t, bal, ClasswiseAllocationMap{enum.Equity: 13, enum.Debt: 14, enum.Gold: 15})
}

func TestPortfolio_GetLastRebalance_Failure(t *testing.T) {
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
				},
				enum.March: ClasswiseAllocation{
					Equity: 7,
					Debt:   8,
					Gold:   9,
				},
				enum.April: ClasswiseAllocation{
					Equity: 10,
					Debt:   11,
					Gold:   12,
				},
			},
		},
		sip:                 ClasswiseAllocation{},
		lastCalculatedMonth: enum.July,
		lastAllocatedYear:   0,
		lastRebalancedMonth: enum.InvalidMonth,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	bal, err := folio.GetLastRebalance()
	assert.Error(t, err, errors.ErrInvalidCommandArguments)
	assert.Nil(t, bal)
}

func TestPortfolio_Change(t *testing.T) {
	folio := portfolio{
		initialAllocation: ClasswiseAllocation{
			Equity: 1,
			Debt:   2,
			Gold:   3,
		},
		allocated: true,
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
				},
				enum.March: ClasswiseAllocation{
					Equity: 7,
					Debt:   8,
					Gold:   9,
				},
				enum.April: ClasswiseAllocation{
					Equity: 10,
					Debt:   11,
					Gold:   12,
				},
				enum.June: ClasswiseAllocation{
					Equity: 13,
					Debt:   14,
					Gold:   15,
				},
				enum.July: ClasswiseAllocation{
					Equity: 16,
					Debt:   17,
					Gold:   18,
				},
			},
		},
		sip: ClasswiseAllocation{
			Equity: 2,
			Debt:   3,
			Gold:   4,
		},
		lastCalculatedMonth: enum.July,
		lastAllocatedYear:   0,
		lastRebalancedMonth: enum.June,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	err := folio.Change(enum.August, Change{
		enum.Equity: 10,
		enum.Debt:   20,
		enum.Gold:   30,
	})

	expectedFolio := portfolio{
		initialAllocation: ClasswiseAllocation{
			Equity: 1,
			Debt:   2,
			Gold:   3,
		},
		allocated: true,
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
				},
				enum.March: ClasswiseAllocation{
					Equity: 7,
					Debt:   8,
					Gold:   9,
				},
				enum.April: ClasswiseAllocation{
					Equity: 10,
					Debt:   11,
					Gold:   12,
				},
				enum.June: ClasswiseAllocation{
					Equity: 13,
					Debt:   14,
					Gold:   15,
				},
				enum.July: ClasswiseAllocation{
					Equity: 16,
					Debt:   17,
					Gold:   18,
				},
				enum.August: ClasswiseAllocation{
					Equity: 19,
					Debt:   24,
					Gold:   28,
				},
			},
		},
		sip: ClasswiseAllocation{
			Equity: 2,
			Debt:   3,
			Gold:   4,
		},
		lastCalculatedMonth: enum.August,
		lastAllocatedYear:   0,
		lastRebalancedMonth: enum.June,
		calculator:          calculator{},
		mutex:               sync.RWMutex{},
	}

	assert.Nil(t, err)
	assert.Equal(t, folio, expectedFolio)
}
