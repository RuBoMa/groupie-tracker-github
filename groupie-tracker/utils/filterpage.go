package utils

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tmplFilter = template.Must(template.ParseFiles("templates/filter.html"))

func FilterPage(artists []Band, data PageData, w http.ResponseWriter, r *http.Request) {
	//	Get formvalues

	memberAmount, err := GetIntFormValues("member-amount", r)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Bad User Request", http.StatusBadRequest)
		return
	}

	creationYearMin, err := GetIntFormValues("creation-year-min", r)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Bad User Request", http.StatusBadRequest)
		return
	}
	creationYearMax, err := GetIntFormValues("creation-year-max", r)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Bad User Request", http.StatusBadRequest)
		return
	}

	albumYearMin, err := GetIntFormValues("album-year-min", r)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Bad User Request", http.StatusBadRequest)
		return
	}
	albumYearMax, err := GetIntFormValues("album-year-max", r)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Bad User Request", http.StatusBadRequest)
		return
	}

	location := r.FormValue("location")

	var matchingArtists []Band

	for _, artist := range artists {
		artistAlbumYear, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
		if err != nil {
			log.Println(err)
			ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if memberAmount == -1 || (len(artist.Members) == memberAmount) {
			if creationYearMax == -1 || (artist.CreationDate >= creationYearMin && artist.CreationDate <= creationYearMax) {
				if albumYearMax == -1 || (artistAlbumYear >= albumYearMin && artistAlbumYear <= albumYearMax) {
					if location != "" {
						for _, loc := range artist.Location {
							if loc == location {
								matchingArtists = append(matchingArtists, artist)
								break
							}
						}
					} else {
						matchingArtists = append(matchingArtists, artist)
					}
				}
			}
		}
	}
	data.Filter.MemberAmount = memberAmount
	if creationYearMin != -1 && creationYearMax != -1 {
		data.Filter.CreationYearMin = creationYearMin
		data.Filter.CreationYearMax = creationYearMax
	}
	if albumYearMin != -1 && albumYearMax != -1 {
		data.Filter.AlbumYearMin = albumYearMin
		data.Filter.AlbumYearMax = albumYearMax
	}
	data.Filter.Location = location
	data.Matches = matchingArtists
	err = tmplFilter.Execute(w, data)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
