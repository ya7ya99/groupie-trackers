package models

type APIURLs struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}
type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
var Result struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
type GeoLocation struct {
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
	Name string `json:"name"`
}
var (
	Var_ApiURL     APIURLs
	Var_Artists    []Artist
	Var_Locations Locations
	Var_Dates     Dates
	Var_Relation  Relation
)

type PageData struct {
	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
	ArtistID  int
	Coordinates []GeoLocation
}

type ErrorData struct {
	Message string
	Number  int
}

var Err ErrorData
var ErrAPI error
