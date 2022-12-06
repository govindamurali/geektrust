package command

import (
	"errors"
	"strings"
)

type CommandResolver struct {
}

func (c *CommandResolver) GetCommand(s string) (iCommand, error) {

	commandArgs := strings.Split(s, " ")

	switch commandArgs[0] {
	case sipCommand:
		return

	default:
		return nil, errors.New("invalid command")
	}
}
