package controllers

import (
	"github.com/gin-gonic/gin"
)

func HandlerServerAddr(r *gin.Engine) {
	r.GET("/serverAddr", func(c *gin.Context) {
		data := make(gin.H)

		data["name"] = "game"
		data["ls_ip"] = "139.129.33.151"
		data["ls_port"] = 7777
		data["gs_ip"] = "139.129.33.151"
		data["gs_port"] = 8888

		c.JSON(200, data)

	})
}
