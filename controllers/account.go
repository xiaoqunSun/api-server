package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqunSun/api-server/mysql"
)

func HandlerAccount(r *gin.Engine) {
	r.POST("/registerAccount", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if len(username) < 6 || len(username) > 15 {
			errorResponse(c, "username length must be in 6-15")
			return
		}
		if len(password) < 6 || len(password) > 15 {
			errorResponse(c, "password length must be in 6-15")
			return
		}
		md5Sum := md5.Sum([]byte(password))
		log.Println("username", username)
		log.Println("password", password, md5Sum)

		db := mysql.DB()
		rows, err := db.Query("call sp_registerAccount(?,?)", username, hex.EncodeToString(md5Sum[:]))
		log.Println("rows", rows)
		var result int
		defer rows.Close()
		if err != nil {
			errorResponse(c, err)
			return
		}
		if rows == nil {
			errorResponse(c, "internal error!")
			return
		}
		for rows.Next() {
			if err := rows.Scan(&result); err != nil {
				errorResponse(c, err)
				return
			}
		}
		if result == 0 {
			c.JSON(200, gin.H{})
		} else if result == 1 {
			errorResponse(c, "username has exist")
		}
	})
}
