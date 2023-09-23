package main

import (
	"educational_program_creator/internal/app"
	"educational_program_creator/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.NewServer(cfg)
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	//if err != nil {
	//	log.Fatalln("Failed to open", err)
	//	return
	//}
	//models.AutoMigrate(db)
	//
	//sqlDB, _ := db.DB()
	//
	//defer func(sqlDB *sql.DB) {
	//	err := sqlDB.Close()
	//	if err != nil {
	//
	//	}
	//}(sqlDB)
}
