package main

import (
	"xyla_prize_api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	prizes = map[string]int{"shoes": 1, "clothes": 1, "pants": 1, "hats": 1}
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/prizes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"prizes": prizes,
			"xx":     "xxx",
		})
	})

	r.POST("/prizes", func(c *gin.Context) {
		key := c.Query("name")
		utils.Logger.Infoln(key)
		if _, ok := prizes[key]; ok {
			if prizes[key] > 0 {
				prizes[key] = prizes[key] - 1
			}
		}
		c.JSON(200, gin.H{
			"message": "good",
		})
	})

	r.Run(":12345")
}
