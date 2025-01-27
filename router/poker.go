package router

import (
	"aquila/api/poker"
	"github.com/gin-gonic/gin"
)

func InitPokerRouter(Router *gin.RouterGroup) {
	pokerRouter := Router.Group("poker")
	calculate := poker.PokerCalculate{}

	// 德州扑克胜率模拟计算相关路由
	pokerRouter.POST("/simulate", calculate.SimulatePokerHandsApi)     // 模拟计算胜率
	pokerRouter.GET("/config", calculate.GetSimulationConfigApi)       // 获取模拟配置
	pokerRouter.POST("/config", calculate.SaveSimulationConfigApi)     // 保存模拟配置
	pokerRouter.POST("/validate", calculate.ValidateHandCardsApi)      // 验证手牌合法性
}
