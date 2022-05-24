package sonbusiness

import (
	"minigin/library/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 举个例子：子业务类型
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": "路由根据业务类型拆分到不同文件中示例",
	})
}
