package mymoney

import (
	"geektrust/command"
	"geektrust/output"
	"geektrust/portfolio"
	"geektrust/reader"
)

type Processor struct {
	console         output.Display
	filePathFetcher reader.FilePathFetcher
}

func GetProcessor() *Processor {
	return &Processor{
		console:         output.GetConsoleDisplay(),
		filePathFetcher: reader.GetConsoleFilePathFetcher(),
	}
}

func (p *Processor) Run() {
	commandStrings := p.getCommands()
	portfolio := portfolio.GetFreshPortfolio()
	p.executeCommands(commandStrings, portfolio)
}

func (p *Processor) getCommands() []string {
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

func (p *Processor) executeCommands(commandStrings []string, portfolio portfolio.Portfolio) {
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
