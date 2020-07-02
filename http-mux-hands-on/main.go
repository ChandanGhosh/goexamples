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
		if err := listener.Close(); err != nil {
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
		if err := conn.Close(); err != nil {
			log.Panicln(err)
		}
	}()

	request(conn)

	respond(conn)
}

func request(conn io.ReadWriteCloser) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line-", i, ":- ", line)
		if i == 0 {
			words := strings.Fields(line)
			fmt.Println("***METHOD*** :", words[0])
			fmt.Println("***URI*** :", words[1])

			// Now the multiplexer(router) kind of naive way :)
			respMux(conn, words[1])
		}
		if line == "" {
			break
		}
		i++

	}

}

func respMux(conn io.ReadWriteCloser, s string) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` +
		s +
		`</strong></body></html>`

	_, _ = fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	_, _ = fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	_, _ = fmt.Fprint(conn, "Content-Type: text/html\r\n")
	_, _ = fmt.Fprint(conn, "\r\n")
	_, _ = fmt.Fprint(conn, body)
}

func respond(conn io.ReadWriteCloser) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	_, _ = fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	_, _ = fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	_, _ = fmt.Fprint(conn, "Content-Type: text/html\r\n")
	_, _ = fmt.Fprint(conn, "\r\n")
	_, _ = fmt.Fprint(conn, body)
}
