package function

import (
	"onepiece/database"
	"onepiece/models"
)

func GetLogintoByName(name string) (error, []models.Loginto, int64) {
	var logintoData []models.Loginto
	page := 1
	pageSize := 1
	searchName := name
	err, logintoData, total := database.GetAllLoginto(map[string]interface{}{"page": page, "limit": pageSize, "searchName": searchName})
	return err, logintoData, total
}
