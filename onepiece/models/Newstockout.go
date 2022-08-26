package models

import "time"

type Newstockout struct {
	CreatedAt  time.Time
	Detime     string  `gorm:"column:detime"`
	Ordernum   string  `gorm:"column:ordernum"`
	Ladingnum  string  `gorm:"column:ladingnum"`
	Obton      string  `gorm:"column:obton"`
	Outpieces  string  `gorm:"column:outpieces"`
	Storagefee float64 `gorm:"column:storsgefee"`
}
