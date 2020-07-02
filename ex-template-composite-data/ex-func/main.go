package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tmpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}

func init() {
	tmpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{
		"Buddha",
		"Gandhi",
		"Martin Luther King",
	}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
