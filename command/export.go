package command

import "geektrust/portfolio"

type iCommand interface {
	Execute(portfolio portfolio.Portfolio) error
}

func GetCommandResolver() *commandResolver {
	return &commandResolver{}
}
