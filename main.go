package main

import "geektrust/mymoney"

func main() {
	processor := mymoney.GetProcessor()
	processor.Run()
}
