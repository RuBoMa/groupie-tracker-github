package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("*.html"))

func main() {

	log.Println("fetching data")
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locationData)
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &datesData)
	if err != nil {
		log.Fatalf("Error fetching dates s: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}

	// for _, artist := range artists {
	// 	fmt.Printf("Artist: %s, Creation Date: %d\n", artist.Name, artist.CreationDate)
	// }

	log.Println("adding data to the artist")
	addLocation()
	addDates()
	addRelations()
	addConcerts()

	// Serving static files like CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", home)

	fmt.Println("Server started on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		log.Println("Error executing index.html: ", err)
		http.Error(w, "Internal Server Error, please try again later.", http.StatusInternalServerError)
		return
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

// request-respond page still missing. POST action is needed in some form. Maybe artist page?
