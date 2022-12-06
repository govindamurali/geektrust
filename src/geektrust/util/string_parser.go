package util

import (
	"strconv"
	"strings"
)

const percentageSymbol = "%"

func GetSlicesStringToInt(stringValues []string) (intValues []int, err error) {

	intValues = make([]int, len(stringValues))

	for i := range stringValues {
		intValues[i], err = strconv.Atoi(strings.TrimSpace(stringValues[i]))
		if err != nil {
			return
		}
	}

	return
}

func GetPercentagesFromString(stringValues []string) (percentages []float64, err error) {

	percentages = make([]float64, len(stringValues))

	for i := range stringValues {
		percentageVal, err := strconv.ParseFloat(strings.TrimRight(strings.TrimSpace(stringValues[i]), percentageSymbol), 64)
		if err != nil {
			return percentages, err
		}
		percentages[i] = percentageVal
	}

	return
}
