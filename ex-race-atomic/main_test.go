package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	n := 3
	var sum int32

	for i := 0; i < n; i++ {
		go func(i int) {
			// this will cause a race condition because the sum is written by multiple
			// go routines. run with go test -race ./...
			sum += int32(i)

			// To avoid the race condition above we should use channel to share the data by communicating.
			// Or, atleast use atomic package like below. So to test, comment the line number 15 and uncomment line 19 and
			// run go test -race ./...
			// atomic.AddInt32(&sum, int32(i))

		}(i)
	}

}
