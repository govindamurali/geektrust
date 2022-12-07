package errors

import "errors"

var (
	ErrPortfolioAlreadyAllocated = errors.New("portfolio already allocated")
	ErrNotAllocated              = errors.New("initial allocation not done")
	ErrInvalidChangeMonth        = errors.New("change to be calculated only for consecutive months. input data invalid")
	ErrInvalidCommand            = errors.New("invalid command")
	ErrInvalidCommandArguments   = errors.New("invalid command arguments")
	ErrInvalidMonth              = errors.New("invalid month parameter")
)
