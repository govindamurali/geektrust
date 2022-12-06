package command

import "geektrust/dal"

const (
	allocateCommand  = "allocate"
	balanceCommand   = "balance"
	changeCommand    = "change"
	rebalanceCommand = "rebalance"
	sipCommand       = "sip"
)

type iCommand interface {
	Execute(portfolio dal.Portfolio)
}

type iDisplay interface {
	Display(val string)
}
