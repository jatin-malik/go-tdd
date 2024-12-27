package xunit_test

import (
	"fmt"
	"testing"

	xunit "github.com/jatin-malik/go-tdd/xUnit"
	"github.com/stretchr/testify/assert"
)

func TestTestMethodInvocation(t *testing.T) {
	test := xunit.NewWasRun("testMethod")
	fmt.Println(test.Flag)
	test.Run()
	assert.True(t, test.Flag)
}
