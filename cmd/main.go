package main

import (
	"database/sql"
	"educational_program_creator/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to open", err)
		return
	}
	models.AutoMigrate(db)

	sqlDB, _ := db.DB()

	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
}
