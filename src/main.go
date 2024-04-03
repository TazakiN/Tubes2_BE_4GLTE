package main

import (
	"net/http"

	logic "github.com/TazakiN/Tubes2_BE_4GLTE/logic"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:metode/:linkMulai/:linkTujuan", getData)

	router.Run("localhost:3321")
}

func getData(c *gin.Context) {
	metode := c.Param("metode")
	linkMulai := c.Param("linkMulai")
	linkTujuan := c.Param("linkTujuan")

	if metode == "bfs" {
		logic.BFS(linkMulai, linkTujuan)
	} else if metode == "ids" {
		logic.IDS(linkMulai, linkTujuan)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Metode tidak ditemukan"})
	}
}
