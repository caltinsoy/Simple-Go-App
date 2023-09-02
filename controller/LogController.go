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

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
