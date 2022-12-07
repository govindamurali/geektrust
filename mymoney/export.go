package mymoney

import (
	"geektrust/output"
	"geektrust/reader"
)

func GetProcessor() *processor {
	return &processor{
		console:         output.GetConsoleDisplay(),
		filePathFetcher: reader.GetConsoleFilePathFetcher(),
	}
}
