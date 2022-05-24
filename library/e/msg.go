package e

var Msg = map[int]string{
	SUCCESS:         "success",
	ERROR:           "fail",
	ILLEGAL_REQUEST: "非法请求",

	REDIS_GET_ERROR:  "redis get error",
	REDIS_SET_ERROR:  "redis set error",
	REDIS_TTL_ERROR:  "redis ttl error",
	REDIS_SADD_ERROR: "redis sadd error",

	JSON_MARSHAL_ERROR: "json.Marshal(value) error",

	AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	AUTH_TOKEN_GENERATE_ERROR: "Token生成失败",
	AUTH_ERROR:                "Token错误",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]

	if ok {
		return msg
	}

	return Msg[ERROR]
}
