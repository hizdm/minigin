package main

import (
	"fmt"
	"minigin/dao"
	"minigin/library/logging"
	"minigin/library/redis"
	"minigin/library/setting"
	"minigin/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	// 配置初始化
	setting.Setup()
	// 日志初始化
	logging.Setup()
	// 数据库初始化
	dao.SetUp()
	redis.Setup()
}
func main() {
	// 设置gin模式
	gin.SetMode(setting.ServerSetting.RunMode)
	r := routers.InitRouter()

	if err := r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)); err != nil {
		logging.Fatal("minigin startup failed, err: %v", err)
	}
}
