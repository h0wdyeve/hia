// controllers/airline_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/h0wdyeve/hia/entity"
)

type AirlineController struct {
	DB *gorm.DB
}

// GetAllAirlines - ดึงข้อมูลสายการบินทั้งหมด
func (ctrl *AirlineController) GetAllAirlines(c *gin.Context) {
	var airlines []entity.Airline
	if err := ctrl.DB.Preload("Point_Calculators").Preload("Benefits").Find(&airlines).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, airlines)
}

// GetAirlineByID - ดึงข้อมูลสายการบินตาม ID
func (ctrl *AirlineController) GetAirlineByID(c *gin.Context) {
	id := c.Param("id")
	var airline entity.Airline
	if err := ctrl.DB.Preload("Point_Calculators").Preload("Benefits").First(&airline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airline not found"})
		return
	}
	c.JSON(http.StatusOK, airline)
}

// CreateAirline - สร้างสายการบินใหม่
func (ctrl *AirlineController) CreateAirline(c *gin.Context) {
	var airline entity.Airline
	if err := c.ShouldBindJSON(&airline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&airline).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, airline)
}

// UpdateAirline - อัพเดตข้อมูลสายการบิน
func (ctrl *AirlineController) UpdateAirline(c *gin.Context) {
	id := c.Param("id")
	var airline entity.Airline
	if err := ctrl.DB.First(&airline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airline not found"})
		return
	}

	if err := c.ShouldBindJSON(&airline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Save(&airline).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, airline)
}

// DeleteAirline - ลบข้อมูลสายการบิน
func (ctrl *AirlineController) DeleteAirline(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&entity.Airline{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Airline deleted"})
}
