package xunit

import "fmt"

type WasRun struct {
	Name           string
	Flag           bool // flag that indicates whether the testMethod was invoked
	methodRegistry map[string]func()
}

func NewWasRun(name string) *WasRun {
	wasRun := WasRun{Name: name}

	// Keep this updated. Could have used reflection but unexported methods are a pain!
	wasRun.methodRegistry = map[string]func(){
		"Run":        wasRun.Run,
		"testMethod": wasRun.testMethod,
	}
	return &wasRun
}

func (w *WasRun) Run() {
	if method, ok := w.methodRegistry[w.Name]; !ok {
		fmt.Printf("method: %s not found", w.Name)
		return
	} else {
		method()
	}
}

func (w *WasRun) testMethod() {
	w.Flag = true
}
