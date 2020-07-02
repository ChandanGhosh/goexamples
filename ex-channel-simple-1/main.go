package main

import "fmt"

func main() {

	// making a unbuffered channel of int
	c := make(chan int)

	// then abruptly close the channel without sending
	// any values to the channel
	close(c)

	// then trying to see what is in the channel
	// The channel will return two things
	// the first one is data which is in this case
	// default int i.e. 0 and the next one is
	// any further data awaiting to be sent in the channel which is
	// referred bellow as more. Every channel when closed, will
	// signal about this. In this case it is false
	// usually it is the best practice to close the channel from the sender
	// so that the sender can decide whether it has data or not
	// and receiver can use the more flag to see if the
	// channel is not closed before asking for data from channel
	// this also avoid blocking/race conditions as well when multiple receivers receives
	// the data
	data, more := <-c
	fmt.Println(data, more)
}
