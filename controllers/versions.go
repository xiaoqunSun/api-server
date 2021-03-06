package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqunSun/api-server/mysql"
)

const static_prfix = "http://static.7171game.com:8080/static"

func HandlerVersions(r *gin.Engine) {
	r.GET("/versions", func(c *gin.Context) {
		appVersion := c.Query("appVersion")
		//platform := c.Query("platform")
		db := mysql.DB()
		rows, err := db.Query("SELECT version FROM versions where appVersion = ?", appVersion)
		defer rows.Close()

		if err != nil {
			errorResponse(c, err)
			return
		}
		data := make(gin.H)
		data["appVersion"] = appVersion
		for rows.Next() {
			var version string
			var url string
			if err := rows.Scan(&version); err != nil {
				errorResponse(c, err)
				return
			}
			url = fmt.Sprintf("%s/%s.zip", static_prfix, version)
			data[version] = url
		}
		if err := rows.Err(); err != nil {
			errorResponse(c, err)
			return
		}

		c.JSON(200, data)

	})
	r.POST("/versions", func(c *gin.Context) {
		if !validClientIP(c) {
			errorResponse(c, "此ip地址不允许执行操作，请联系管理员")
			return
		}
		appVersion := c.PostForm("appVersion")
		log.Println("/version post", appVersion)
		db := mysql.DB()
		rows, err := db.Query("select version from versions where appVersion = ? order by version desc limit 1;", appVersion)
		defer rows.Close()

		if err != nil {
			errorResponse(c, err)
			return

		}
		var version int
		for rows.Next() {
			if err := rows.Scan(&version); err != nil {
				errorResponse(c, err)
				return

			}
		}
		if err := rows.Err(); err != nil {
			errorResponse(c, err)
			return

		}
		version = version + 1
		_, err = db.Query("insert into versions set version = ?, appVersion = ?, created_at = ?", version, appVersion, time.Now())
		if err != nil {
			errorResponse(c, err)
			return

		}
		c.Status(200)

	})
}
