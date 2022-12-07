package reader

import (
	"bufio"
	"fmt"
	"os"
)

func GetStrings(filePath string) (inputs []string, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	fmt.Println("File opened")

	inputs = make([]string, 0)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		args := scanner.Text()
		fmt.Println(args)
		inputs = append(inputs, args)
	}
	return
}
