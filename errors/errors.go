package errors

import (
	"errors"
	"strings"
)

var (
	ErrPortfolioAlreadyAllocated = errors.New("portfolio already allocated")
	ErrPortfolioNotAllocated     = errors.New("portfolio not allocated")
	ErrNotAllocated              = errors.New("initial allocation not done")
	ErrInvalidChangeMonth        = errors.New("change to be calculated only for consecutive months. input data invalid")
	ErrInvalidCommand            = errors.New("invalid command")
	ErrInvalidCommandArguments   = errors.New("invalid command arguments")
	ErrInvalidMonth              = errors.New("invalid month parameter")
)

func GetAppendedErrors(errs ...error) error {
	if len(errs) == 0 {
		return errs[0]
	}
	sb := strings.Builder{}
	for _, err := range errs {
		sb.WriteString(err.Error() + " ")
	}
	return errors.New(sb.String())
}
