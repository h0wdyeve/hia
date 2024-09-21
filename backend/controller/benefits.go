package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/h0wdyeve/hia/config"
	"github.com/h0wdyeve/hia/entity"
)

func ListBenefits(c *gin.Context) {
	var moviepackages []entity.Benefits
	db := config.DB()

	results := db.Select("id, Package_name, Price, Duration").Find(&moviepackages)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, moviepackages)
}
