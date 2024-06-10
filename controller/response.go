package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义错误响应码，并封装返回响应的方法
/*
{
	"code": 1000	//程序中的错误码
	"mes": xx		//错误码对应的提示信息
	"data": xxx		//返回的响应的数据
}
*/

type responseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(ctx *gin.Context, code ResCode) {
	ctx.JSON(http.StatusOK, &responseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &responseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResCode, msg interface{}) {
	ctx.JSON(http.StatusOK, &responseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
