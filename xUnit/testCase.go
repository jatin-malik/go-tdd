package xunit

import (
	"log"
	"reflect"
)

type Test interface {
	Run(Test, *TestResult) //TODO: signature looks weird?
	SetUp()
	TearDown()
	Suite() *TestSuite
}

// Base test struct that every other Test embeds.
type TestCase struct {
	Name string
}

func (t TestCase) SetUp() {

}

func (t TestCase) TearDown() {

}

func (t TestCase) Suite() *TestSuite {
	log.Println("suite not implemented in embedding test")
	return NewTestSuite()
}

func (t TestCase) Run(test Test, result *TestResult) {
	defer func() {
		if r := recover(); r != nil {
			test.TearDown()
			result.TestFailed()
			return
		}
	}()

	result.TestStarted()
	refValue := reflect.ValueOf(test)
	method := refValue.MethodByName(t.Name)
	if !method.IsValid() {
		result.TestFailed()
		return
	}

	test.SetUp()
	method.Call(nil)
	test.TearDown()
}
