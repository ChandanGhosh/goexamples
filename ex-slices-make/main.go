package main

import "fmt"

func main() {
	// create a slice using make. this creates a nil slice
	p := make([]string, 3)

	fmt.Println("Hello", p)

}
