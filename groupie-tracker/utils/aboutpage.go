package utils

import (
	"log"
	"net/http"
	"text/template"
)

var tmplAbout = template.Must(template.ParseFiles("templates/about.html"))

func AboutPage(w http.ResponseWriter) {
	if err := tmplAbout.Execute(w, nil); err != nil {
		log.Println(err)
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
