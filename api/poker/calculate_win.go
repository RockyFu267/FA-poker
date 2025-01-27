package poker

import (
	"aquila/cardFunc"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PokerCalculate struct {
}

// SimulatePokerHandsApi 模拟德州扑克手牌胜率计算
func (p PokerCalculate) SimulatePokerHandsApi(ctx *gin.Context) {
	var req cardFunc.HandConfig
	err := ctx.ShouldBind(&req)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "参数绑定失败："+err.Error())
		return
	}

	// 调用模拟计算函数
	result, err := cardFunc.HandWinRateSimulationWeb01(req)
	if err != nil {
		utils.Fail(ctx, "计算失败："+err.Error())
		return
	}

	utils.Success(ctx, result)
}

// GetSimulationConfigApi 获取模拟配置
func (p PokerCalculate) GetSimulationConfigApi(ctx *gin.Context) {
	configPath := ctx.Query("configPath")
	if configPath == "" {
		utils.Fail(ctx, "配置路径不能为空")
		return
	}

	config, err := cardFunc.ReadConfig(configPath)
	if err != nil {
		utils.Fail(ctx, "读取配置失败："+err.Error())
		return
	}

	utils.Success(ctx, config)
}

// SaveSimulationConfigApi 保存模拟配置
func (p PokerCalculate) SaveSimulationConfigApi(ctx *gin.Context) {
    var req struct {
        ConfigPath string             `json:"configPath"`
        Config     cardFunc.HandConfig `json:"config"`
    }
    
    err := ctx.ShouldBind(&req)
    if err != nil {
        utils.Fail(ctx, "参数绑定失败："+err.Error())
        return
    }

    // 这里需要实现配置保存的逻辑
    // TODO: 实现配置文件的保存功能
    
    utils.Success(ctx, nil)
}

// ValidateHandCardsApi 验证手牌合法性
func (p PokerCalculate) ValidateHandCardsApi(ctx *gin.Context) {
    var req struct {
        HandCards []cardFunc.HandCard `json:"handCards"`
    }
    
    err := ctx.ShouldBind(&req)
    if err != nil {
        utils.Fail(ctx, "参数绑定失败："+err.Error())
        return
    }

    // 验证手牌是否合法
    var cardMap = make(map[cardFunc.Card]bool)
    for i, hand := range req.HandCards {
        for _, card := range hand.HandCard {
            if card.Rank < 2 || card.Rank > 14 {
                utils.Fail(ctx, fmt.Sprintf("第%d副手牌的点数不合法", i+1))
                return
            }
            if card.Suit != "黑桃" && card.Suit != "红桃" && card.Suit != "方片" && card.Suit != "梅花" {
                utils.Fail(ctx, fmt.Sprintf("第%d副手牌的花色不合法", i+1))
                return
            }
            if cardMap[card] {
                utils.Fail(ctx, fmt.Sprintf("第%d副手牌中包含重复的牌", i+1))
                return
            }
            cardMap[card] = true
        }
    }

    utils.Success(ctx, nil)
}
