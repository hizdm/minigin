package routers

import (
	"fmt"
	"minigin/controllers"
	"minigin/library/middleware"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func InitRouter() *gin.Engine {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 静态资源存放地址
	r.Static("/static", "static")

	// 模版存放地址
	r.LoadHTMLGlob("templates/*")

	// Controllers
	r.GET("/", controllers.IndexHandler)
	r.GET("/articles", controllers.FindArticle) // 文章列表
	r.POST("/article", controllers.CreateArticle)

	// Redis
	r.GET("/redis/get", controllers.RedisGet)

	// 路由根据业务类型拆分到不同文件中
	InitSonRouter(r)

	// 秒杀活动实例
	r.GET("/activity/seckill", controllers.Seckill)
	r.GET("/activity/seckilllist", controllers.SeckillList)

	// JWT
	r.GET("/jwt/generate", controllers.JwtGenerate)
	apiV1 := r.Group("jwt")
	apiV1.Use(middleware.JWT())
	{
		apiV1.GET("/index", controllers.JwtIndex)
	}

	return r
}
