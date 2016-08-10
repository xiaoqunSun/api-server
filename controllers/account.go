package controllers

import (
	"crypto/md5"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqunSun/api-server/mysql"
)

func HandlerAccount(r *gin.Engine) {
	r.POST("/registerAccount", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if len(username) < 6 || len(username) > 15 {
			errorResponse(c, "username length must be in 6-15")
		}
		if len(password) < 6 || len(password) > 15 {
			errorResponse(c, "password length must be in 6-15")
		}
		db := mysql.DB()
		rows, err := db.Query("call sp_registerAccount(?,?)", username, md5.Sum([]byte(password)))
		var result int
		defer rows.Close()
		if err != nil {
			errorResponse(c, err)
		}
		for rows.Next() {
			if err := rows.Scan(&result); err != nil {
				errorResponse(c, err)
			}
		}
		if result == 0 {
			c.JSON(200, gin.H{})
		} else if result == 1 {
			c.JSON(200,gin.H{
				"error":"username has exist"
			})
		}
	})
}
