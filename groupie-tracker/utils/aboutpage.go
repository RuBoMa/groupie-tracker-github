package utils

import (
	"log"
	"net/http"
)

func AboutPage(w http.ResponseWriter) {
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println("Error executing about.html: ", err)
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
