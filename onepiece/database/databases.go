package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"onepiece/models"
)

var db *gorm.DB

type Model struct {
	CreatedOn  int `json:"created_on" column:"created_on"`
	ModifiedOn int `json:"modified_on" column:"modified_on"`
	DeletedOn  int `json:"deleted_on"  column:"deleted_on"`
}

func Setup() {
	var err error
	dsn := "host=39.99.153.228 user=postgres password=postgres dbname=postgres port=5432"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.AutoMigrate(&models.Newstockin{})
	db.AutoMigrate(&models.Newstockout{})
}

func Getdatabase() *gorm.DB {
	return db
}
