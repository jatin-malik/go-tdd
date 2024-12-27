package tests

import xunit "github.com/jatin-malik/go-tdd/xUnit"

type wasRun struct {
	xunit.TestCase
	runFlag   bool // flag that indicates whether the testMethod was invoked
	setupFlag bool // flag that indicates whether the setUp was done
}

func newWasRun(name string) wasRun {
	return wasRun{TestCase: xunit.TestCase{Name: name}}
}

func (w *wasRun) TestMethod() {
	w.runFlag = true
}

func (w *wasRun) SetUp() {
	w.setupFlag = true
}
