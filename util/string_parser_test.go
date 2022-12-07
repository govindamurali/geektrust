package util

import (
	"geektrust/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSlicesStringToInt_Success(t *testing.T) {
	intVals, err := GetSlicesStringToInt([]string{"1 ", " 2", "3"})
	assert.Nil(t, err)
	assert.Equal(t, intVals, []int{1, 2, 3})
}

func TestGetSlicesStringToInt_Failure(t *testing.T) {
	intVals, err := GetSlicesStringToInt([]string{"1", "2xy", "3"})
	assert.Error(t, err)
	assert.Nil(t, intVals)

	intVals2, err := GetSlicesStringToInt([]string{"1", "2.0", "3"})
	assert.Error(t, err)
	assert.Nil(t, intVals2)
}

func TestGetPercentagesFromString_Success(t *testing.T) {
	percentageVals, err := GetPercentagesFromString([]string{"1.2% ", " 2.3%", "4.4%"})
	assert.Nil(t, err)
	assert.Equal(t, percentageVals, []float64{1.2, 2.3, 4.4})
}

func TestGetPercentagesFromString_Failure(t *testing.T) {
	_, err := GetPercentagesFromString([]string{"1.2% ", " 2.3%", "4.4"})
	assert.Error(t, err, errors.ErrInvalidCommandArguments)

	_, err = GetPercentagesFromString([]string{"1.2x% ", " 2.3%", "4.4%"})
	assert.Error(t, err, errors.ErrInvalidCommandArguments)

}
