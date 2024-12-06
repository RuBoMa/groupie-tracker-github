package utils

import (
	"strings"
)

// Adding data to Concerts by looping locations (chronological) and matching the relations
// Concerts in chronologial order
func AddConcerts(artists []Band) {

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

// Reading locations from LocationURL and adding to Band struct by matching IDs
// Locations in chronologial order
func AddLocation(artists []Band, locationData LocationURL) {

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
func AddDates(artists []Band, datesData DatesURL) {

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
	var location string

	for i, char := range s {
		if char == '_' {
			location += " "
		} else if char == '-' {
			location += ", "
		} else if i == 0 || s[i-1] == '_' || s[i-1] == '-' {
			location += strings.ToUpper(string(char))
		} else {
			location += string(char)
		}
	}
	return location
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

func CleanInput(input string) string {
	cleanStr := ""

	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			cleanStr += string(char)
		}
	}
	return cleanStr
}

func AddLocationsToData(artists []Band, data *PageData) {
	data.AllLocations = append(data.AllLocations, "north_carolina-usa")

	for _, artist := range artists {
		for _, loc := range artist.Location {
			for index, prevLoc := range data.AllLocations {
				if loc == prevLoc {
					break
				}
				if index == len(data.AllLocations)-1 {
					data.AllLocations = append(data.AllLocations, loc)
				}
			}
		}
	}
}

func AddMaxMembersToData(artists []Band, data *PageData) {
	data.MaxMembers = append(data.MaxMembers, 1)
	for _, artist := range artists {
		for i := data.MaxMembers[len(data.MaxMembers)-1] + 1; i <= len(artist.Members); i++ {
			data.MaxMembers = append(data.MaxMembers, i)
		}
	}
}
