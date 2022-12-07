package command

import (
	"geektrust/portfolio"
	"geektrust/portfolio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSip_Execute(t *testing.T) {
	porfolioMock := mocks.NewPortfolio(t)
	porfolioMock.On("StartSip", portfolio.Allocation{1: 1, 2: 2, 3: 3}).Return(nil)
	sipCom := sip{amount: []int{1, 2, 3}}
	err := sipCom.Execute(porfolioMock)

	assert.Nil(t, err)
	porfolioMock.AssertCalled(t, "StartSip", portfolio.Allocation{1: 1, 2: 2, 3: 3})
}

func Test_GetSip(t *testing.T) {
	sipCom, err := getSipCommand([]string{"SIP", "6000", "3000", "1000"})
	assert.Nil(t, err)
	assert.Equal(t, sipCom, &sip{[]int{6000, 3000, 1000}})
}
