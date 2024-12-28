package tests

import (
	"fmt"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
)

type TestCaseTest struct {
	xunit.TestCase
}

func NewTestCaseTest(name string) *TestCaseTest {
	return &TestCaseTest{TestCase: xunit.TestCase{Name: name}}
}

func (t *TestCaseTest) TestTemplateMethod() {
	wasRun := newWasRun("TestMethod")
	result := xunit.TestResult{}
	wasRun.Run(wasRun, &result)
	if wasRun.callLog != "setup TestMethod tearDown" {
		panic("testMethod invocation failed")
	}
}

func (t *TestCaseTest) TestReport() {
	wasRun := newWasRun("TestMethod")
	result := xunit.TestResult{}
	wasRun.Run(wasRun, &result)
	if result.Summary() != "1 run, 0 failed" {
		panic("reporting is not correct!")
	}
}

func (t *TestCaseTest) TestReportForFailedTests() {
	wasRun := newWasRun("")

	result := xunit.TestResult{}
	wasRun.Run(wasRun, &result)
	if result.Summary() != "1 run, 1 failed" {
		panic("reporting is not correct!")
	}
}

func (t *TestCaseTest) TestSuite() {
	suite := xunit.NewTestSuite()
	suite.Add(newWasRun("TestMethodBroken"))
	suite.Add(newWasRun("TestMethod"))
	result := xunit.TestResult{}
	suite.Run(&result)
	if result.Summary() != "2 run, 1 failed" {
		panic(fmt.Sprintf("got %s", result.Summary()))
	}
}

func (t *TestCaseTest) TestCreatingSuiteFromTestClass() {
	suite := newWasRun("").Suite()
	result := xunit.TestResult{}
	suite.Run(&result)
	if result.Summary() != "2 run, 1 failed" {
		panic(fmt.Sprintf("got %s", result.Summary()))
	}
}

func (t TestCaseTest) Suite() *xunit.TestSuite {
	suite := xunit.NewTestSuite()
	suite.Add(NewTestCaseTest("TestTemplateMethod"))
	suite.Add(NewTestCaseTest("TestReport"))
	suite.Add(NewTestCaseTest("TestReportForFailedTests"))
	suite.Add(NewTestCaseTest("TestSuite"))
	suite.Add(NewTestCaseTest("TestCreatingSuiteFromTestClass"))
	return suite
}
