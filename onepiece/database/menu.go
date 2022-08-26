package database

import (
	"onepiece/models"
	"onepiece/packages/gredis"
	"strconv"
)

func GetMenuByLevel(level int) (error, []models.Menu) {
	var menuData []models.Menu
	cache := "mksys:menu" + strconv.Itoa(level)
	err := db.Table("menu").Where("menu_level = ?", level).Find(&menuData).Error
	if len(menuData) != 0 {
		gredis.RedisConn.Set(cache, menuData, 10000)
	}
	return err, menuData
}
