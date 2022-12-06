package enum

type Month int

const (
	January Month = iota + 1
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

func (m Month) String() string {
	return [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}[m-1]
}

func (m Month) isRebalanceRequired() bool {
	return m == June || m == December
}
