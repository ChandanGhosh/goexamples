package main

import (
	"fmt"
	"math"
	"reflect"
)

type square struct {
	height float32
	width  float32
}

func (sq square) area() float32 {
	return sq.height * sq.width
}

type circle struct {
	radius float32
}

func (circle circle) area() float32 {
	return 2 * math.Pi * circle.radius
}

type shape interface {
	area() float32
}

func calculateArea(sh shape) {
	typeStr := reflect.TypeOf(sh)
	fmt.Println("shape type is", typeStr, `and its area is`, sh.area())
}

func main() {
	sq := square{
		4,
		3,
	}

	ci := circle{
		5,
	}

	calculateArea(sq)
	calculateArea(ci)
}
