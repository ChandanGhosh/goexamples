package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(c chan int) {
	val := rand.Intn(10)
	fmt.Println("Calculated random value is : ", val)
	time.Sleep(1 * time.Second)
	c <- val
	fmt.Println("This line will only be shown for buffered channel")
}

func main() {
	c := make(chan int, 3)
	defer close(c)
	go sendData(c)
	go sendData(c)
	res := <-c
	fmt.Println(res)
}
