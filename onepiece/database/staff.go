package database

import (
	"onepiece/models"
)

func GetAllLoginto(logintoParam map[string]interface{}) (error, []models.Loginto, int64) {
	var logintoData []models.Loginto
	page := logintoParam["page"].(int)
	pageSize := logintoParam["limit"].(int)
	searchName := logintoParam["searchName"].(string)
	var total int64
	// err := db.Table("loginto").Where("loginto_name like ? and is_del = false", searchName+"%").Order("loginto_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&logintoData).Error
	err := db.Table("loginto").Where("loginto_name like ?", searchName+"%").Order("loginto_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&logintoData).Error
	return err, logintoData, total
}

func GetLogintoById(id int) (error, models.Loginto) {
	var logintoData models.Loginto
	// err := db.Table("loginto").Where("loginto_id = ? and is_del = false", id).Find(&logintoData).Error
	err := db.Table("loginto").Where("loginto_id = ?", id).Find(&logintoData).Error
	return err, logintoData
}

func AddLoginto(loginto models.Loginto) error {
	err := db.Table("loginto").Select("loginto_name", "loginto_password", "loginto_level", "loginto_remarks", "update_time").Create(&loginto).Error
	return err
}

func UpdateLoginto(loginto models.Loginto) error {
	err := db.Table("loginto").Where("loginto_id = ?", loginto.LogintoId).Updates(&loginto).Error
	return err
}

func DelLoginto(id int) error {
	err := db.Table("loginto").Where("loginto_id = ?", id).Update("is_del", true).Error
	return err
}
