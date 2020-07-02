package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	file, err1 := os.Create("index.html")
	if err1 != nil {
		log.Fatalln(err1)
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
