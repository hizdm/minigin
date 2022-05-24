package controllers

import (
	"fmt"
	"minigin/library/e"
	"minigin/library/middleware"
	"minigin/library/redis"
	"minigin/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Contrller 创建文章
func CreateArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)
	result, err := article.CreateArticle()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": result,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": result,
	})
}

// Find
func FindArticle(c *gin.Context) {
	var article models.Article
	result, err := article.FindArticle()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": result,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": result,
	})
}

// Redis
func RedisGet(c *gin.Context) {

	scoreMap := make(map[string]int, 2)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// a := redis.Set("name", scoreMap, 1000)
	// fmt.Println("----------")
	// fmt.Println(a)
	// result, err := redis.Get("name")

	a := redis.Set("name", scoreMap, 10000000000)

	fmt.Println(a)
	result := redis.Get("name")

	b := redis.TTL("name")
	fmt.Println(b)

	fmt.Println("-----------")

	c.JSON(http.StatusOK, gin.H{
		"code": 300,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": result,
	})

}

// JWT
func JwtGenerate(c *gin.Context) {
	token, err := middleware.GenerateToken("zhang", "123456")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.AUTH_TOKEN_GENERATE_ERROR,
			"msg":  e.GetMsg(e.AUTH_TOKEN_GENERATE_ERROR),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": token,
	})
}

func JwtIndex(c *gin.Context) {
	hToken := c.GetHeader("Authorization")
	fmt.Println(hToken)
	bearerLength := len("Bearer ")
	token := strings.TrimSpace(hToken[bearerLength:])
	fmt.Println(token)

	a, err := middleware.ParseToken(token)
	fmt.Println(a.Username)
	fmt.Println(a.Issuer)
	fmt.Println(err)
	ua2 := c.Request.Header.Get("Authorization")
	fmt.Println(ua2)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": "JWT Index Test",
	})
}
