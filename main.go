package main

import (
	"fmt"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
	xunitTests "github.com/jatin-malik/go-tdd/xUnit/tests"
)

func main() {
	suite := xunitTests.NewTestCaseTest("").Suite()
	result := xunit.TestResult{}
	suite.Run(&result)
	fmt.Println(result.Summary())
}
