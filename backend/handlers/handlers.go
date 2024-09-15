package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"groupie-trackers/backend/api"
	"groupie-trackers/backend/function"
	"groupie-trackers/backend/models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := api.GitAllData()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error\nError 500")
		return
	}
	temp, err := template.ParseFiles("../../frontend/templates/HomePage.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error\nError 500")
		return
	}
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, "The page does not exist\nError 404")
		return
	}
	if r.Method != "GET" {
		handleError(w, http.StatusMethodNotAllowed, "The method is not allowed\nError 405")
		return
	}
	search := r.FormValue("search")
	creationYearFrom := r.FormValue("creation-date-from")
	creationYearTo := r.FormValue("creation-date-to")

	albumYearFrom := r.FormValue("album-date-from")
	albumYearTo := r.FormValue("album-date-to")

	members := r.Form["members"]
	location := r.FormValue("location")

	if len(members) != 0 || location != "" || creationYearFrom != "" || creationYearTo != "" || albumYearFrom != "" || albumYearTo != "" {
		models.Var_Artists = function.Filter(creationYearFrom, creationYearTo, albumYearFrom, albumYearTo, location, members)
	} else if search != "" {
		models.Var_Artists = function.Search(search)
	}
	data := models.PageData{
		Artists:   models.Var_Artists,
		Locations: models.Var_Locations,
	}
	errE := temp.Execute(w, data)
	if errE != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error\nError 500")
		return
	}
}

func Artists(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../../frontend/templates/More.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error\nError 500")
		return
	}
	if r.Method != "GET" {
		handleError(w, http.StatusMethodNotAllowed, "The method is not allowed\nError 405")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, 400, "Invalid artist ID")
		return
	}
	// Check if artist ID is valid
	if id < 1 || id > 52 {
		handleError(w, 404, "Artist not found")
		return
	}

	var locations []string
	var coordinates []models.GeoLocation

	// Find the locations for the specific artist
	for _, loc := range models.Var_Locations.Index {
		if loc.ID == id {
			locations = loc.Locations
			break
		}
	}

	// Geocode each location
	for _, location := range locations {
		coord, err := api.Geocode(location)
		if err != nil {
			fmt.Println("Error geocoding %s: %v\n", location, err)
			continue
		}
		coordinates = append(coordinates, *coord)
	}

	data := models.PageData{
		Artists:     models.Var_Artists,
		Locations:   models.Var_Locations,
		Dates:       models.Var_Dates,
		Relation:    models.Var_Relation,
		ArtistID:    id,
		Coordinates: coordinates,
	}

	errE := temp.Execute(w, data)
	if errE != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error\nError 500")
		return
	}
}

func handleError(w http.ResponseWriter, StatusCode int, message string) {
	temp, err := template.ParseFiles("../../frontend/templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error\nError 500", 500)
		return
	}

	w.WriteHeader(StatusCode)

	models.Err.Number = StatusCode
	models.Err.Message = message

	errE := temp.Execute(w, models.Err)
	if errE != nil {
		http.Error(w, "Internal Server Error\nError 500", 500)
		return
	}
}

func ServeStyle(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/static")))

	if r.URL.Path == "/static/" {
		handleError(w, http.StatusNotFound, "The page does not exist\nError 404")
		return
	}

	fs.ServeHTTP(w, r)
}
