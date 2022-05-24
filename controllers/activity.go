package controllers

import (
	"fmt"
	"math/rand"
	"minigin/library/e"
	"minigin/library/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var seckillNumber int64 = 10
var setKey = "succ_user_no_set_6"

// 秒杀
func Seckill(c *gin.Context) {
	// no := c.PostForm("no")
	// no := c.Query("no")
	no := "no:" + GenerateUnique()

	// 黑产
	// 时间
	number := redis.SCard(setKey)
	if number >= seckillNumber {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  "本轮商品已被秒杀完毕，如秒杀成功，请查看个人账户",
			"data": "",
		})

		return
	}

	if no != "" {
		result := redis.SAdd(setKey, no)
		if result == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code": e.SUCCESS,
				"msg":  "恭喜你秒杀成功",
				"data": fmt.Sprintf("您的优惠券编号：%s 请牢记！", GenerateUnique()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": e.SUCCESS,
				"msg":  "已经秒杀成功",
				"data": "",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ILLEGAL_REQUEST,
			"msg":  e.GetMsg(e.ILLEGAL_REQUEST),
			"data": "",
		})
	}
}

// 成功用户列表
func SeckillList(c *gin.Context) {
	members := redis.SMembers(setKey)

	for _, v := range members {
		fmt.Printf("%s\r\n", v)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": members,
	})
}

func GenerateUnique() string {
	var r = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0")
	uniq := make([]rune, 20)
	for i := range uniq {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i<<20))
		uniq[i] = r[rand.Intn(len(r))]
	}
	return string(uniq)
}
