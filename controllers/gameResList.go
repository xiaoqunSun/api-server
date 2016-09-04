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
}
