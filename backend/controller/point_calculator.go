package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/h0wdyeve/hia/config"
	"github.com/h0wdyeve/hia/entity"
)

func GetAllPoint(c *gin.Context) {
	var Point []entity.Point_Calculator
	db := config.DB()

	results := db.Select("id, Package_name, Price, Duration").Find(&Point)  //อีฟแก้เอานะเราไม่รู้ต้องเอาไรบ้าง

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, Point)
}

func GetPointByID(c *gin.Context) {
	var Point entity.Point_Calculator
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&Point, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Point_Calculator not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": Point})
}

func DeletePoint(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Point_Calculator WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
