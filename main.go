package main

import (
	"github.com/ferdianzh/ferdianzh.github.api/controllers"
	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	
	// still need security with jwt
	r.GET("/roles", controllers.FindRoles)
	r.POST("/roles", controllers.CreateRoles)
	r.GET("/roles/:id", controllers.FindRole)
	r.PATCH("/roles/:id", controllers.UpdateRole)
	r.DELETE("/roles/:id", controllers.DeleteRole)

	r.GET("/works", controllers.FindWorks)
	r.POST("/works", controllers.CreateWorks)

	r.Static("/img", "./storages/images")

	r.Run("localhost:8080")
}