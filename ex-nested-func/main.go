package main

import (
	"fmt"
)

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
