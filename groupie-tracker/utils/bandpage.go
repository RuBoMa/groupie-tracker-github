package utils

import (
	"log"
	"net/http"
	"strings"
)

func BandPage(artists []Band, w http.ResponseWriter, r *http.Request) {
	for _, artist := range artists {
		//	Case insensitive word matching
		if strings.EqualFold(artist.Name, r.URL.Path[1:]) {
			err := tmpl.ExecuteTemplate(w, "artist.html", artist)
			if err != nil {
				log.Println("Error executing artist.html: ", err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
	}
	log.Println("Error: artist page not found.")
	ErrorPage(w, "Page not found", http.StatusNotFound)
}
