package controllers

import (
	"net/http"

	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

type WorksValidation struct {
	IDRole      string `json:"role" binding:"required"`
	Name 			string `json:"name" binding:"required"`
	Tech        string `json:"tech" binding:"required"`
	Description string `json:"description"`
}

// GET /works
// Get all works
func FindWorks(c *gin.Context) {
	var works []models.Work
	models.DB.Find(&works)

	c.JSON(http.StatusOK, gin.H{"data": works})
}

// POST /works
// Create new works
func CreateWorks(c *gin.Context) {
	var input WorksValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	work := models.Work{
		IDRole		: input.IDRole,
		Name			: input.Name,
		Tech			: input.Tech,
		Description	: input.Description,
	}
	models.DB.Create(&work)

	c.JSON(http.StatusOK, gin.H{"data": work})
}