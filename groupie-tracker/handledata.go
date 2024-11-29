package main

import "strings"

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
					modWord += ", "
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

// Reading locations from LocationURL and adding to Band struct by matching IDs
func addDates(artists []Band, datesData DatesURL) {

	cleanDates := func(s []string) []string {
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
