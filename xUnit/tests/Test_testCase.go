package tests

import (
	"log"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
)

type TestCaseTest struct {
	xunit.TestCase
	wasRunFixture wasRun
}

func NewTestCaseTest(name string) TestCaseTest {
	return TestCaseTest{TestCase: xunit.TestCase{Name: name}}
}

func (t *TestCaseTest) SetUp() {
	t.wasRunFixture = newWasRun("TestMethod")
}

func (t *TestCaseTest) TestRunning() {
	t.wasRunFixture.Run(&t.wasRunFixture)
	if !t.wasRunFixture.runFlag {
		log.Fatal("run flag should be true")
	}
	log.Println("TestRunning ran successfully")
}

func (t *TestCaseTest) TestSetup() {
	t.wasRunFixture.Run(&t.wasRunFixture)
	if !t.wasRunFixture.setupFlag {
		log.Fatal("setUp flag should be true")
	}
	log.Println("TestSetup ran successfully")
}
