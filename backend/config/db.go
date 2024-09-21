package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/h0wdyeve/hia/entity"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("G11_PROJECT.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&entity.Admin{},
		&entity.Airline{},
		&entity.Benefits{},
		&entity.Member{},
		&entity.Point_Calculator{},
	)
	db = database
}
