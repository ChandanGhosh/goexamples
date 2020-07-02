package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.txt"))
}

func main() {
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "file1.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "file2.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "file3.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
