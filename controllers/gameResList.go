package controllers

import (
	"fmt"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqunSun/api-server/mysql"
)

type FileInfo struct {
	Md5      string
	FilePath string
	Url      string
}

var cache map[string]string
var cacheExpiration map[string]int64

const Cache_Expiration = 5 * time.Second

//返回游戏所需的资源文件列表
func HandlerGameReslist(r *gin.Engine) {
	cache = make(map[string]string)
	cacheExpiration = make(map[string]int64)
	r.GET("/gameResList", func(c *gin.Context) {
		gameID := c.Query("gameID")
		var filelist string
		filelist, ok := cache[gameID]

		fmt.Println(ok, cacheExpiration[gameID], time.Now().Unix())
		if !ok || cacheExpiration[gameID] < time.Now().Unix() {
			db := mysql.DB()
			rows, err := db.Query("SELECT filelist FROM gameResList where gameID = ?", gameID)
			if err != nil {
				errorResponse(c, err)
				return
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&filelist); err != nil {
					errorResponse(c, err)
					return
				}
				break
			}
			if filelist == "" {
				c.JSON(200, make(gin.H))
				return
			}
			cache[gameID] = filelist
			cacheExpiration[gameID] = time.Now().Add(Cache_Expiration).Unix()

		}
		data, err := simplejson.NewJson([]byte(filelist))
		if err != nil {
			errorResponse(c, err)
			return
		}
		c.JSON(200, data)
	})

	r.POST("/gameResList", func(c *gin.Context) {
		print("post gamereslist")
		if !validClientIP(c) {
			errorResponse(c, "此ip地址不允许执行操作，请联系管理员")
			return
		}
		print("post gamereslist2")

		gameID := c.PostForm("gameID")
		filelist := c.PostForm("filelist")
		if gameID == "" || filelist == "" {
			errorResponse(c, "参数无效")
			return
		}
		print(gameID, filelist)
		db := mysql.DB()
		rows, err := db.Query("call sp_updateGameResList(?,?)", gameID, filelist)
		if err != nil {
			errorResponse(c, err)
			return
		}
		defer rows.Close()
		c.Status(200)
	})
}
