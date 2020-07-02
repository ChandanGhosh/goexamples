package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", doggy)
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doggy(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Println(request.URL.Path)

	_, _ = fmt.Fprintln(writer, "Doogy route")
}
