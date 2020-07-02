package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greet string
}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v World!", mh.greet)))
}

func main() {
	// http.Handle("/", &myHandler{greet: "Hello"})
	http.ListenAndServe(":8000", &myHandler{"Hello"})
}
