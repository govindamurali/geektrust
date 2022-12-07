package command

import (
	"geektrust/enum"
	mocks2 "geektrust/output/mocks"
	"geektrust/portfolio"
	"geektrust/portfolio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRebalance_Execute_Success(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	displayMock := mocks2.NewDisplay(t)

	porfolioMock.On("IsRebalanced").Return(true)
	porfolioMock.On("GetLastRebalance").Return(portfolio.Allocation{enum.Equity: 100, enum.Debt: 22, enum.Gold: 44}, nil)
	displayMock.On("Output", "100 22 44").Return()

	rebalance := rebalance{displayMock}
	err := rebalance.Execute(porfolioMock)

	assert.Nil(t, err)
	porfolioMock.AssertCalled(t, "IsRebalanced")
	porfolioMock.AssertCalled(t, "GetLastRebalance")
	displayMock.AssertCalled(t, "Output", "100 22 44")
}

func TestRebalance_Execute_Failure(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	displayMock := mocks2.NewDisplay(t)

	porfolioMock.On("IsRebalanced").Return(false)
	displayMock.On("Output", "CANNOT_REBALANCE").Return()

	rebalance := rebalance{displayMock}
	err := rebalance.Execute(porfolioMock)

	assert.Nil(t, err)
	porfolioMock.AssertCalled(t, "IsRebalanced")
	porfolioMock.AssertNotCalled(t, "GetLastRebalance")
	displayMock.AssertCalled(t, "Output", "CANNOT_REBALANCE")
}

func Test_GetRebalance(t *testing.T) {
	displayMock := mocks2.Display{}
	rebalanceCom, err := getRebalanceCommand(&displayMock)
	assert.Nil(t, err)
	assert.Equal(t, rebalanceCom, &rebalance{
		&displayMock,
	})
}
