package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})

	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})

	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})

	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})

	})
	r.Run(":8888")
}
