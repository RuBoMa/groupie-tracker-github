package utils

import (
	"log"
	"net/http"
	"strings"
)

func BandPage(artists []Band, w http.ResponseWriter, r *http.Request) {
	artistExists := false
	name := r.FormValue("bandName")

	if name != r.URL.Path[1:] {
		log.Println("Error: Post not matching URL")
		ErrorPage(w, "Bad Request", http.StatusBadRequest)
		return
	}
	for _, artist := range artists {
		//	Case insensitive word matching
		if strings.EqualFold(artist.Name, name) {
			
			err := tmpl.ExecuteTemplate(w, "artist.html", artist)
			if err != nil {
				log.Println("Error executing artist.html: ", err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}
			artistExists = true
		}
	}
	if !artistExists {
		log.Println("Error: artist page not found.")
		ErrorPage(w, "Page not found", http.StatusNotFound)
	}
	
}
