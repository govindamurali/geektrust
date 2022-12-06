package enum

type PortfolioType int

const (
	Equity PortfolioType = iota + 1
	Debt
	Gold
)

func (p PortfolioType) String() string {
	return [...]string{"Equity", "Debt", "Gold"}[p-1]
}
