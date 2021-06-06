package main

import (
	"fmt"
)

// “A closure is a nested function that
// has access to the variables of its parent function,
// even after the parent has executed.”
// so essentially the anonymous function
// gets a new copy of the nested function's
// state variable i and increment it and
// nested function store it in memory until
// the program exists
// This is why everytime there is a call to
// nested function would return
// a newer incremented value of i
// as every call to nested function
// would automatically execute the anonymous function
// inside it
func nested() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}
}

func main() {
	myfunc := nested()
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
}
