package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/h0wdyeve/hia/controllers"
	"github.com/h0wdyeve/hia/entity"
)

func main() {
	// สร้างการเชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(sqlite.Open("sa-entity.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// อัพเดทโครงสร้างของตารางในฐานข้อมูล
	db.AutoMigrate(&entity.Admin{}, &entity.Airline{}, &entity.Benefits{}, 
		&entity.Member{}, &entity.Point_Calculator{})

	// สร้างเซิร์ฟเวอร์ Gin
	r := gin.Default()

	// สร้างคอนโทรลเลอร์
	airlineController := &controllers.AirlineController{DB: db}

	// ตั้งค่าเส้นทาง API
	api := r.Group("/api")
	{
		// เส้นทางสำหรับ Admin
		api.POST("/admins", func(c *gin.Context) {
			var admin entity.Admin
			if err := c.ShouldBindJSON(&admin); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			db.Create(&admin)
			c.JSON(200, admin)
		})

		api.GET("/admins", func(c *gin.Context) {
			var admins []entity.Admin
			db.Find(&admins)
			c.JSON(200, admins)
		})

		api.GET("/admins/:id", func(c *gin.Context) {
			id := c.Param("id")
			var admin entity.Admin
			if err := db.First(&admin, id).Error; err != nil {
				c.JSON(404, gin.H{"error": "Admin not found"})
				return
			}
			c.JSON(200, admin)
		})

		api.PUT("/admins/:id", func(c *gin.Context) {
			id := c.Param("id")
			var admin entity.Admin
			if err := db.First(&admin, id).Error; err != nil {
				c.JSON(404, gin.H{"error": "Admin not found"})
				return
			}
			if err := c.ShouldBindJSON(&admin); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			db.Save(&admin)
			c.JSON(200, admin)
		})

		api.DELETE("/admins/:id", func(c *gin.Context) {
			id := c.Param("id")
			if err := db.Delete(&entity.Admin{}, id).Error; err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, gin.H{"message": "Admin deleted"})
		})

		// Airline routes ใช้ AirlineController
		api.GET("/airlines", airlineController.GetAllAirlines)
		api.GET("/airlines/:id", airlineController.GetAirlineByID)
		api.POST("/airlines", airlineController.CreateAirline)
		api.PUT("/airlines/:id", airlineController.UpdateAirline)
		api.DELETE("/airlines/:id", airlineController.DeleteAirline)

		// เส้นทางอื่นๆ เช่น Benefits, Member, และ Point_Calculator สามารถเพิ่มโค้ดคล้ายๆ กันในภายหลัง
	}

	// เริ่มเซิร์ฟเวอร์
	r.Run(":8080")
}
