package command

import (
	"geektrust/enum"
	"geektrust/errors"
	mocks2 "geektrust/output/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandResolver_GetCommand_Success(t *testing.T) {
	displayMock := mocks2.NewDisplay(t)
	cm := GetCommandResolver()
	com, err := cm.GetCommand("ALLOCATE 8000 6000 3500", displayMock)
	assert.Nil(t, err)
	assert.Equal(t, com, &allocate{amount: []int{8000, 6000, 3500}})

	com2, err := cm.GetCommand("CHANGE -6.00% 21.00% -3.00% FEBRUARY", displayMock)
	assert.Nil(t, err)
	assert.Equal(t, com2, &change{enum.February, []float64{-6.0, 21.0, -3.0}})

}

func TestCommandResolver_GetCommand_Failure(t *testing.T) {
	displayMock := mocks2.NewDisplay(t)
	cm := GetCommandResolver()
	com, err := cm.GetCommand("ALLOCATE 8000 ", displayMock)
	assert.Error(t, err, errors.ErrInvalidCommandArguments)
	assert.Nil(t, com)

	com2, err := cm.GetCommand("", displayMock)
	assert.Error(t, err, errors.ErrInvalidCommand)
	assert.Nil(t, com2)

	com3, err := cm.GetCommand("RANDOM argument", displayMock)
	assert.Error(t, err, errors.ErrInvalidCommand)
	assert.Nil(t, com3)
}
