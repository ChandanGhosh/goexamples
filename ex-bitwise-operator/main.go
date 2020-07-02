package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1 << 1 | 1 << 5  - This says have a uint8 binary representation of 1 (i.e. 00000001) and then shift truthy bit one position from right at first. Which will be 00000010
	// Then with the result, shift truthy bit another 5 steps from the right again which will be 00100000.
	// Then OR both of the values(i.e. 00000010 and 00100000). So now the result would be 00100010
	// This is now the value of x. So x = 00100010
	var x = 1<<1 | 1<<5

	// So here y finally would be 00000110
	var y = 1<<1 | 1<<2

	// now print them in binary nice 8bit padding format
	fmt.Printf("x : %08b \n", x)
	fmt.Printf("y : %08b \n", y)

	// print all the logic gate operations, supporte by golang for this two numbers x, y
	// bitwise AND (&)
	fmt.Printf("x&y : %08b \n", x&y)

	// bitwise OR (|)
	fmt.Printf("x|y :%08b \n", x|y)

	// bitwise XOR(^)
	fmt.Printf("x^y : %08b \n", x^y)

	// bitwise AND NOT (&^)
	fmt.Printf("x&^y : %08b \n", x&^y)

	// test the membership sets. For example x belongs to {1, 5} as the values of x floats between 1 -5 bit positions {00000010 - 00100000}

	s := "{"
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			s = fmt.Sprintf("%s%v, ", s, i)
		}
	}
	s = strings.Trim(s, ", ") + "}"
	fmt.Println("The set of x is: ", s)

	// let's have so more fun with bitwise shift from right
	// x << 1 - here we are shifting the x truthy bits by one again from right. That means x (00100010) becomes 01000100. The set also shifts by one accordingly.
	// Now the sets become {1, 5} to {2, 6}
	fmt.Printf("x<<1 : %08b \n", x<<1)

	// let's have so more fun with bitwise shift from left
	// x >> 1 - here we are shifting the x's truthy bits by one again from left. That means x (00100010) becomes 00010001. The set also shifts by one accordingly.
	// Now the sets become {1, 5} to {0, 4}
	fmt.Printf("x>>1 : %08b \n", x>>1)
}
