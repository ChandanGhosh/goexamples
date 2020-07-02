package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Fprintf(connection, "An error occured in creating a connection: %s", err)
		}

		io.WriteString(connection, "\nYour connection to this tcp server is successful!\n")
		io.WriteString(connection, "Good day!, bye!!\n")
		connection.Close()
	}
}
