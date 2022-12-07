package reader

import (
	"bufio"
	"os"
)

func GetStrings(filePath string) (inputs []string, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	inputs = make([]string, 0)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		args := scanner.Text()
		inputs = append(inputs, args)
	}
	return
}
