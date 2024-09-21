package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Email		string
	Password	string
	FirstName	string
	LastName	string
	BirthDay	string
	Gender		string
	TotalPoint	uint

	Point_Calculators []Point_Calculator `gorm:"many2many:Member_has_Point_Calculator"`
}