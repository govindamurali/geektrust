package command

import (
	"geektrust/enum"
	"geektrust/errors"
)

type validator struct{}

func (v *validator) validateCount(args []string, commandType enum.CommandType) error {
	if len(args) != commandType.GetArgsCount() {
		return errors.ErrInvalidCommandArguments
	}
	return nil
}
