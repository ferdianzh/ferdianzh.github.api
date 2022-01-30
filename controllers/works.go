package controllers

import (
	"net/http"

	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

// GET /works
// Get all works
func FindWorks(c *gin.Context) {
	var works []models.Work
	models.DB.Find(&works)

	c.JSON(http.StatusOK, gin.H{"data": works})
}

type CreateWorkInput struct {
	IDRole      string `json:"role" binding:"required"`
	Name 			string `json:"name" binding:"required"`
	Tech        string `json:"tech" binding:"required"`
	Description string `json:"description"`
}
// POST /works
// Create new works
func CreateWork(c *gin.Context) {
	var input CreateWorkInput
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

// GET /works/:id
// Find a work
func FindWork(c *gin.Context) {
	var work models.Work
	if err := models.DB.Where("id = ?", c.Param("id")).First(&work).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": work})
}

type UpdateWorkInput struct {
	IDRole      string `json:"role"`
	Name 			string `json:"name"`
	Tech        string `json:"tech"`
	Description string `json:"description"`
}
// PATCH /works/:id
// Update a work
func UpdateWork(c *gin.Context) {
	var work models.Work
	if err := models.DB.Where("id = ?", c.Param("id")).First(&work).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	var input UpdateWorkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&work).Updates(models.Work{
		IDRole		: input.IDRole,
		Name			: input.Name,
		Tech			: input.Tech,
		Description	: input.Description,
	})

	c.JSON(http.StatusOK, gin.H{"data": work})
}

// DELETE /works/:id
// Delete a work
func DeleteWork(c *gin.Context) {
	var work models.Work
	if err := models.DB.Where("id = ?", c.Param("id")).First(&work).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&work)

	c.JSON(http.StatusOK, gin.H{"data": true})
}