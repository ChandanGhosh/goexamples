package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Quote string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	data := sage{
		"Buddha",
		"Love everyone, everything",
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
