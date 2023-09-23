package db

import (
	"educational_program_creator/internal/config"
	"educational_program_creator/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(cfg config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	models.AutoMigrate(db)

	return db, nil
}
