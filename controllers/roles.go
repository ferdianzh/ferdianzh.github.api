package controllers

import (
	"net/http"

	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

// GET /roles
// Get all roles
func FindRoles(c *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

type CreateRoleInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
// POST /roles
// Create new roles
func CreateRoles(c *gin.Context) {
	var input CreateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	role := models.Role{
		Name			: input.Name,
		Description	: input.Description,
	}
	models.DB.Create(&role)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// GET /roles/:id
// Find a role
func FindRole(c *gin.Context) {
	var role models.Role
	if err := models.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

type UpdateRoleInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
// PATCH /roles/:id
// Update a role
func UpdateRole(c *gin.Context) {
	var role models.Role
	if err := models.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	var input UpdateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&role).Updates(models.Role{
		Name			: input.Name,
		Description	: input.Description,
	})
	
	c.JSON(http.StatusOK, gin.H{"data": input, "role": role})
}

// DELETE /roles/:id
// Delete a role
func DeleteRole(c *gin.Context) {
	var role models.Role
	if err := models.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}