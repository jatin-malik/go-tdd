package xunit

import "fmt"

type TestResult struct {
	totalRan int
	failed   int
}

func (t *TestResult) TestStarted() {
	t.totalRan += 1
}

func (t *TestResult) TestFailed() {
	t.failed += 1
}

func (t TestResult) Summary() string {
	return fmt.Sprintf("%d run, %d failed", t.totalRan, t.failed)
}
