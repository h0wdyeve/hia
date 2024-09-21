package main

import (
	"github.com/h0wdyeve/hia/config"
	"github.com/h0wdyeve/hia/controller"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	// เริ่มต้นเชื่อมต่อกับฐานข้อมูล
	config.SetupDatabase()

	// สร้าง Gin engine ใหม่
	r := gin.Default()
	
	// ใช้ CORS middleware
	r.Use(CORSMiddleware())
	r.GET("/benefits", controller.GetAllBenefits)
	r.GET("/benefits/:id", controller.GetAllBenefitsByID)
	r.DELETE("/benefits/:id", controller.DeleteBenefits)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // หากเป็นคำขอ OPTIONS ให้ตอบกลับด้วย 204
			return
		}
		c.Next() // ดำเนินการต่อกับคำขอที่ไม่ใช่ OPTIONS
	}
}