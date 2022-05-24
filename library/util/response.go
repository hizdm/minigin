package util

import (
	"encoding/json"
	"fmt"
	"minigin/library/logging"
)

// Response Map2Json
func Response(code int64, msg string, data interface{}) string {
	jsonData := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	str, err := json.Marshal(jsonData)

	if err != nil {
		logging.Error(fmt.Sprintf("util.Response json.Marshal error, err: %v", err))
	}

	return string(str)
}
