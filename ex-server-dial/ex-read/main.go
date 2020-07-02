package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Println(err)
	}

	defer connection.Close()
	// make sure the ex-server/ex-server-tcp-write/main.go is running in
	// different terminal before running this code
	bs, err := ioutil.ReadAll(connection)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
