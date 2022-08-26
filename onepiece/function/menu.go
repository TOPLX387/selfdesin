package function

import (
	"onepiece/database"
	"onepiece/models"
)

func GetMenuByLevel(level int) (error, []models.Menu) {
	err, menuData := database.GetMenuByLevel(level)
	return err, menuData
}
