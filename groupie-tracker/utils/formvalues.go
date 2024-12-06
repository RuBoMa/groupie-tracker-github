package utils

import (
	"net/http"
	"strconv"
)

func GetIntFormValues(value string, r *http.Request) (int, error) {

	data := r.FormValue(value)
	if data == "" {
		return -1, nil
	}

	//	Convert formvalues

	output, err := strconv.Atoi(data)
	if err != nil {
		return -1, err
	}
	return output, nil
}
