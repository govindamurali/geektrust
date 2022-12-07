package mymoney

import (
	"geektrust/command"
	"geektrust/output"
	"geektrust/portfolio"
	"geektrust/reader"
)

type processor struct {
	console         output.Display
	filePathFetcher reader.FilePathFetcher
}

func (p *processor) Run() {
	portfolio := portfolio.GetFreshPortfolio()
	p.executeCommands(p.getCommands(), portfolio)
}

func (p *processor) getCommands() []string {
	var commandStrings []string
	var err error
	for commandStrings == nil {
		filePath := p.filePathFetcher.GetFilePath()
		commandStrings, err = reader.GetStrings(filePath)
		if err != nil {
			p.console.Output(err.Error())
			p.console.Output("Try again")
		}
	}
	return commandStrings
}

func (p *processor) executeCommands(commandStrings []string, portfolio portfolio.Portfolio) {
	commandResolver := command.GetCommandResolver()
	for _, val := range commandStrings {
		command, err := commandResolver.GetCommand(val, p.console)
		if err != nil {
			p.console.Output(err.Error())
		}
		err = command.Execute(portfolio)
		if err != nil {
			p.console.Output(err.Error())
		}
	}
}
