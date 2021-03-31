package channel

import (
	"reflect"
	"sync"
)

func pipe(out reflect.Value, in reflect.Value) {
	for {
		if item, ok := in.Recv(); ok {
			out.Send(item)
		} else {
			break
		}
	}
}

func multiplex(out reflect.Value, in reflect.Value) {
	var wait sync.WaitGroup

	for i := 0; i < in.Len(); i++ {
		wait.Add(1)

		go func(to reflect.Value, from reflect.Value) {
			defer wait.Done()

			pipe(to, from)
		}(out, in.Index(i))
	}

	wait.Wait()

	out.Close()
}

// Multiplex redirects multiple-input streams to single-output stream.
func Multiplex(out interface{}, in interface{}) {
	go multiplex(reflect.ValueOf(out), reflect.ValueOf(in))
}
