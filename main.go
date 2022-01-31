package main

import (
	"github.com/ferdianzh/ferdianzh.github.api/controllers"
	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/ferdianzh/ferdianzh.github.api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	
	// still need security with jwt
	r.GET("/roles", controllers.FindRoles)
	r.POST("/roles", controllers.CreateRole)
	r.GET("/roles/:id", controllers.FindRole)
	r.PATCH("/roles/:id", controllers.UpdateRole)
	r.DELETE("/roles/:id", controllers.DeleteRole)

	r.GET("/works", controllers.FindWorks)
	r.POST("/works", controllers.CreateWork)
	r.GET("/works/:id", controllers.FindWork)
	r.PATCH("/works/:id", controllers.UpdateWork)
	r.DELETE("/works/:id", controllers.DeleteWork)

	r.Static("/images", "./storages/images")
	r.POST("/upload/:folder", services.UploadImage)

	r.Run("localhost:8080")
}