package controller

import (
	"bullmoon/logic"
	"bullmoon/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(ctx *gin.Context) {
	// 参数校验
	p := new(models.ParamVotedData)
	if err := ctx.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("PostVoteController, ctx.ShouldBindJSON(p) error", zap.Error(err))
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		zap.L().Error("PostVoteController, ctx.ShouldBindJSON(p) error",
			zap.Any("removeTopStruct(errs.Translate(trans))", errData))
		ResponseErrorWithMsg(ctx, CodeInvalidParam, errData)
		return
	}
	userid, err := getCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	err = logic.PostVote(userid, p)
	if err != nil {
		zap.L().Error("logic.PostVote(userid, p) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, nil)
}
