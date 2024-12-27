package xunit

import (
	"log"
	"reflect"
)

type TestCase struct {
	Name string
}

func (t TestCase) Run(any interface{}) {
	refValue := reflect.ValueOf(any)
	method := refValue.MethodByName(t.Name)
	if !method.IsValid() {
		log.Printf("Method: %s not found\n", t.Name)
		return
	}
	method.Call(nil)
}

type WasRun struct {
	TestCase
	Flag bool // flag that indicates whether the testMethod was invoked
}

func (w *WasRun) TestMethod() {
	w.Flag = true
}
