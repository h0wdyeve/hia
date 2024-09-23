package entity

import ("gorm.io/gorm"
		"time")

type Admin struct {
	gorm.Model
	Email		string
	Password	string
	FirstName	string
	LastName	string
	BirthDay	time.Time

	// Benefits []Benefits `gorm:"foreignKey:AdminID"`
}