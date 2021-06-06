package main

import "fmt"

func main() {
	var ch chan string = make(chan string)

	go func() {
		msg := <-ch
		fmt.Println("waiting for messages")
		fmt.Println(msg)
		ch <- "pong"
	}()
	ch <- "ping"
	fmt.Println(<-ch)
}
