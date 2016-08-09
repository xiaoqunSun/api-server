package controllers

//import (
//	"encoding/json"
//	"log"
//	"net/http"

//	"github.com/gin-gonic/gin"
//	"github.com/xiaoqunSun/api-server/mysql"
//)

//func HandlerAccount(r *gin.Engine) {
//	r.GET("/versions", func(c *gin.Context) {
//		appVersion := c.Query("appVersion")
//		//platform := c.Query("platform")
//		db := mysql.DB()

//		rows, err := db.Query("SELECT appVersion FROM versions where appVersion = ?", appVersion)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer rows.Close()
//		for rows.Next() {
//			var appVersion string
//			if err := rows.Scan(&appVersion); err != nil {
//				log.Fatal(err)
//			}
//		}
//		if err := rows.Err(); err != nil {
//			log.Fatal(err)
//		}

//		c.String(200, "param %s %d", c.Query("xx"), http.StatusOK)
//	})
//}
