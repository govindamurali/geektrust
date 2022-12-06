package command

import (
	"geektrust/enum"
	errors2 "geektrust/errors"
	"geektrust/output"
	"strings"
)

type CommandResolver struct {
	validator
}

const commandSeparator = " "

func (c *CommandResolver) GetCommand(s string, display output.IDisplay) (iCommand, error) {

	// todo add basic validations

	commandArgs := strings.Split(s, commandSeparator)

	if len(commandArgs) == 0 {
		return nil, errors2.ErrInvalidCommand
	}

	commandType, err := enum.GetCommandType(commandArgs[0])
	if err != nil {
		return nil, err
	}

	err = c.validator.validateCount(commandArgs, commandType)
	if err != nil {
		return nil, err
	}

	switch commandType {
	case enum.AllocateCommand:
		return getAllocateCommand(commandArgs)
	case enum.BalanceCommand:
		getBalanceCommand(commandArgs, display)
	case enum.ChangeCommand:
		getChangeCommand(commandArgs)
	case enum.RebalanceCommand:
		getRebalanceCommand(display)
	case enum.SipCommand:
		return getSipCommand(commandArgs)
	}
	return nil, errors2.ErrInvalidCommand
}

func GetCommandResolver() *CommandResolver {
	return &CommandResolver{}
}
