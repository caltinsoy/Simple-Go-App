package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-crud/initializers"
	"go-crud/model"
	"go-crud/producer"
	"net/http"
)

func CreateLog(c *gin.Context, ctx context.Context) {
	var input model.LogDocument
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logDocument := model.LogDocument{AnyString: input.AnyString}
	initializers.DB.Create(&logDocument)
	go producer.Produce(ctx, logDocument.String())

	c.JSON(http.StatusOK, gin.H{"data": logDocument})
}

func GetLog(c *gin.Context) {

	id := c.Param("id")

	// Query the database to find the user with the specified ID
	var logDocument model.LogDocument
	if err := initializers.DB.Where("id = ?", id).First(&logDocument).Error; err != nil {
		c.JSON(404, gin.H{"error": "LogDocument not found"})
		return
	}

	// Return the user as JSON response
	c.JSON(200, logDocument)

}

func DeleteLog(c *gin.Context) {

	id := c.Param("id")

	// Query the database to find the user with the specified ID
	var logDocument model.LogDocument
	// Delete the user with the specified ID from the database
	if err := initializers.DB.Where("id = ?", id).Delete(&model.LogDocument{}).Error; err != nil {
		c.JSON(404, gin.H{"error": "LogDocument not found"})
		return
	}

	// Return the user as JSON response
	c.JSON(200, logDocument)

}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
