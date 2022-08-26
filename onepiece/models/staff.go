package models

import "time"

type Loginto struct {
	LogintoId       int       `gorm:"column:loginto_id"`
	LogintoName     string    `grom:"column:loginto_name"`
	LogintoPassword string    `grom:"column:loginto_password"`
	LogintoLevel    int       `grom:"column:loginto_level"`
	LogintoRemarks  string    `grom:"column:loginto_remarks"`
	UpdateTime      time.Time `gorm:"column:update_time"`
}
