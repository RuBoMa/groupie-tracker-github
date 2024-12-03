package utils

import (
	"log"
	"net/http"
)

func AboutPage(w http.ResponseWriter) {
	if err := tmpl.ExecuteTemplate(w, "about.html", nil); err != nil {
		log.Println("Error executing about.html: ", err)
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
