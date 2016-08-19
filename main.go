package main

import (
	"io/ioutil"
	"log"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqunSun/api-server/controllers"
	"github.com/xiaoqunSun/api-server/mysql"
)

func main() {
	content, err := ioutil.ReadFile("etc/config.json")
	if err != nil {
		log.Println("cannot open et/config.json", err)
		return
	}
	config, err := simplejson.NewJson(content)
	if err != nil {
		log.Println("etc/config is not a json", err)
		return
	}

	dsn, err := config.Get("mysql_dsn").String()
	if err != nil {
		log.Println("config get mysql_dsn error", err)
		return
	}
	//init mysql
	err = mysql.Init(dsn)
	if err != nil {
		log.Println("mysql open error", err)
		return
	}
	r := gin.Default()

	r.Static("/static", "./static")
	controllers.HandlerVersions(r)
	controllers.HandlerAccount(r)
	controllers.HandlerServerAddr(r)
	r.Run(":8080")
}
