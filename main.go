package main

import (
	"geektrust/command"
	"geektrust/output"
	"geektrust/portfolio"
	"geektrust/reader"
)

func main() {

	outputMode := output.GetConsoleDisplay()
	filePath := reader.GetFilePath()
	comms, err := reader.GetStrings(filePath)
	if err != nil {
		outputMode.Output(err.Error())
	}
	portfolio := portfolio.GetFreshPortfolio()
	commandResolver := command.GetCommandResolver()
	for _, val := range comms {
		command, err := commandResolver.GetCommand(val, outputMode)
		if err != nil {
			outputMode.Output(err.Error())
		}
		err = command.Execute(portfolio)
		if err != nil {
			outputMode.Output(err.Error())
		}
	}
}
