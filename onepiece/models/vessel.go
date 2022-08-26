package models

type Vessel struct {
	Ladingnum   string `gorm:"column:ladingnum"`
	Cname       string `gorm:"column:cname"`
	Destination string `gorm:"column:destination`
}
