package entity

import "gorm.io/gorm"

type Airline struct {
	gorm.Model
	AirlineName		string

	// Point_Calculators []Point_Calculator `gorm:"foreignKey:AirlineID"`

	// Benefits []Benefits `gorm:"foreignKey:AirlineID"`
}