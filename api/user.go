package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoList/pkg/utils"
	"todoList/service"
	"todoList/types"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserRegister(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LoggerObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// UserLoginHandler 用户登录
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LoggerObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
