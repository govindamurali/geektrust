package enum

type PortfolioType int

const (
	Unsupported PortfolioType = iota
	Equity
	Debt
	Gold
)

var portfolios = map[string]PortfolioType{
	"equity": Equity,
	"debt":   Debt,
	"gold":   Gold,
}
