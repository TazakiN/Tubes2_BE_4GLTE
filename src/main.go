package main

import (
	"net/http"
	"strings"

	logic "github.com/TazakiN/Tubes2_BE_4GLTE/logic"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:metode/:linkMulai/:linkTujuan", getData)

	router.Run("localhost:3321")
}

func getData(c *gin.Context) {
	metode := strings.ToLower(c.Param("metode"))
	linkMulai := c.Param("linkMulai")
	linkTujuan := c.Param("linkTujuan")
	hasil := []string{}

	if metode == "bfs" {
		hasil = logic.BFS(linkMulai, linkTujuan)
	} else if metode == "ids" {
		hasil = logic.IDS(linkMulai, linkTujuan)
	}

	if metode != "bfs" && metode != "ids" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terjadi kesalahan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hasil": hasil,
	})
}
