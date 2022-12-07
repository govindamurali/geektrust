package command

import (
	"geektrust/portfolio"
	"geektrust/portfolio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllocation_Execute(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	porfolioMock.On("Allocate", portfolio.Allocation{1: 1, 2: 2, 3: 3}).Return(nil)
	allocate := allocate{amount: []int{1, 2, 3}}
	err := allocate.Execute(porfolioMock)
	
	assert.Nil(t, err)
	porfolioMock.AssertCalled(t, "Allocate", portfolio.Allocation{1: 1, 2: 2, 3: 3})
}

func Test_GetAllocate(t *testing.T) {
	allocation, err := getAllocateCommand([]string{"ALLOCATE", "6000", "3000", "1000"})
	assert.Nil(t, err)
	assert.Equal(t, allocation, &allocate{[]int{6000, 3000, 1000}})
}
