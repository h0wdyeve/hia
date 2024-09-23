package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/h0wdyeve/hia/config"
	"github.com/h0wdyeve/hia/entity"
)

func GetAllAirline(c *gin.Context) {
    var Airline []entity.Airline

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&Airline).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"airlines": Airline})
}


func GetAirlineByID(c *gin.Context) {
	var Airline entity.Airline
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&Airline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airline not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": Airline})
}

func DeleteAirline(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Airline WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}


