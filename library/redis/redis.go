package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"minigin/library/e"
	"minigin/library/logging"
	"minigin/library/setting"
	"minigin/library/util"
	"time"

	"github.com/go-redis/redis"
)

var RedisConn *redis.Client

func Setup() {
	// 连接服务器
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,     // Addr
		Password: setting.RedisSetting.Password, // Password
		DB:       0,                             // use default DB
	})

	// 心跳
	pong, err := RedisConn.Ping().Result()
	log.Println(pong, err) // Output: PONG <nil>
}

// --------- String ----------
// Redis Get
func Get(key string) string {
	value, err := RedisConn.Get(key).Result()
	if err != nil {
		logging.Error(util.Response(e.REDIS_GET_ERROR, fmt.Sprintf("redis.Get error, err: %v", err), ""))
	}

	return value
}

// Redis Set
func Set(key string, value interface{}, expiration time.Duration) bool {
	value, err := json.Marshal(value)
	if err != nil {
		logging.Error(util.Response(e.JSON_MARSHAL_ERROR, fmt.Sprintf("redis.Set %s, err: %v", e.GetMsg(e.JSON_MARSHAL_ERROR), err), ""))
		return false
	}

	// todo 时间
	err = RedisConn.Set(key, value, 1000000000*expiration).Err()

	if err != nil {
		logging.Error(util.Response(e.REDIS_SET_ERROR, fmt.Sprintf("redis.Set error, err: %v", err), ""))
		return false
	}

	return true
}

// Redis TTL
func TTL(key string) int {
	expiration, err := RedisConn.TTL("key").Result()

	if err != nil {
		logging.Error(util.Response(e.REDIS_TTL_ERROR, fmt.Sprintf("redis.TTL error, err: %v", err), ""))
	}

	return int(expiration)
}

// --------- Set ----------
// SAdd
func SAdd(key string, value string) int64 {
	result, err := RedisConn.SAdd(key, value).Result()

	if err != nil {
		logging.Error(util.Response(e.REDIS_SADD_ERROR, fmt.Sprintf("redis.Sadd error, err: %v", err), ""))
	}

	return result
}

// SMembers
func SMembers(key string) []string {
	members, err := RedisConn.SMembers(key).Result()

	if err != nil {
		fmt.Println(err)
	}

	return members
}

// SCard
func SCard(key string) int64 {
	count, err := RedisConn.SCard(key).Result()

	if err != nil {
		fmt.Println(err)
	}

	return count
}

// Hash
