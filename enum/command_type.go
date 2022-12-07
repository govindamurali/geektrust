package enum

import (
	"geektrust/errors"
	"strings"
)

//todo think of moving
type CommandType string

const (
	AllocateCommand    CommandType = "allocate"
	BalanceCommand     CommandType = "balance"
	ChangeCommand      CommandType = "change"
	RebalanceCommand   CommandType = "rebalance"
	SipCommand         CommandType = "sip"
	unsupportedCommand CommandType = "unsupported"
)

var commandTypes = map[string]CommandType{
	"allocate":  AllocateCommand,
	"balance":   BalanceCommand,
	"change":    ChangeCommand,
	"rebalance": RebalanceCommand,
	"sip":       SipCommand,
}

func (c CommandType) GetArgsCount() int {
	switch c {
	case RebalanceCommand:
		return 1
	case BalanceCommand:
		return 2
	case AllocateCommand, SipCommand:
		return 4
	case ChangeCommand:
		return 5
	}
	return 0
}

func GetCommandType(s string) (CommandType, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if c, ok := commandTypes[s]; ok {
		return c, nil
	}

	return unsupportedCommand, errors.ErrInvalidCommand
}
