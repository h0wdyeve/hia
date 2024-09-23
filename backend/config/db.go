package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/h0wdyeve/hia/entity"
	"fmt"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("Netfilm2.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
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

	Airlines := []entity.Airline{
		{AirlineName: "AirAsia"},
		{AirlineName: "Thai Airways"},
		{AirlineName: "NokAir"},
		{AirlineName: "Vietjet"},
	}
	for _, pkg := range Airlines {
		db.FirstOrCreate(&pkg,entity.Airline{AirlineName: pkg.AirlineName})
	}

	BenefitsDetails := []entity.Benefits{
		{BenefitsName: "Bangkok (Suvarnabhumi) - Samui", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Samui", PointRequired: 6500, Quantity: 10, Code: "sFlfg487", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Samui", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Samui", PointRequired: 4000, Quantity: 10, Code: "Agjyj201", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Maldives", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Maldives", PointRequired: 14000, Quantity: 10, Code: "liuaq006", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Maldives", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Maldives", PointRequired: 11500, Quantity: 10, Code: "jsoem493", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Lampang", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Lampang", PointRequired: 5000, Quantity: 10, Code: "Yoash039", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Lampang", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Lampang", PointRequired: 2500, Quantity: 10, Code: "Hqpeo348", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Phuket", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Phuket", PointRequired: 5000, Quantity: 10, Code: "cnosd882", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Phuket", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Phuket", PointRequired: 2500, Quantity: 10, Code: "Udcnn235", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Trat", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Trat", PointRequired: 5000, Quantity: 10, Code: "Wewld398", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Trat", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Trat", PointRequired: 2500, Quantity: 10, Code: "spdmk948", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Chiang Mai", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Chiang Mai", PointRequired: 5000, Quantity: 10, Code: "Vosdl455", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Chiang Mai", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Chiang Mai", PointRequired: 2500, Quantity: 10, Code: "kdfai588", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Sukhothai", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Sukhothai", PointRequired: 5000, Quantity: 10, Code: "Dsaop358", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Bangkok (Suvarnabhumi) - Sukhothai", FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Sukhothai", PointRequired: 2500, Quantity: 10, Code: "akdpU482", Class: "Economy",Trip: "One Way"},
		{BenefitsName: "Chiang Mai - Krabi", FlyingFrom: "Chiang Mai", GoingTo: "Krabi", PointRequired: 9000, Quantity: 10, Code: "hyKrs329", Class: "Business",Trip: "One Way"},
		{BenefitsName: "Chiang Mai - Krabi", FlyingFrom: "Chiang Mai", GoingTo: "Krabi", PointRequired: 6500, Quantity: 10, Code: "sRpeo396", Class: "Economy",Trip: "One Way"},
	}
	for _, pkg := range BenefitsDetails {
		db.FirstOrCreate(&pkg,entity.Benefits{Code: pkg.Code})
	}

	PointCal := []entity.Point_Calculator{
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Samui", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Samui", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Maldives", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Maldives", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Lampang", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Lampang", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Phuket", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Phuket", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Trat", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Trat", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Chiang Mai", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Chiang Mai", Point: 500, Class: "Economy"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Sukhothai", Point: 750, Class: "Business"},
		{FlyingFrom: "Bangkok (Suvarnabhum)", GoingTo: "Sukhothai", Point: 500, Class: "Economy"},
		{FlyingFrom: "Chiang Mai", GoingTo: "Krabi", Point: 750, Class: "Business"},
		{FlyingFrom: "Chiang Mai", GoingTo: "Krabi", Point: 500, Class: "Economy"},
	}
	for _, pcloop := range PointCal {
		db.FirstOrCreate(&pcloop,entity.Point_Calculator{GoingTo: pcloop.GoingTo, Point: pcloop.Point})
	}
}

