package main

import (
	"geektrust/command"
	"geektrust/output"
	"geektrust/portfolio"
	"geektrust/reader"
)

func main() {

	filePath := reader.GetFilePath()
	comms := reader.GetStrings(filePath)
	portfolio := portfolio.GetEmptyPortfolio()
	commandResolver := command.GetCommandResolver()
	outputMode := output.GetConsoleDisplay()
	for _, val := range comms {
		command, err := commandResolver.GetCommand(val, outputMode)
		if err != nil {
			outputMode.Display(err.Error())
		}
		err = command.Execute(portfolio)
		if err != nil {
			outputMode.Display(err.Error())
		}
	}
}
