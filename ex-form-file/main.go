package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", fileUploadHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Panicln(http.ListenAndServe(":8080", nil))
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		fileContent, _, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer fileContent.Close()
		fs, err := ioutil.ReadAll(fileContent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		s = string(fs)

		// write to a file
		if err := ioutil.WriteFile("data.txt", fs, 0644); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

	w.Header().Set("content-type", "text/html; charset=utf-8")
	_, _ = io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q" />
		<input type="submit" />
	</form>
	<br/>`+s)
}
