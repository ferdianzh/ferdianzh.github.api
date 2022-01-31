package services

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	extendsion := filepath.Ext(file.Filename)
	file.Filename = fmt.Sprint(time.Now().Unix()) + RandomString(5) + extendsion
	filepath := "/images/" + c.Param("folder") + "/" + file.Filename

	if err:= c.SaveUploadedFile(file, "./storages/" + filepath); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}