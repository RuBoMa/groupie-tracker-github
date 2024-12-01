package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomePage(artists []Band, data PageData) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search" && r.URL.Path != "/" {
			BandPage(artists, data, w, r)
			return
		} else if r.URL.Path == "/search" {
			SearchPage(artists, data, w, r)
			return
		}
		if r.Method != http.MethodGet {
			ErrorPage(w, "Wrong user method", http.StatusMethodNotAllowed)
			return
		} else {
			tmplHome, err := template.ParseFiles("templates/index.html")
			if err != nil {
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if err := tmplHome.Execute(w, data); err != nil {
				fmt.Println(err)
				ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	})

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
