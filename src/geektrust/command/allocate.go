package command

import "geektrust/dal"

type allocate struct {
	amount map[string]int
}

func Execute(portfolio dal.Portfolio) {

	allocation :=
		portfolio.Allocate()
}
