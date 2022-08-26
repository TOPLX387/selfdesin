package models

import "time"

type Newstockin struct {
	CreatedAt time.Time
	Whtime    string `gorm:"column:whtime"`
	Lpnum     string `gorm:"column:lpnum"`
	Ibton     string `gorm:"column:ibton"`
	Ibpieces  string `gorm:"column:ibpieces"`
}
