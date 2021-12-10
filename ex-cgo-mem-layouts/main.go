// Go has ability to work even more lowlevel with C for certain situations
// like graphics program, driver access which are only written in C
// while this should be avoided as much as possible(https://dave.cheney.net/2016/01/18/cgo-is-not-go)
// but sometimes this is the only option left to work directly with kernels and drivers

// Here we shall inspect memory layouts of the built-in types as well as user-defined types using
// Cgo because when it comes to addresses and memory allocations, go has a nice package called
// "unsafe" to track and find very interesting details of the programs. It provides access to
// a number of language features which are not ordinarily available because they exposes
// Go's memory layouts.

package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Order of struct fields can impact on memory
// allocation as well. Like the below one takes
// 5 words = 40bytes (in 64bit) os.
type more_memory struct {
	a bool
	c []int
	b int16
}

// Order of the struct fields should be
// ascending or descending in terms of
// memory allocation by fields
// to optimize on memory allocation of
// struct. Here this struct takes 4words = 32byte of memory(64bit os)
type less_memory struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println("OS type:", runtime.GOARCH)
	// unsafe has built-in methods like unsafe.Sizeof, unsafe.Allignof, unsafe.Offsetof
	// Printing float64 of 0 show how much memory Go allocates - 8bytes
	fmt.Printf("Sizeof(float64(0)): %vbytes\n", unsafe.Sizeof(float64(0)))
	// Printing int of 0 shows how much memory Go allocates - 8bytes
	// in 64bit OS int represents int64. 8bytes is also called 1word in Go's terminalogy and in 32bit system 4bytes = 1word
	// So, int allocates 1word of memory where as float64 uses 2words of memory
	fmt.Printf("sizeof(int(0)): %vbytes\n", unsafe.Sizeof(int(0)))

	var x more_memory
	// Struct memory layout
	fmt.Printf("Sizeof(x): %vbytes\n", unsafe.Sizeof(x))
	fmt.Printf("Sizeof(x.a): %vbytes\tAllignof(x.a): %v\tOffsetof(x.a): %v\n", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("Sizeof(x.b): %vbytes\tAllignof(x.b): %v\tOffsetof(x.b): %v\n", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	// For array it is not only the values(or data) the array store, it also takes space for
	// len and capacity. 8bytes - data(int64), 8bytes - len(int64), 8bytes - cap(int64)
	fmt.Printf("Sizeof(x.c): %vbytes\tAllignof(x.c): %v\tOffsetof(x.c): %v\n", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))

	var y less_memory
	// Struct memory layout
	fmt.Printf("Sizeof(y): %vbytes\n", unsafe.Sizeof(y))
	fmt.Printf("Sizeof(y.a): %vbytes\tAllignof(y.a): %v\tOffsetof(y.a): %v\n", unsafe.Sizeof(y.a), unsafe.Alignof(y.a), unsafe.Offsetof(y.a))
	fmt.Printf("Sizeof(y.b): %vbytes\tAllignof(y.b): %v\tOffsetof(y.b): %v\n", unsafe.Sizeof(y.b), unsafe.Alignof(y.b), unsafe.Offsetof(y.b))
	// For array it is not only the values(or data) the array store, it also takes space for
	// len and capacity. 8bytes - data(int64), 8bytes - len(int64), 8bytes - cap(int64)
	fmt.Printf("Sizeof(y.c): %vbytes\tAllignof(y.c): %v\tOffsetof(y.c): %v\n", unsafe.Sizeof(y.c), unsafe.Alignof(y.c), unsafe.Offsetof(y.c))

}
