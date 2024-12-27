package xunit

import (
	"log"
)

type TestCaseTest struct {
	TestCase
}

func (t *TestCaseTest) RunTest() {
	test := WasRun{TestCase: TestCase{Name: "TestMethod"}}
	if test.Flag {
		log.Fatal("flag should be false")
	}
	test.Run(&test)
	if !test.Flag {
		log.Fatal("flag should be true")
	}
}
