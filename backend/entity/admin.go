package entity

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email		string
	Password	string
	FirstName	string
	LastName	string
	BirthDay	string

	Benefits []Benefits `gorm:"foreignKey:AdminID"`
}