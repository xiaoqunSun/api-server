package main

import "github.com/gin-gonic/gin"
import _ "github.com/go-sql-driver/mysql"
import "net/http"

func main() {
	r := gin.Default()

	r.GET("/version", func(c *gin.Context) {
		c.String(200, "param %s %d", c.Query("xx"), http.StatusOK)
	})
	r.POST("/account", func(c *gin.Context) {
		c.String(200, "param %s %d", c.PostForm("ggg"), http.StatusOK)
	})

	r.Run(":8080")
}
