package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{"Buddha", "The belief of no beliefs"}
	gandhi := sage{"Gandhi", "Be the change"}
	mlk := sage{"Martin Luthar King", "Hatred never wins"}
	jesus := sage{"Jesus", "Love is the only beautiful thing in the world!"}
	muhammad := sage{"Muhammad", "Insah Allah!"}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	if err := tmpl.Execute(os.Stdout, sages); err != nil {
		log.Fatalln(err)
	}

}
