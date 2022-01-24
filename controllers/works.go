package controllers

import (
	"net/http"

	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

func FindWorks(c *gin.Context) {
	var works []models.Works
	models.DB.Find(&works)

	c.JSON(http.StatusOK, gin.H{
		"data": works,
	})
}