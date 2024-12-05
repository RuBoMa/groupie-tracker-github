Web site for displaying RESTful API

# Groupie Trackers

## Description:

In all simpleness, Groupie Trackers is a web application that consists on consuming given RESTful APIs and displaying all information on a website.

## Authors:

Fanni, Johannes & Roope

## Features:

Displaying artist details
Displaying concerts, as well as all locations and dates the artists have had concerts separately
Data visualization in form of cards and pages
Client-server interaction when requesting an artists page

## Usage:

1. Start the Server: by running the following command in your terminal: go run main.go

2. Open in Browser: In your browser, navigate to http://localhost:8080

3. Scroll around and learn more about the bands

- Get to the band pages by clicking the pictures/names
- Read about the authors and the project by clicking About Us

## Implemention details:

API Parsing: fetchData gets the response from the API, reads the response and then parse JSON encoded data and store it into the target interface (pointer of a struct). All four API endpoints are parsed separately and the data stored to their own structs.

Data Linking: Dates, locations and relations are all added to the artists variable in separate functions (eg. AddLocation) by matching the ID from the Band struct to the one on Index struct.

Client-Server Communication: Clicking an artist name or image on the home page, takes the client to the artist page. BandPage function searches for a match from the artists variable and if found, renders the artist.html template with the artist data.

Visualization: On both home and artist page cards are used to display artists/artists data. The data visualisation in form of cards makes it easy to browse artists on the home page. Collapsibles on the artist page give the client a choice to view selected information and have fun discovering facts about the artist.

Error Handling: Since the web application only allows GET method, any other method results on an error and no input validation is needed. Trying to get any other than root, About or a valid artist page returns page not found error. All error are logged on terminal.
