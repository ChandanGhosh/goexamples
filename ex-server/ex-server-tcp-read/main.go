package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Fprintf(connection, "An error occured in creating a connection: %s", err)
		}

		scanner := bufio.NewScanner(connection)

		for scanner.Scan() {
			receivedLine := scanner.Text()
			fmt.Println(receivedLine)
		}

		connection.Close()

		fmt.Println("Bye!!")
	}
}
