package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetFilePath() string {
	reader := bufio.NewReader(os.Stdin)

	var filePath string
	var err error

	fmt.Println("Please provide the file location")

	for {
		fmt.Print("-> ")
		filePath, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please provide valid string as input", err)
			continue
		}
		// convert CRLF to LF
		filePath = strings.Replace(filePath, "\n", "", -1)
		break
	}

	return filePath
}
