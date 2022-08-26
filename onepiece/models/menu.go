package models

type Menu struct {
	MenuId      int    `gorm:"column:menu_id"`
	MenuAddress string `gorm:"column:menu_address"`
	MenuLevel   int    `gorm:"column:menu_level"`
	MenuName    string `grom:"column:menu_name"`
}
