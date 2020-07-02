package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Panicln(err)
		}
	}()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handle(connection)
	}

}

func handle(conn io.ReadWriteCloser) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Panicln(err)
		}
	}()

	_, err := io.WriteString(conn,
		"Welcome to the database\n"+
			"Use GET key \n"+
			"SET key value \n"+
			"DEL key\n")
	if err != nil {
		log.Println("Not able to write to the connection")
	}

	scanner := bufio.NewScanner(conn)
	data := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			words := strings.Fields(line)

			if len(words) < 1 {
				continue
			}

			switch words[0] {
			case "GET":
				key := words[1]
				val := data[key]
				_, _ = fmt.Fprintf(conn, "%s\r\n", val)

			case "SET":
				key := words[1]
				val := words[2]
				data[key] = val
			case "DEL":
				key := words[1]
				delete(data, key)
				_, _ = fmt.Fprintf(conn, "DELETED %s\n", key)
			default:
				_, _ = fmt.Fprintf(conn, "INVALID COMMAND %s\n", words)
				continue
			}

		}
	}
	fmt.Println("Connection closed")
}
