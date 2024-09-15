package function

import (
	"strconv"
	"strings"

	"groupie-trackers/backend/models"
)

func Search(search string ) []models.Artist {
	var ArtistsF []models.Artist
	search = strings.ToLower(search)
	artistMap := make(map[int]bool) 
	for _, artist := range models.Var_Artists {
		addArtist := false
		// Check artist's name
		if strings.HasPrefix(strings.ToLower(artist.Name), search) {
			addArtist = true
		}
		// Check band members
		for _, members := range artist.Members {
			if strings.Contains(strings.ToLower(members), search) {
				addArtist = true
				break
			}
		}
		// Check first album
		if artist.FirstAlbum == search {
			addArtist = true
		}
		// Check creation date
		cdate := strconv.Itoa(artist.CreationDate)
		if search == cdate {
			addArtist = true
		}
		// Check location
		for _, loc := range models.Var_Locations.Index[artist.ID-1].Locations {
			if strings.Contains(loc, search) {
				addArtist = true
				break
			}
		}
		// Add artist to result list if not already added
		if addArtist && !artistMap[artist.ID] {
			ArtistsF = append(ArtistsF, artist)
			artistMap[artist.ID] = true
		}
	}
	return ArtistsF
}