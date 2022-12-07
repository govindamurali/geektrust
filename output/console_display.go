package output

import "fmt"

type consoleDisplay struct{}

func (c *consoleDisplay) Output(val string) {
	fmt.Println(val)
}

func GetConsoleDisplay() Display {
	return &consoleDisplay{}
}
