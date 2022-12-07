package enum

import (
	"geektrust/errors"
	"strings"
)

type Month int

const (
	InvalidMonth Month = iota
	January
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var months = map[string]Month{
	"january":   January,
	"february":  February,
	"march":     March,
	"april":     April,
	"may":       May,
	"june":      June,
	"july":      July,
	"august":    August,
	"september": September,
	"october":   October,
	"november":  November,
	"december":  December,
}

func (m Month) String() string {
	return [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}[m-1]
}

func GetMonthFromString(s string) (Month, error) {
	if month, ok := months[strings.ToLower(strings.TrimSpace(s))]; ok {
		return month, nil
	}
	return InvalidMonth, errors.ErrInvalidMonth
}

func (m Month) IsRebalanceMonth() bool {
	return m == June || m == December
}
