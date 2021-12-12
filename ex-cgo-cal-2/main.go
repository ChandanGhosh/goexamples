// In ex-cgo-cal-1 we learned to call a basic c function from directly within
// go program. Now here we shall learn how to write a c file with the implementation
// create a header for it and then create a object using gcc compiler : gcc -c -o cal.o cal.c
// and then make a library archive using "ar" binary from object file(cal.o): ar -rsc libcal.a cal.o
// so that go can link it at the time of creating the go binary.
// Because when Go compiles it, it can compile based on the signature from header file
// But when it creates the binary it needs the actual c implementation to get the function symbol to link
// so that it can create a binary. So here we have added the sum.h header file to provide
// signature for compiling but also provided cgo pragma linker with LDFLAGS and location with -L
// to the current directory and then specify the library to use using -l<library_name>.
// It will pickup libcal.a from current directory and link it to execute the sum function
// when running "go run main.go"

package main

// #cgo LDFLAGS: -L . -lcal
// #include "cal.h"
import "C"

func main() {
	C.sum(10, 15)
}
