package config

import (
	"rumahbelajar-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data1.db")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	// Creating the table Wali Kelas
	if !db.HasTable(&models.WaliKelass{}) {
		db.CreateTable(&models.WaliKelass{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.WaliKelass{})
	}

	return db
}
