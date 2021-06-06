package main

import (
	"fmt"
	"time"
)

func main() {

	var ch chan int

	select {
	case m := <-ch:
		// this is blocking forever
		// as there is nothing to receive from the channel
		fmt.Println(m)
	case <-time.After(time.Second * 10):
		// But in a different case here we are using
		// different channel offered by time API
		// to simulate a message to be sent to this case
		// and the select exists. This is very
		// nice way of blocking something in an asynchronous thread safe way
		// Since the select chooses which ever case completes first,
		// hence we get exits finally from this select block
		// Channel with select offers extremely powerful combination for Go
		fmt.Println("Timeout is over. Exiting..")
	}

}
