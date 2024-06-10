package routes

import (
	"bullmoon/controller"
	_ "bullmoon/docs" // 千万不要忘了导入把你上一步生成的docs
	"bullmoon/logger"
	"bullmoon/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/post", controller.GetPostListHandler)
		v1.GET("/posts2", controller.GetPostListHandler2)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "认证Token成功")
	})
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
