package enum

import (
	"geektrust/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMonthFromString_Valid(t *testing.T) {
	month, err := GetMonthFromString("JANUARY ")
	assert.Nil(t, err)
	assert.Equal(t, January, month)
}
func TestGetMonthFromString_Invalid(t *testing.T) {
	month, err := GetMonthFromString("Nomonth ")
	assert.Error(t, err, errors.ErrInvalidMonth)
	assert.Equal(t, InvalidMonth, month)
}
