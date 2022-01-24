package controllers

import (
	"net/http"

	"github.com/ferdianzh/ferdianzh.github.api/models"
	"github.com/gin-gonic/gin"
)

func FindMahasiswa(c *gin.Context) {
	var mhs []models.Mahasiswa
	models.DB.Table("mahasiswa").Find(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
	})
}