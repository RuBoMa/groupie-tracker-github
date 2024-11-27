package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Band structure of the JSON data
type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	var artists []Band
	var locations []Location
	var dates []Date
	var relations []Relation

	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}

	for _, artist := range artists {
		fmt.Printf("Artist: %s, Creation Date: %d\n", artist.Name, artist.CreationDate)
	}
	// Serving static files like CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", artists)
	})

	fmt.Println("Server started on http://localhost:8090")
	http.ListenAndServe(":8090", nil)

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
	return json.Unmarshal(data, target)
}
