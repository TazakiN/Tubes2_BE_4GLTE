package main

import (
	"net/http"
	"time"

	logic "github.com/TazakiN/Tubes2_BE_4GLTE/logic"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	router.GET("/:bahasa/:metode/:linkMulai/:linkTujuan", getData)

	router.Run("localhost:3321")
}

func getData(c *gin.Context) {
	bahasa := c.Param("bahasa")
	metode := c.Param("metode")
	linkMulai := c.Param("linkMulai")
	linkTujuan := c.Param("linkTujuan")
	hasil := []string{}

	startTime := time.Now()
	if metode == "BFS" {
		hasil = logic.BFS(linkMulai, linkTujuan, bahasa)
	} else if metode == "IDS" {
		hasil = logic.IDS(linkMulai, linkTujuan, bahasa)
	}
	endTime := time.Now()
	elapseTime := endTime.Sub(startTime).String()

	if metode != "BFS" && metode != "IDS" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terjadi kesalahan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hasil":          hasil,
		"panjang solusi": len(hasil) - 1, // asumsi link awal ga dihitung
		"waktu":          elapseTime,
	})
}
