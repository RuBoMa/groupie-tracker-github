package utils

import (
	"fmt"
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
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if err := tmplHome.Execute(w, data); err != nil {
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
			//	Change to method post when changing request/response stuff
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
