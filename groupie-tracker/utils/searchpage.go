package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func SearchPage(artists []Band, data PageData, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
	} else {
		var matches []Band
		data.Input = CleanInput(r.FormValue("search"))
		fmt.Println(data.Input)
		for _, group := range artists {
			if strings.Contains(strings.ToLower(group.Name), strings.ToLower(data.Input)) {
				matches = append(matches, group)
				continue
			}
			for _, member := range group.Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(data.Input)) {
					matches = append(matches, group)
					break
				}
			}
		}
		data.Matches = matches
		searchtmpl, err := template.ParseFiles("templates/search.html")
		if err != nil {
			ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err = searchtmpl.Execute(w, data)
		if err != nil {
			ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

}
