package tests

import xunit "github.com/jatin-malik/go-tdd/xUnit"

type wasRun struct {
	xunit.TestCase
	callLog string
}

func newWasRun(name string) wasRun {
	return wasRun{TestCase: xunit.TestCase{Name: name}}
}

func (w *wasRun) TestMethod() {
	w.callLog += " TestMethod"
}

func (w *wasRun) SetUp() {
	w.callLog = "setup"
}

func (w *wasRun) TearDown() {
	w.callLog += " tearDown"
}
