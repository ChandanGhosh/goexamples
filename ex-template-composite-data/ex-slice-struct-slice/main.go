package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name string
}

type wisdom struct {
	Title string
}

type items struct {
	S []sage
	W []wisdom
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []sage{
		sage{"Gandhi"},
		sage{"MLK"},
		sage{"Buddha"},
	}

	wisdoms := []wisdom{
		wisdom{"More you know, more you realize you don't know"},
		wisdom{"Love is the best thing in this world"},
		wisdom{"The Earth is the beatiful thing for us"},
	}

	data := items{
		sages,
		wisdoms,
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}

}
