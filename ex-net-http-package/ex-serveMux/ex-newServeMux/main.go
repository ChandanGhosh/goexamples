package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog string

func (h *hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/cat":
		_, _ = fmt.Fprintln(w, "cat cat catty")

	case "/dog":
		_, _ = fmt.Fprintln(w, "dog dog doggy")
	default:
		_, _ = fmt.Fprintln(w, "Hello world!")

	}
}

func main() {

	var d *hotdog

	mux := http.NewServeMux()

	mux.Handle("/cat", d)
	mux.Handle("/dog", d)

	log.Panicln(http.ListenAndServe(":8080", mux))
}
