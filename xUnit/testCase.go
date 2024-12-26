package xunit

import (
	"log"
	"reflect"
)

type WasRun struct {
	Name string
	Flag bool // flag that indicates whether the testMethod was invoked
}

func (w *WasRun) Run() {
	refValue := reflect.ValueOf(w)
	method := refValue.MethodByName(w.Name)
	if !method.IsValid() {
		log.Printf("Method: %s not found\n", w.Name)
		return
	}
	method.Call(nil)
}

func (w *WasRun) TestMethod() {
	w.Flag = true
}
