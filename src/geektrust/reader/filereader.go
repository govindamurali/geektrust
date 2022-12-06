package reader

import (
	"bufio"
	"fmt"
	"os"
)

func GetStrings(filePath string) (inputs []string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	fmt.Println("opened")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fmt.Println("scanning")
		args := scanner.Text()
		//argList := strings.Fields(args)

		fmt.Println(args)

	}
	return
}
