package utils

import (
	"net/http"
	"text/template"
)

var tmplLocations = template.Must(template.ParseFiles("templates/locations.html"))

func LocationsPage(w http.ResponseWriter) {
	if err := tmplLocations.Execute(w, nil); err != nil {
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
