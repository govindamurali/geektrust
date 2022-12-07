package enum

import (
	"geektrust/errors"
	"strings"
)

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

var commandTypeArgumentCountMap = map[CommandType]int{
	RebalanceCommand: 1,
	BalanceCommand:   2,
	AllocateCommand:  4,
	SipCommand:       4,
	ChangeCommand:    5,
}

func (c CommandType) GetArgsCount() int {
	return commandTypeArgumentCountMap[c]
}

func GetCommandTypeFromString(s string) (CommandType, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if c, ok := commandTypes[s]; ok {
		return c, nil
	}

	return unsupportedCommand, errors.ErrInvalidCommand
}
