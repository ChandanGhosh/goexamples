package main

import (
	"fmt"
)

// Stringer interface
type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// implementing Stringer interface in fakeString struct
// with these changes, the fakeString type is now of type
// Stringer interface
func (fs *fakeString) String() string {
	return fs.content
}

// printString function proves that fakeString struct
// is now Stringer interface type as it implements
// String() method
func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println("string: " + str)
	case Stringer:
		fmt.Println("Stringer: " + str.String())
	case fakeString:
		fmt.Println("fakeString: " + str.content)
	}
}

func main() {
	s := &fakeString{content: "I am from fakeString struct"}
	printString(s)
	printString("Hello Gopher!!")
}
