package main

import (
	"fmt"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
)

func main() {
	test := xunit.NewWasRun("testMethod")
	fmt.Println(test.Flag)
	test.Run()
	fmt.Println(test.Flag)
}
