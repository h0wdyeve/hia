package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/h0wdyeve/hia/config"
	"github.com/h0wdyeve/hia/entity"
)

func GetAllBenefits(c *gin.Context) {
	var Benefits []entity.Benefits
	db := config.DB()

	results := db.Select("id, Package_name, Price, Duration").Find(&Benefits)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, Benefits)
}

func GetAllBenefitsByID(c *gin.Context) {
	var Benefits entity.Benefits
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&Benefits, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Benefits not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": Benefits})
}

func DeleteBenefits(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM moviepackages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}


