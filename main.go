package main

import (
	"github.com/ferdianzh/ferdianzh.github.api/controllers/works"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/works", works.GetAll)

	router.Run()
}