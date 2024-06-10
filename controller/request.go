package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "UserId"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(ctx *gin.Context) (userId int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
	}
	userId, ok = uid.(int64)
	return userId, nil
}

func getPageInfo(ctx *gin.Context) (page int64, size int64) {
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")
	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
