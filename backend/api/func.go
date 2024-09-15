package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"groupie-trackers/backend/models"
)

func GetData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching URL: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}

func GitAllData() error {
	url := "https://groupietrackers.herokuapp.com/api"

	if err := GetData(url, &models.Var_ApiURL); err != nil {
		return err
	}

	if err := GetData(models.Var_ApiURL.Artists, &models.Var_Artists); err != nil {
		return err
	}

	if err := GetData(models.Var_ApiURL.Locations, &models.Var_Locations); err != nil {
		return err
	}

	if err := GetData(models.Var_ApiURL.Dates, &models.Var_Dates); err != nil {
		return err
	}

	if err := GetData(models.Var_ApiURL.Relation, &models.Var_Relation); err != nil {
		return err
	}

	return nil
}

// Geocode uses Google Maps Geocoding API to get the latitude and longitude for a given address.
func Geocode(address string) (*models.GeoLocation, error) {
	baseURL, _ := url.Parse("https://maps.googleapis.com/maps/api/geocode/json")

	// creating a RowUrl to be used as the Api to extract the info of the locations
	params := url.Values{}
	params.Add("address", address)
	params.Add("key", "key google map")

	baseURL.RawQuery = params.Encode()
	urlge := baseURL.String()
	GetData(urlge, &models.Result)

	return &models.GeoLocation{
		Lat:  fmt.Sprintf("%f", models.Result.Results[0].Geometry.Location.Lat),
		Lon:  fmt.Sprintf("%f", models.Result.Results[0].Geometry.Location.Lng),
		Name: address,
	}, nil
}
