package main

import (
	"github.com/ferdianzh/ferdianzh.github.api/controllers"
	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	
	r.GET("/works", controllers.FindWorks)

	r.GET("/mhs", controllers.FindMahasiswa)

	r.Run("localhost:8080")
}