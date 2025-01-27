package router

import (
	"aquila/api/system"
	"aquila/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("menu", middleware.AuthMiddleWare())
	menu := system.Menu{}
	resisterRouter.POST("", menu.CreateMenuApi)          // 创建菜单
	resisterRouter.GET("", menu.GetMenuApi)             // 获取单个菜单
	resisterRouter.GET("list", menu.GetMenuAllApi)      // 获取所有菜单树形结构
	resisterRouter.DELETE("", menu.DeleteMenuApi)       // 删除菜单
	resisterRouter.PUT("", menu.UpdateMenuApi)          // 更新菜单
	resisterRouter.GET("page", menu.GetMenuPageApi)     // 分页查询菜单
}
