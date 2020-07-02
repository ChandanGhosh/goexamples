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

	connection, err := listener.Accept()

	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintf(connection, "ready to accept connection from %s\n", "you")
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "I hear you said: %s\n", line)
	}

	defer connection.Close()
}
