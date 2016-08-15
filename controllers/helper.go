package controllers

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, err interface{}) {
	switch e := err.(type) {
	case error:
		c.JSON(200, gin.H{
			"error": e.Error(),
		})
	case string:
		c.JSON(200, gin.H{
			"error": e,
		})
	}
}
