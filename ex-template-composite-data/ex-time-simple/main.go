package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var tpl *template.Template

var fm = template.FuncMap{
	"monthDayYear": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tmpl.html", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
