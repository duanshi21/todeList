package ctl

import "todoList/pkg/e"

// Response 最基础的返回resp
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带数组形式返回
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// TokenData 用户登录专属的，带用户信息和token的返回resp
type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}

// TrackedErrorResponse 带有追踪信息的错误返回
type TrackedErrorResponse struct {
	Response
	TraceId string `json:"trace_id"`
}

// RespSuccess 返回成功的函数
func RespSuccess(data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
}

// RespError 错误返回
func RespError(err error, data string, code ...int) *TrackedErrorResponse {
	status := e.Error
	if code != nil {
		status = code[0]
	}

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		// TrackId:  // TODO 加上track id
	}

	return r
}
