package main

import (
	"fmt"
	"mini-services/pkg/mgo"
	"mini-services/pkg/rabbit"
	"mini-services/pkg/redis"
	"mini-services/pkg/setting"
	"mini-services/router"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	redis.Setup()
	mgo.Setup()
	rabbit.Setup()
}

func main() {
	r := gin.Default()
	gin.SetMode(setting.ServerSetting.RunMode) // 发版本模式

	router.InitRouter(r)

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	r.Run(endPoint) // listen and serve on 0.0.0.0:8080
}
