package utils

import (
	"strings"
)

// Reading locations from LocationURL and adding to Band struct by matching IDs
func AddLocation(artists []Band, locationData LocationURL) {

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, loc := range locationData.Index {
			var cleanLoc []string
			if loc.ID == artists[i].ID {
				for _, place := range loc.Locations {
					addLoc := cleanLocation(place)
					cleanLoc = append(cleanLoc, addLoc)
				}
				artists[i].Location = cleanLoc
				//artists[i].Location = loc.Locations
				break
			}
		}
	}
}

// Reading dates from DatesURL and adding to Band struct by matching IDs
// Dates in chronologial order
func AddDates(artists []Band, datesData DatesURL) {

	for i := range artists {
		// Find the corresponding location data for the artist based on ID
		for _, date := range datesData.Index {
			if date.ID == artists[i].ID {
				cleanDates := cleanDates(date.Dates)
				//not changing source material
				artists[i].Dates = cleanDates
				break
			}
		}
	}
}

// Reading relations from RelationsURL and adding to Band struct by matching IDs
// Concerts in alphabetical order after the town name
func AddRelations(artists []Band, relations RelationsURL) {

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

// Adding data to Concerts by looping the relations
// Concerts include same data as relations but the places are capitalized
func AddConcerts(artists []Band) {

	for i := range artists {
		concertDates := make(map[string][]string)
		for place, dates := range artists[i].Relation {
			newLoc := cleanLocation(place)
			concertDates[newLoc] = dates
		}
		artists[i].Concerts = concertDates
	}
}

// Capitalizes first letter of the city and country (separated by "-"), adding space instead of "_"
func cleanLocation(loc string) string {
	location := strings.Split(loc, "-")

	for i := 0; i < len(location); i++ {
		item := location[i]
		if item == "usa" || item == "uk" {
			location[i] = strings.ToUpper(item)
		} else {
			modLoc := ""
			for j, char := range item {
				if char == '_' {
					modLoc += " "
				} else if j == 0 || item[j-1] == '_' || item[j-1] == '-' {
					modLoc += strings.ToUpper(string(char))
				} else {
					modLoc += string(char)
				}
			}
			location[i] = modLoc
		}
	}
	return strings.Join(location, ", ")
}

// Removes the "*" from the beginning for the dates
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
