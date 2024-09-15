package handlers

import (
	"groupie-trackers/backend/api"
	"groupie-trackers/backend/models"
)

func GetData() {
	url := "https://groupietrackers.herokuapp.com/api"

	models.ErrAPI = api.GetData(url, &models.Var_ApiURL)

	models.ErrAPI = api.GetData(models.Var_ApiURL.Artists, &models.Var_Artists)

	models.ErrAPI = api.GetData(models.Var_ApiURL.Locations, &models.Var_Locations)

	models.ErrAPI = api.GetData(models.Var_ApiURL.Dates, &models.Var_Dates)

	models.ErrAPI = api.GetData(models.Var_ApiURL.Relation, &models.Var_Relation)
}
