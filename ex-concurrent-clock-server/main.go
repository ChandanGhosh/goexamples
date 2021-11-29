package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Concurrent clock server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Unable to listen to port 8080, %s", err)
	}
	defer listener.Close()

	fmt.Println("Clock server started at :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept ant tcp connection", err)

		}
		log.Println("Client connected", conn.RemoteAddr().String())
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var err error
	for {
		now := time.Now()
		formattedTime := now.Format("15:04:05\n")
		_, err = io.WriteString(conn, formattedTime)
		if err != nil {
			fmt.Printf("Error occured %s\nClient disconnected %s", err, conn.RemoteAddr().String())
			break
		}
		time.Sleep(1 * time.Second)
	}
}
