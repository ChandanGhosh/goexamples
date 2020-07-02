package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cat", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Cat cat catty")
	})
	http.HandleFunc("/dog/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "dog dog doggy")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
