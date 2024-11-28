package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Band structure of the JSON data
type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     []string
	Dates        DatesURL
	Relation     RelationsURL
}

type LocationURL struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type DatesURL struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type RelationsURL struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	var artists []Band
	var locationData LocationURL
	var dates DatesURL
	var relations RelationsURL

	log.Println("fething artist data")
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	log.Println("fething location data")
	err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locationData)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	log.Println("fething dates data")
	err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	log.Println("fething releations data")
	err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}

	// for _, artist := range artists {
	// 	fmt.Printf("Artist: %s, Creation Date: %d\n", artist.Name, artist.CreationDate)
	// }

	addLocation(artists, locationData)

	// Serving static files like CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Println("rendering BandPAge")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, artists)
	})

	fmt.Println("Server started on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))

}

// Reading locations from LocationURL and adding to Band struct by matching IDs
func addLocation(artists []Band, locationData LocationURL) {

	cleanLocation := func(s []string) []string {
		var location []string

		for _, word := range s {
			modWord := ""
			for i, char := range word {
				if char == '_' {
					modWord += " "
				} else if char == '-' {
					modWord += " - "
				} else if i == 0 || word[i-1] == '_' || word[i-1] == '-' {
					modWord += strings.ToUpper(string(char))
				} else {
					modWord += string(char)
				}
			}
			location = append(location, modWord)
		}
		return location
	}

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, loc := range locationData.Index {
			if loc.ID == artists[i].ID {
				cleanLocation := cleanLocation(loc.Locations)
				artists[i].Location = cleanLocation
				break
			}
		}
	}
}

func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(data)) // Debug the API response
	return json.Unmarshal(data, target)
}


// possible changes: Reading locations into a 