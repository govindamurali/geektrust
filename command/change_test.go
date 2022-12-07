package command

import (
	"geektrust/enum"
	"geektrust/portfolio"
	"geektrust/portfolio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChange_Execute(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	month := enum.March
	porfolioMock.On("Change", month, portfolio.Change{
		enum.Equity: 1.0,
		enum.Debt:   2.2,
		enum.Gold:   4.5,
	}).Return(nil)
	changeCom := change{month, []float64{1.0, 2.2, 4.5}}
	err := changeCom.Execute(porfolioMock)
	assert.Nil(t, err)
}

func Test_GetChange(t *testing.T) {
	changeCom, err := getChangeCommand([]string{"CHANGE", "1.0", "2.2", "4.5", "MARCH"})
	assert.Nil(t, err)
	assert.Equal(t, changeCom, &change{enum.March, []float64{1.0, 2.2, 4.5}})
}
