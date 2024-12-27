package main

import (
	xunitTests "github.com/jatin-malik/go-tdd/xUnit/tests"
)

func main() {
	test := xunitTests.NewTestCaseTest("TestTemplateMethod")
	test.Run(&test)

}
