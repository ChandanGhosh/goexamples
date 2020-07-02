package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template
var sessionUsers = map[string]string{}
var dbUsers = map[string]user{}

type user struct {
	username string
	password string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/signin", signin)

	log.Panicln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		fmt.Fprintln(w, "User logged in")
	} else {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		password := r.FormValue("password")

		sid := uuid.NewV4()

		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		})

		sessionUsers[sid.String()] = un

		dbUsers[un] = user{username: un, password: password}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		_ = tpl.ExecuteTemplate(w, "signin.html", nil)
	}

}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	var sid = cookie.Value
	if _, ok := sessionUsers[sid]; ok {
		return ok
	}
	return false
}
