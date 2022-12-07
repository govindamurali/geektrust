package mymoney

import (
	mocks2 "geektrust/output/mocks"
	"geektrust/reader/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var filePathwithRebalance = "./test_files/input_with_rebalance.txt"
var filePathwithoutRebalance = "./test_files/input_without_rebalance.txt"

func TestProcessor_Run_Rebalanced(t *testing.T) {

	mockDisplay := mocks2.NewDisplay(t)
	mockFilePathFetcher := mocks.NewFilePathFetcher(t)

	printedValues := make([]string, 0)

	mockDisplay.On("Output", mock.AnythingOfType("string")).Return().Run(
		func(args mock.Arguments) {

			arg := args.Get(0).(string)
			printedValues = append(printedValues, arg)

		})
	mockFilePathFetcher.On("GetFilePath").Return(filePathwithRebalance)

	processor := Processor{
		console:         mockDisplay,
		filePathFetcher: mockFilePathFetcher,
	}

	processor.Run()

	assert.Equal(t, printedValues, []string{"10593 7897 2272", "23619 11809 3936"})
}

func TestProcessor_Run_NonRebalanced(t *testing.T) {

	mockDisplay := mocks2.NewDisplay(t)
	mockFilePathFetcher := mocks.NewFilePathFetcher(t)

	printedValues := make([]string, 0)

	mockDisplay.On("Output", mock.AnythingOfType("string")).Return().Run(
		func(args mock.Arguments) {

			arg := args.Get(0).(string)
			printedValues = append(printedValues, arg)

		})
	mockFilePathFetcher.On("GetFilePath").Return(filePathwithoutRebalance)

	processor := Processor{
		console:         mockDisplay,
		filePathFetcher: mockFilePathFetcher,
	}

	processor.Run()

	assert.Equal(t, printedValues, []string{"15937 14552 6187", "23292 16055 7690", "CANNOT_REBALANCE"})
}
