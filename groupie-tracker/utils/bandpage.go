package utils

import (
	"net/http"
	"strings"
	"text/template"
)

var tmplArtist = template.Must(template.ParseFiles("templates/artist.html"))

func BandPage(artists []Band, data PageData, w http.ResponseWriter, r *http.Request) {
	artistExists := false
	for _, artist := range artists {
		//	Case insensitive word matching
		if strings.EqualFold(artist.Name, r.URL.Path[1:]) {
			err := tmplArtist.Execute(w, artist)
			if err != nil {
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}
			artistExists = true
		}
	}
	if !artistExists {
		ErrorPage(w, "Page not found", http.StatusNotFound)
	}
}

