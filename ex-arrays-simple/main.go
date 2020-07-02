package main

import "fmt"

func main() {
	// declaring an array
	var a [2]string
	// setting array values
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// another way
	b := [2]string{"Hello", "World"}
	fmt.Println(b[0], b[1])
	fmt.Println(b)

	// Using ellipsis to declare implicit array declaration
	c := [...]string{"Hello", "World", "again", ":)"}
	fmt.Printf("%q", c)
	fmt.Println(c)
}
