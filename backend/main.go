package main

import (
	"github.com/sa67-lab5/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sa-entity.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Admin{}, &entity.Airline{}, &entity.Benefits{}, 
&entity.Member{}, &entity.Point_Calculator{})
}