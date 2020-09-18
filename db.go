package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

// gorm fibonacci database model
type FibDB struct {
	ID       uint8 `gorm:"primary_key"`
	Current  uint64
	Previous uint64
}

// checks to see if DB exists, if not creates it.
// if database exists, then reads fibonacci current and previous numbers from it
// if error is received, logs error and sets fibonacci numbers to O and O
func checkDB() {
	database, err := gorm.Open("sqlite3", "fibonacci.db")
	if err != nil {
		log.Println(err)
		f.Current = 0
		f.Previous = 0
	}
	if err := database.AutoMigrate(&FibDB{}).Error; err != nil {
		log.Println(err)
	}
	DB = database
	var fibSetup FibDB
	if err := DB.Where("ID = 1").First(&fibSetup).Error; err != nil {
		log.Println(err)
		f.Current = 0
		f.Previous = 0
		fibSetup = FibDB{1, 0, 0}
		if err := DB.Create(&fibSetup).Error; err != nil {
			log.Println(err)
		}
		return
	}
	f.Current = fibSetup.Current
	f.Previous = fibSetup.Previous
}

// writes values of current and previous fibonacci numbers to database
// logs error if write process fails.
func writeToDB(current, previous uint64) {
	var fibWrite = FibDB{1, current, previous}
	if err := DB.Model(&FibDB{ID: 1}).Update(fibWrite).Error; err != nil {
		log.Println(err)
	}
}
