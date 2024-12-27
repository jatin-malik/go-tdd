package tests

import (
	"log"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
)

type TestCaseTest struct {
	xunit.TestCase
}

func NewTestCaseTest(name string) TestCaseTest {
	return TestCaseTest{TestCase: xunit.TestCase{Name: name}}
}

func (t *TestCaseTest) TestTemplateMethod() {
	wasRun := newWasRun("TestMethod")
	wasRun.Run(&wasRun)
	if wasRun.callLog != "setup TestMethod tearDown" {
		log.Fatal("testMethod invocation failed")
	}
	log.Println("TestTemplateMethod ran successfully")
}
