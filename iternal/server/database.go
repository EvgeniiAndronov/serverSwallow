package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"serverSwallow/iternal/models"
)

func SetupDatabase() {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserData{})
}

func NewDbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

//func NewDataStore() *
