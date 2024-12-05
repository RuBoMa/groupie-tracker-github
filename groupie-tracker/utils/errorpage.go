package utils

import (
	"net/http"
	"strconv"
)

func ErrorPage(w http.ResponseWriter, errorMessage string, errorStatus int) {
	w.WriteHeader(errorStatus)
	data := PageData{
		ErrorMessage: errorMessage,
		ErrorStatus:  "Error " + strconv.Itoa(errorStatus),
	}
	err := tmpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, errorMessage, errorStatus)
	}
}
