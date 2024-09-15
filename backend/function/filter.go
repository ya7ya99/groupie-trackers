package function

import (
	"strconv"

	"groupie-trackers/backend/models"
)

func Filter(creationYearFrom, creationYearTo, albumYearFrom, albumYearTo, location string, members []string) []models.Artist {
	
	var ArtistsF []models.Artist
	if len(members) == 0 {
		ArtistsF = Filters(creationYearFrom, creationYearTo, albumYearFrom, albumYearTo, location, "")
	} else {
		for _, num := range members {
			ArtistsF = append(ArtistsF, Filters(creationYearFrom, creationYearTo, albumYearFrom, albumYearTo, location, num)...)
		}
	}
	return ArtistsF
}

func Filters(creationYearFrom, creationYearTo, albumYearFrom, albumYearTo, location, members string) []models.Artist {
	var ArtistsF []models.Artist
	cYearFrom, _ := strconv.Atoi(creationYearFrom)
	cYearTo, _ := strconv.Atoi(creationYearTo)
	numMembers, _ := strconv.Atoi(members)

	for _, artist := range models.Var_Artists {
		matchesAll := true
		// Check creation year

		if artist.CreationDate < cYearFrom {
			matchesAll = false
		}
		if artist.CreationDate > cYearTo {
			matchesAll = false
		}
		// Check number of members
		if len(members) != 0 && len(artist.Members) != numMembers {
			matchesAll = false
		}

		// Check album year
		if artist.FirstAlbum[len(artist.FirstAlbum)-4:] < albumYearFrom {
			matchesAll = false
		}
		if  artist.FirstAlbum[len(artist.FirstAlbum)-4:] > albumYearTo {
			matchesAll = false
		}
		// Check location
		if location != "" {
			locationMatch := false
			for _, loc := range models.Var_Locations.Index[artist.ID-1].Locations {
				if loc == location {
					locationMatch = true
					break
				}
			}
			if !locationMatch {
				matchesAll = false
			}
		}
		// If all criteria match, add to filtered list
		if matchesAll {
			ArtistsF = append(ArtistsF, artist)
		}
	}
	return ArtistsF
}
