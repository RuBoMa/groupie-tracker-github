package main

import (
	"encoding/json"
	"fmt"
	"grp/utils"
	"io"
	"log"
	"net/http"
)

func main() {
	var artists []utils.Band
	var locationData utils.LocationURL
	var dates utils.DatesURL
	var relations utils.RelationsURL

	log.Println("fetching data")
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locationData)
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		log.Fatalf("Error fetching dates s: %v", err)
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}

	log.Println("adding data to artists variable")
	utils.AddLocation(artists, locationData)
	utils.AddDates(artists, dates)
	utils.AddRelations(artists, relations)
	utils.AddConcerts(artists)

	data := utils.PageData{
		Matches: artists,
	}

	// Serving static files like CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Println("rendering PageHandler")
	utils.PageHandler(artists, data)
	fmt.Println("Server started on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))

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
