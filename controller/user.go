package controller

import (
	"bullmoon/dao/mysql"
	"bullmoon/logic"
	"bullmoon/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	// 1.获取提交参数
	var p = new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("Login with invalid param", zap.Error(err))
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			zap.L().Error("SignUp with invalid param", zap.Any("err", errs.Translate(trans)))
			ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}
	// 2.业务逻辑
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Warn("logic SignUp err", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(ctx, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(ctx, CodeInvalidPassword)
			return
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 登录成功生成token
	// 3.返回响应
	ResponseSuccess(ctx, gin.H{
		// 防止前端id失真问题，将int64的数值转换为string类型
		//前端 id值范围 1<<53-1,后端id值范围 1<<63-1
		"user_id":  fmt.Sprintf("%d", user.UserId),
		"username": user.Username,
		"token":    user.Token,
	})
}

func SignUpHandler(ctx *gin.Context) {
	// 1.获取提交参数
	var p = new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}
	fmt.Println(p)
	// 2.业务逻辑处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Warn("logic SignUp err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
			return
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(ctx, nil)
}
