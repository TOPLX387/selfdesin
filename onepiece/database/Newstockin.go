package database

import (
	"onepiece/models"
)

func GetAllstockinlist(info models.InstockSearch) (list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := db.Model(&models.Newstockin{})
	var stockinData []models.Newstockin
	if info.Lpnum != "" {
		db = db.Where("lpnum = ?", info.Lpnum)
	}
	if info.Whtime != "" {
		db = db.Where("whtime = ?", info.Whtime)
	}
	err = db.Count(&total).Offset(offset).Limit(limit).Find(&stockinData).Error
	return list, total, err

}

func GetAllstockinstruct(info models.InstockSearch) (stockin []models.Newstockin, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := db.Model(&models.Newstockin{})
	var stockinData []models.Newstockin
	if info.Lpnum != "" {
		db = db.Where("lpnum = ?", info.Lpnum)
	}
	if info.Whtime != "" {
		db = db.Where("whtime = ?", info.Whtime)
	}
	err = db.Count(&total).Offset(offset).Limit(limit).Find(&stockinData).Error
	return stockinData, total, err

}

func AddInstock(stockinData models.Newstockin) (err error) {
	err = db.Create(&stockinData).Error
	return err
}

func DeleteInstock(id string) (err error) {
	var stockinData models.Newstockin
	println(id)
	err = db.Table("newstockins").Where("Lpnum = ?", id).Delete(&stockinData).Error
	return err
}

func UpdateInstock(stockinData models.Newstockin) (err error) {
	err = db.Table("newstockins").Where("Lpnum = ? And Whtime = ?", stockinData.Lpnum, stockinData.Whtime).Updates(&stockinData).Error
	return err
}

func FindInstock(stockio map[string]interface{}) (error, []models.Newstockin, int64) {
	var stockioData []models.Newstockin
	page := stockio["page"].(int)
	pageSize := stockio["limit"].(int)
	searchLpnum := stockio["searchLpnum"].(string)
	var total int64
	err := db.Table("newstockins").Where("lpnum like ?", searchLpnum+"%").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&stockioData).Error
	return err, stockioData, total
}
