package utils

import (
	"log"
	"net/http"
	"text/template"
)

func PageHandler(artists []Band, data PageData) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {

		case "/":
			if r.Method != http.MethodGet {
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			tmplHome, err := template.ParseFiles("templates/index.html")
			if err != nil {
				log.Println(err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if err := tmplHome.Execute(w, data); err != nil {
				log.Println(err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		case "/Filter":
			if r.Method != http.MethodPost {
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			FilterPage(artists, data, w, r)

		case "/About":
			if r.Method != http.MethodGet {
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			AboutPage(w)

		default:
			if r.Method != http.MethodGet {
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			BandPage(artists, data, w, r)

		}
	})
}
