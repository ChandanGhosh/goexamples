package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h *hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Chandan", "Ghosh")
	w.Header().Set("content-type", "text/html; charset=utf-8;")
	_, _ = fmt.Fprintln(w, "<h1>This is an html output</h1>")
}

func main() {
	var d *hotdog
	log.Panicln(http.ListenAndServe(":8080", d))
}
