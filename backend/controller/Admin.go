// controllers/admin_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/h0wdyeve/hia/entity"
)

type Admin struct {
	DB *gorm.DB
}

// GetAllAdmins - ดึงข้อมูลผู้ดูแลระบบทั้งหมด
func (a *Admin) GetAllAdmins(c *gin.Context) {
	var admins []entity.Admin
	if err := a.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// GetAdminByID - ดึงข้อมูลผู้ดูแลระบบตาม ID
func (a *Admin) GetAdminByID(c *gin.Context) {
	id := c.Param("id")
	var admin entity.Admin
	if err := a.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// CreateAdmin - สร้างผู้ดูแลระบบใหม่
func (a *Admin) CreateAdmin(c *gin.Context) {
	var admin entity.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := a.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// UpdateAdmin - อัพเดตข้อมูลผู้ดูแลระบบ
func (a *Admin) UpdateAdmin(c *gin.Context) {
	id := c.Param("id")
	var admin entity.Admin
	if err := a.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// DeleteAdmin - ลบข้อมูลผู้ดูแลระบบ
func (a *Admin) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if err := a.DB.Delete(&entity.Admin{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted"})
}
