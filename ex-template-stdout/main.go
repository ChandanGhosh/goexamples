package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Fatalln("error parsing file", err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("error in outputting the file content")
	}
}
