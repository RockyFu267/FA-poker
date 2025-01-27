package router

import (
	"aquila/api/bot"

	"github.com/gin-gonic/gin"
)

func InitBotRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("bot")
	qianxun := bot.Qianxun{}
	resisterRouter.POST("/", qianxun.HandleWechatMsgApi)
	resisterRouter.GET("/shJob", qianxun.FetchJobAnnouncements)
	resisterRouter.GET("/shBird", qianxun.FetchBirdReport)
}
