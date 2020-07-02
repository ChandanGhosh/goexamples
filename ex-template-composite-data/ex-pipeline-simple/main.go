package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

func sqr(n int) float64 {
	return math.Pow(float64(n), float64(2))
}

func dbl(n int) int {
	return n + n
}

func sqrt(n int) float64 {
	return math.Sqrt(float64(2))
}

var fm = template.FuncMap{
	"sqr":  sqr,
	"sqrt": sqrt,
	"dbl":  dbl,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.html"))
}

func main() {
	if err := tpl.ExecuteTemplate(os.Stdout, "tmpl.html", 3); err != nil {
		log.Fatalln(err)
	}
}
