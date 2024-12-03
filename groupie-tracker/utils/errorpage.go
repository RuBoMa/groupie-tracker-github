package utils

import (
	"net/http"
	"strconv"
	"text/template"
)

var tmplError = template.Must(template.ParseFiles("templates/error.html"))

func ErrorPage(w http.ResponseWriter, errorMessage string, errorStatus int) {
	w.WriteHeader(errorStatus)
	data := PageData{
		ErrorMessage: errorMessage,
		ErrorStatus:  "Error " + strconv.Itoa(errorStatus),
	}
	tmplError.Execute(w, data)
}
