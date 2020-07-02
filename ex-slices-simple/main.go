package main

import "fmt"

func main() {
	// simple slice
	p := []string{"Hello", "World", "Again", "!"}

	fmt.Printf("%q", p)

	// you can slice a slice with underlying array intact
	s1 := p[1:3]
	fmt.Printf("\n%q", s1)

	// omit the first include the first
	s2 := p[:2]
	fmt.Printf("\n%q", s2)

	// omit the last part
	s3 := p[2:]
	fmt.Printf("\n%q", s3)
}
