package errors

import "errors"

var (
	PortfolioAlreadyAllocatedError = errors.New("portfolio already allocated")
	NotAllocatedError              = errors.New("initial allocation not done")
	InvalidChangeMonth             = errors.New("change to be calculated only for consecutive months. input data invalid")
	InvalidCommandError            = errors.New("invalid command")
	InvalidCommandArgumentsError   = errors.New("invalid command arguments")
)
