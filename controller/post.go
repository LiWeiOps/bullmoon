package controller

import (
	"bullmoon/logic"
	"bullmoon/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(ctx *gin.Context) {
	//接收参数
	post := new(models.Post)
	if err := ctx.ShouldBindJSON(post); err != nil {
		zap.L().Error("post should bind error", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//获取userid
	var err error
	post.AuthorID, err = getCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	//创建post
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost(post) error", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(ctx, nil)
}

func GetPostDetailHandler(ctx *gin.Context) {
	// 获取post_id
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	// 根据id查询post
	data, err := logic.GetPostDetail(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail(id) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}

func GetPostListHandler(ctx *gin.Context) {
	page, size := getPageInfo(ctx)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}

// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照分数排序
// 1. 获取参数
// 2. 去redis查询id列表
// 3. 根据id去数据库查询帖子详情

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 查询帖子列表相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(ctx *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 invalid params", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}
