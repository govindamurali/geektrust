package command

import (
	"geektrust/enum"
	mocks2 "geektrust/output/mocks"
	"geektrust/portfolio"
	"geektrust/portfolio/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBalance_Execute_Success(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	displayMock := mocks2.NewDisplay(t)

	month := enum.March
	porfolioMock.On("GetBalance", month).Return(portfolio.Allocation{enum.Equity: 100, enum.Debt: 22, enum.Gold: 44}, nil)
	displayMock.On("Output", "100 22 44").Return()

	balance := balance{displayMock, month}
	err := balance.Execute(porfolioMock)
	displayMock.AssertCalled(t, "Output", "100 22 44")
	assert.Nil(t, err)
}

func TestBalance_Execute_Failure(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	displayMock := mocks2.NewDisplay(t)

	month := enum.March
	porfolioMock.On("GetBalance", month).Return(nil, nil)
	displayMock.On("Output", mock.Anything).Return()

	balance := balance{displayMock, month}
	err := balance.Execute(porfolioMock)
	displayMock.AssertCalled(t, "Output", "BALANCE UNAVAILABLE")
	assert.Nil(t, err)
}

func Test_GetBalance(t *testing.T) {
	displayMock := mocks2.NewDisplay(t)
	bal, err := getBalanceCommand([]string{"BALANCE", "MARCH"}, displayMock)
	assert.Nil(t, err)
	assert.Equal(t, bal, &balance{displayMock, enum.March})
}
