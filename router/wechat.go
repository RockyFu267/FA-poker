package router

import (
	"aquila/api/wechat"

	"github.com/gin-gonic/gin"
)

func InitWechatRouter(Router *gin.RouterGroup) {
	WechatRouter := Router.Group("wechat")
	{
		login := wechat.Login{}
		WechatRouter.POST("wxlogin", login.WxLoginApi)
		WechatRouter.POST("updateUser", login.UpdateUserApi)
	}
	{
		group := wechat.Group{}
		WechatRouter.POST("group/create", group.CreateGroupApi)
		WechatRouter.GET("group/get", group.GetGroupApi)
		WechatRouter.GET("group/getPage", group.GetGroupPageApi)
		WechatRouter.POST("group/update", group.UpdateGroupApi)
		WechatRouter.POST("group/delete", group.DeleteGroupApi)
	}
	{
		groupType := wechat.GroupType{}
		WechatRouter.POST("groupType/create", groupType.CreateGroupTypeApi)
		WechatRouter.GET("groupType/get", groupType.GetGroupTypeApi)
		WechatRouter.GET("groupType/getPage", groupType.GetGroupTypePageApi)
		WechatRouter.POST("groupType/update", groupType.UpdateGroupTypeApi)
		WechatRouter.POST("groupType/delete", groupType.DeleteGroupTypeApi)
	}
	{
		banner := wechat.Banner{}
		WechatRouter.POST("banner/create", banner.CreateBannerApi)
		WechatRouter.GET("banner/get", banner.GetBannerApi)
		WechatRouter.GET("banner/getPage", banner.GetBannerPageApi)
		WechatRouter.POST("banner/update", banner.UpdateBannerApi)
		WechatRouter.POST("banner/delete", banner.DeleteBannerApi)
	}
	{
		article := wechat.Article{}
		WechatRouter.POST("article/create", article.CreateArticleApi)
		WechatRouter.GET("article/get", article.GetArticleApi)
		WechatRouter.GET("article/getPage", article.GetArticlePageApi)
		WechatRouter.POST("article/update", article.UpdateArticleApi)
		WechatRouter.GET("article/delete", article.DeleteArticleApi)
	}

	{
		articleType := wechat.ArticleType{}
		WechatRouter.POST("articleType/create", articleType.CreateArticleTypeApi)
		WechatRouter.GET("articleType/get", articleType.GetArticleTypeApi)
		WechatRouter.GET("articleType/getList", articleType.GetArticleTypeListApi)
		WechatRouter.POST("articleType/update", articleType.UpdateArticleTypeApi)
		WechatRouter.GET("articleType/delete", articleType.DeleteArticleTypeApi)
	}
}
