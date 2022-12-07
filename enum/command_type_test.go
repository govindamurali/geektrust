package enum

import (
	"geektrust/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandType_GetArgsCount_Valid(t *testing.T) {
	com := AllocateCommand
	assert.Equal(t, 4, com.GetArgsCount())
}
func TestCommandType_GetArgsCount_Invalid(t *testing.T) {
	comInvalid := CommandType("random")
	assert.Equal(t, 0, comInvalid.GetArgsCount())
}

func TestGetCommandType_Valid(t *testing.T) {
	com, err := GetCommandTypeFromString("Allocate ")
	assert.Nil(t, err)
	assert.Equal(t, com, AllocateCommand)
}
func TestGetCommandType_Invalid(t *testing.T) {
	com, err := GetCommandTypeFromString("random")
	assert.Error(t, err, errors.ErrInvalidCommand)
	assert.Equal(t, com, unsupportedCommand)
}
