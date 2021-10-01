package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/follower/:USERNAME", func(c *gin.Context) {
		USERNAME := c.Param("USERNAME")
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "jumlah follower USERNAME = " + USERNAME,
		})
	})
	router.GET("/:USERID", func(c *gin.Context) {
		USERID := c.Param("USERID")
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "username dan follower USERID = " + USERID,
		})
	})

	router.Run()
}
