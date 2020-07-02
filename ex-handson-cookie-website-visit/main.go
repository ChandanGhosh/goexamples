package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", readCookies)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func readCookies(writer http.ResponseWriter, request *http.Request) {
	c1, err := request.Cookie("visited")
	if err != nil {
		_, _ = fmt.Fprintln(writer, "Could not retrieve cookie details", err)
	}
	_, _ = fmt.Fprintln(writer, "The site visited:", c1.Value)
}

func set(writer http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("visited")
	if err != nil {
		http.SetCookie(writer, &http.Cookie{Name: "visited", Value: "1"})
		_, _ = fmt.Println("Cookie not set yet.", err)
		_, _ = fmt.Fprintln(writer, "Visitor cookie is set with value")
		return
	}

	val, _ := strconv.Atoi(c1.Value)
	c1.Value = strconv.Itoa(val + 1)
	http.SetCookie(writer, c1)
	_, _ = fmt.Fprintln(writer, "Visitor cookie value is", c1.Value)

}
