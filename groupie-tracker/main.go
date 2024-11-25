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

var tmpl = template.Must(template.ParseFiles("index.html"))
var bands []Band

func main() {
	// Serving static files like CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	// Map data in response
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	// Ensures the response body is closed after using it to avoid ressource leaks
	defer response.Body.Close()

	// Read data with io.ReadAll
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse JSON into bands slice
	json.Unmarshal(responseData, &bands)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server started on http://localhost:8090")
	http.ListenAndServe(":8090", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	err := tmpl.ExecuteTemplate(w, "index.html", bands)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
