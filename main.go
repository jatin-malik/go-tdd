package main

import (
	xunitTests "github.com/jatin-malik/go-tdd/xUnit/tests"
)

func main() {

	// test suite for xunit
	test := xunitTests.NewTestCaseTest("TestRunning")
	test.Run(&test)

	test = xunitTests.NewTestCaseTest("TestSetup")
	test.Run(&test)

}
