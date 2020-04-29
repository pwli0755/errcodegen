// Code generated by github.com/pwli0755/errcodegen DO NOT EDIT
// Package sample const code comment msg
package sample

// messages get msg from const comment
var messages = map[int]string{

	ERROR:                       "内部错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "鉴权相关失败",
	ERROR_EXIST_TAG:             "标签已存在",
	INVALID_PARAMS:              "请求参数错误",
}

// GetMsg get error msg
func GetMsg(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "UNKNOWN ERROR"
}
