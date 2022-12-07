package command

import (
	"geektrust/enum"
	errors2 "geektrust/errors"
	"geektrust/output"
	"strings"
)

type commandResolver struct{}

const commandSeparator = " "

func (c *commandResolver) GetCommand(s string, display output.Display) (iCommand, error) {

	// todo add basic validations

	commandArgs := strings.Split(s, commandSeparator)

	if len(commandArgs) == 0 {
		return nil, errors2.ErrInvalidCommand
	}

	commandType, err := enum.GetCommandType(commandArgs[0])
	if err != nil {
		return nil, err
	}

	err = c.validateCount(commandArgs, commandType)
	if err != nil {
		return nil, err
	}

	switch commandType {
	case enum.AllocateCommand:
		return getAllocateCommand(commandArgs)
	case enum.BalanceCommand:
		return getBalanceCommand(commandArgs, display)
	case enum.ChangeCommand:
		return getChangeCommand(commandArgs)
	case enum.RebalanceCommand:
		return getRebalanceCommand(display)
	case enum.SipCommand:
		return getSipCommand(commandArgs)
	}
	return nil, errors2.ErrInvalidCommand
}

func (c *commandResolver) validateCount(args []string, commandType enum.CommandType) error {
	if len(args) != commandType.GetArgsCount() {
		return errors2.ErrInvalidCommandArguments
	}
	return nil
}
