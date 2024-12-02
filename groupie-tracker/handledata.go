package main

import (
	"strings"
)

// Adding data to Concerts by looping locations (chronological) and matching the relations
// Capitalized concert data
func addConcerts() {

	for i := range artists {
		concertDates := make(map[string][]string)
		for _, loc := range artists[i].Location {
			for place, dates := range artists[i].Relation {
				if loc == place {
					newLoc := cleanLocation(place)
					concertDates[newLoc] = dates
				}
			}
		}
		artists[i].Concerts = concertDates
	}
}

// Reading relations from RelationsURL and adding to Band struct by matching IDs
// Concerts in alphabetical order after the town name
func addRelations() {

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, rel := range relations.Index {
			if rel.ID == artists[i].ID {
				artists[i].Relation = rel.DatesLocations
				break
			}
		}

	}
}

// Reading locations from LocationURL and adding to Band struct by matching IDs
// Locations in chronologial order
func addLocation() {

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, loc := range locationData.Index {
			if loc.ID == artists[i].ID {
				// cleanLocation := cleanLocation(loc.Locations)
				// artists[i].Location = cleanLocation
				artists[i].Location = loc.Locations
				break
			}
		}
	}
}

// Reading dates from DatesURL and adding to Band struct by matching IDs
// Dates in chronologial order
func addDates() {

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, date := range datesData.Index {
			if date.ID == artists[i].ID {
				cleanDates := cleanDates(date.Dates)
				artists[i].Dates = cleanDates
				break
			}
		}
	}
}

func cleanLocation(s string) string {
	location := strings.Split(s, "-")

	for i := 0; i < len(location); i++ {
		word := location[i]
		if word == "usa" {
			location[i] = strings.ToUpper(word)
		} else {
			newLoc := ""
			for j, char := range word {
				if char == '_' {
					newLoc += " "
				} else if j == 0 || word[j-1] == '_' {
					newLoc += strings.ToUpper(string(char))
				} else {
					newLoc += string(char)
				}
			}
			location[i] = newLoc
		}
	}
	return strings.Join(location, ", ")
}

func cleanDates(s []string) []string {
	var dates []string

	for _, date := range s {
		modWord := ""
		for _, char := range date {
			if char != '*' {
				modWord += string(char)
			}
		}
		dates = append(dates, modWord)
	}
	return dates
}
