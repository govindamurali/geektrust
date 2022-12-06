package output

import "fmt"

type consoleDisplay struct{}

func (c *consoleDisplay) Display(val string) {
	fmt.Println(val)
}
