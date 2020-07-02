package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (h *hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	_ = tpl.ExecuteTemplate(w, "index.html", r.Form)
}
func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d *hotdog
	log.Fatalln(http.ListenAndServe(":8080", d))
}
