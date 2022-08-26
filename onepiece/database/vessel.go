package database

import (
	"onepiece/models"
)

func GetAllVessel(vessel map[string]interface{}) (error, []models.Vessel, int64) {
	var vesselData []models.Vessel
	page := vessel["page"].(int)
	pageSize := vessel["limit"].(int)
	searchLadingnum := vessel["searchLadingnum"].(string)
	var total int64
	err := db.Table("Vessel").Select("ladingnum, cname,destination").Where("ladingnum like ?", searchLadingnum+"%").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&vesselData).Error
	return err, vesselData, total
}

func GetLadingnum(vessel map[string]interface{}) (error, []models.TTE, int64) {
	var vesselDate []models.TTE
	var total int64
	err := db.Table("Vessel").Select("ladingnum, cname,destination").Count(&total).Error
	return err, vesselDate, total
}
func AddVessel(vessel models.Vessel) error {
	err := db.Table("Vessel").Select("ladingnum", "cname", "destination").Create(&vessel).Error
	return err
}

func UpdateVessel(vessel models.Vessel) error {
	err := db.Table("Vessel").Where("ladingnum = ?", vessel.Ladingnum).Updates(&vessel).Error
	return err
}

func DelVessel(id string) error {
	var vessel models.Vessel
	err := db.Table("Vessel").Where("cname = ?", id).Delete(&vessel).Error
	return err
}
