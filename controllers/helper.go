package controllers

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, err error) {
	c.String(403, err.Error())
	c.AbortWithError(403, err)
}
