package xunit

import (
	"log"
	"reflect"
)

type Test interface {
	Run(Test) //TODO: signature looks weird?
	SetUp()
}

// Base test struct that every other Test embeds.
type TestCase struct {
	Name string
}

func (t TestCase) SetUp() {

}

func (t TestCase) Run(test Test) {
	refValue := reflect.ValueOf(test)
	method := refValue.MethodByName(t.Name)
	if !method.IsValid() {
		log.Printf("Method: %s not found\n", t.Name)
		return
	}

	test.SetUp()
	method.Call(nil)
}
