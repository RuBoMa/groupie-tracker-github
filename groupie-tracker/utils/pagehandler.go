package utils

import (
	"fmt"
	"net/http"
)

func PageHandler(artists []Band, data PageData) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {

		case "/":

			if r.Method != http.MethodGet {
				ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
				return
			}
			if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
				fmt.Println(err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}

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
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
