package utils

import (
	"net/http"
)

func AboutPage(w http.ResponseWriter) {
	if err := tmpl.ExecuteTemplate(w, "about.html", nil); err != nil {
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
