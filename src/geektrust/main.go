package main

import (
	"fmt"
	"geektrust/reader"
)

func main() {

	filePath := reader.GetFilePath()
	comms := reader.GetStrings(filePath)
	fmt.Println(comms)

}
