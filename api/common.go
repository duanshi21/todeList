package api

import (
	"encoding/json"
	"todoList/pkg/ctl"
	"todoList/pkg/e"
)

func ErrorResponse(err error) *ctl.TrackedErrorResponse {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "JSON类型不匹配")
	}
	return ctl.RespError(err, "参数错误", e.InvalidParams)
}
