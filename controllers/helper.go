package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, err interface{}) {
	switch e := err.(type) {
	case error:
		c.String(403, e.Error())
		c.AbortWithError(403, e)
	case string:
		c.String(403, e)
		log.Println(e)
		c.Abort()
	}
}
