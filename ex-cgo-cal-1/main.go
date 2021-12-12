// Here we are going implement a simple sum
// to understand a very simple execution of C program from
// go.

package main

// #include <stdio.h>
//
// void sum(int a, int b) {
//	 printf("The sum of %d and %d is %d\n", a, b, a+b);
// }
import "C" // This is a special import. Though there is no package as "C" but it tells compiler that we shall use "C" language here.

func main() {
	C.sum(4, 3)
}
