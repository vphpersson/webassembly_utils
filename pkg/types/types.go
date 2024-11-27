package types

import "syscall/js"

type Promise struct {
	Resolve js.Value
	Reject  js.Value

	JsPromise js.Value
}

func MakePromise() *Promise {
	var promise Promise

	promiseConstructor := js.Global().Get("Promise")
	promise.JsPromise = promiseConstructor.New(
		js.FuncOf(
			func(this js.Value, args []js.Value) any {
				promise.Resolve = args[0]
				promise.Reject = args[1]
				return nil
			},
		),
	)

	return &promise
}
