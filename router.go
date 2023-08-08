package main

import (
	"github.com/abuziming/dousheng_demo/config"
	"github.com/abuziming/dousheng_demo/controller"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	// public directory is used to serve static resources
	r.Static("/static", config.Global.Path.StaticSourcePath)

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                                      // 这个 token 得特殊处理
	apiRouter.GET("/user/", controller.JwtAuth(), controller.UserInfoHandler)     // 获取用户信息
	apiRouter.POST("/user/register/", controller.RegisterHandler)                 // 登录
	apiRouter.POST("/user/login/", controller.LoginHandler)                       // 注册
	apiRouter.POST("/publish/action/", controller.JwtAuth(), controller.Publish)  // 发布视频
	apiRouter.GET("/publish/list/", controller.JwtAuth(), controller.PublishList) // 发布列表

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.JwtAuth(), controller.FavoriteAction) // 点赞
	apiRouter.GET("/favorite/list/", controller.JwtAuth(), controller.FavoriteList)      // 喜欢列表
	apiRouter.POST("/comment/action/", controller.JwtAuth(), controller.CommentAction)   // 评论
	apiRouter.GET("/comment/list/", controller.JwtAuth(), controller.CommentList)        // 评论列表

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
	return r
}
