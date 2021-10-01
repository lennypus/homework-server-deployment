package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Homework Rakamin Server Deployment - Lenny Puspita Sari")
	})
	router.GET("/follower/:USERNAME", func(c *gin.Context) {
		USERNAME := c.Param("USERNAME")
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "total follower USERNAME = " + USERNAME,
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
