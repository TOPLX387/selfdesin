package database

import (
	"onepiece/models"
)

func GetAllstockoutlist(info models.OutstockSearch) (list interface{}, total int64, err error) {
	var stockoutData []models.Newstockout
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := db.Model(&models.Newstockout{})
	if info.Ordernum != "" {
		db = db.Where("ordernum = ?", info.Ordernum)
	}
	if info.Detime != "" {
		db = db.Where("detime = ?", info.Detime)
	}
	err = db.Order("created_at asc").Count(&total).Offset(offset).Limit(limit).Find(&stockoutData).Error
	return list, total, err
}

func GetAllstockoutstruct(info models.OutstockSearch) (stockout []models.Newstockout, total int64, err error) {
	var stockoutData []models.Newstockout
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := db.Model(&models.Newstockout{})
	if info.Ordernum != "" {
		db = db.Where("ordernum = ?", info.Ordernum)
	}
	if info.Detime != "" {
		db = db.Where("detime = ?", info.Detime)
	}
	err = db.Order("created_at asc").Count(&total).Offset(offset).Limit(limit).Find(&stockoutData).Error
	return stockoutData, total, err
}

func AddOutstock(stockoutData models.Newstockout) (err error) {
	err = db.Create(&stockoutData).Error
	return err
}

func DeleteOutstock(id string) (err error) {
	var stockoutData models.Newstockout
	println(id)
	err = db.Table("newstockouts").Where("Ordernum = ?", id).Delete(&stockoutData).Error
	return err
}

func UpdateOutstock(stockoutData models.Newstockout) (err error) {
	err = db.Table("newstockouts").Where("Ordernum = ? And Detime = ?", stockoutData.Ordernum, stockoutData.Detime).Updates(&stockoutData).Error
	return err
}

func FindOutstock(stockio map[string]interface{}) (error, []models.Newstockout, int64) {
	var stockioData []models.Newstockout
	page := stockio["page"].(int)
	pageSize := stockio["limit"].(int)
	searchOrdernum := stockio["searchOrdernum"].(string)
	var total int64
	err := db.Table("newstockouts").Where("ordernum like ?", searchOrdernum+"%").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&stockioData).Error
	return err, stockioData, total
}
