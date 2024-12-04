package utils

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func PageHandler(artists []Band) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {

		case "/":

			if r.Method != http.MethodGet {
				log.Println("Wrong user method requesting /")
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			if err := tmpl.ExecuteTemplate(w, "index.html", artists); err != nil {
				log.Println("Error executing index.html: ", err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}

		case "/About":

			if r.Method != http.MethodGet {
				log.Println("Wrong user method requesting /About")
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			AboutPage(w)

		default:

			if r.Method != http.MethodPost {
				log.Println("Wrong user method requesting band pages")
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			BandPage(artists, w, r)

		}
	})
}
