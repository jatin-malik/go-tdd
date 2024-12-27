package main

import xunit "github.com/jatin-malik/go-tdd/xUnit"

func main() {

	// test suite for xunit
	test := xunit.TestCaseTest{TestCase: xunit.TestCase{"RunTest"}}
	test.Run(&test)
}
