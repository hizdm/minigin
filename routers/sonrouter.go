package routers

import (
	"minigin/controllers/sonbusiness"

	"github.com/gin-gonic/gin"
)

// 按照子业务类型进行路由文件存储
func InitSonRouter(route *gin.Engine) {
	sonBusiness := route.Group("sonbusiness")
	{
		sonBusiness.GET("/index", sonbusiness.Index)
	}
}
