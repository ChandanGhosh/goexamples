package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h *hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/dog":
		_, _ = fmt.Fprintln(w, "dog dog doggy")
	case "/cat":
		_, _ = fmt.Fprintln(w, "cat cat catty")

	}

}

func main() {
	var d *hotdog

	log.Panicln(http.ListenAndServe(":8080", d))
}
