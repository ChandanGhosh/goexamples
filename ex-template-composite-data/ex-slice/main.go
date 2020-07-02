package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := []string{"Gandhi", "MLK", "Buddha", "Christ"}

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
