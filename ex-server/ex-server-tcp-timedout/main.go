package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}

	defer listener.Close()

	connection, err := listener.Accept()
	err = connection.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(connection, "Ready to accept connection from %s\n", "you!!")
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "Received: %s\n", line)
	}

	defer connection.Close()

	fmt.Println("Code reaches here this time because of deadline a.k.a timeout is set")
}
