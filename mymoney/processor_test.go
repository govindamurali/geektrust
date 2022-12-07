package mymoney

import (
	mocks2 "geektrust/output/mocks"
	"geektrust/reader/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProcessor_Run_Rebalanced(t *testing.T) {

	mockDisplay := mocks2.NewDisplay(t)
	mockFilePathFetcher := mocks.NewFilePathFetcher(t)

	printedValues := make([]string, 0)

	mockDisplay.On("Output", mock.AnythingOfType("string")).Return().Run(
		func(args mock.Arguments) {

			arg := args.Get(0).(string)
			printedValues = append(printedValues, arg)

		})
	mockFilePathFetcher.On("GetFilePath").Return("./test_files/input1.txt")

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
	mockFilePathFetcher.On("GetFilePath").Return("./test_files/input2.txt")

	processor := Processor{
		console:         mockDisplay,
		filePathFetcher: mockFilePathFetcher,
	}

	processor.Run()

	assert.Equal(t, printedValues, []string{"15937 14552 6187", "23292 16055 7690", "CANNOT_REBALANCE"})
}
