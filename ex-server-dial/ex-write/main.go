package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panicln(err)
	}
	defer connection.Close()

	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(connection, "I dial you %s", "dear")
}
