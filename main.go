package main

import (
	"fmt"
	"min-services/pkg/mgo"
	"min-services/pkg/rabbit"
	"min-services/pkg/redis"
	"min-services/pkg/setting"
	"min-services/router"

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
