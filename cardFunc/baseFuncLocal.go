package cardFunc

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

var entertainmentIDs = []string{
	"扑克小鲁班",
	"牌桌范德彪",
	"赌圣他二舅",
	"气质男孩",
	"顺子制造机",
	"葫芦小王子",
	"皇家同花顺饲养员",
	//以下都是台球名人的名字，姓都改成了傅，也可以是其他体育明行的名字
	"傅沙利文",
	"傅尔摩斯",
	"傅俊辉",
	"傅尔比",
	"傅金斯",
	"傅得利",
	"傅廉姆斯",
	"傅马赫",
	"傅俊辉",
}

// 单机模式功能1：模拟牌局统计胜率、牌型
func HandWinRateSimulation(input HandConfig) error {
	var cardMap = make(map[Card]bool)
	if input.PlayerNumber < 2 || input.PlayerNumber > 10 {
		return fmt.Errorf("playNumber must 大于等于 2，小于等于 10")
	}
	if input.RoundNumber < 1 || input.RoundNumber > 100000 {
		return fmt.Errorf("roundNumber must 大于等于 1,小于等于 100000")
	}
	if len(input.HandCardList) > input.PlayerNumber {
		return fmt.Errorf("playNumber must 大于等于 HandCardList长度")
	}
	for k, v := range input.HandCardList {
		if v.HandCard[0].Rank < 2 || v.HandCard[0].Rank > 14 || v.HandCard[1].Rank < 2 || v.HandCard[1].Rank > 14 {
			return fmt.Errorf("第" + strconv.Itoa(k+1) + "个元素的Rank值有问题，正确范围2-14之间")
		}
		if v.HandCard[0].Suit != "黑桃" && v.HandCard[0].Suit != "红桃" && v.HandCard[0].Suit != "方片" && v.HandCard[0].Suit != "梅花" {
			return fmt.Errorf("第" + strconv.Itoa(k+1) + "个元素的Suit值有问题,花色范围只在 黑桃  红桃  方片  梅花 中选择")
		}
		if ok := cardMap[v.HandCard[0]]; ok {
			return fmt.Errorf("第" + strconv.Itoa(k+1) + "个元素的手牌有问题，不能有重复的牌")
		} else {
			cardMap[v.HandCard[0]] = true
		}
		if ok := cardMap[v.HandCard[1]]; ok {
			return fmt.Errorf("第" + strconv.Itoa(k+1) + "个元素的手牌有问题，不能有重复的牌")
		} else {
			cardMap[v.HandCard[1]] = true
		}
		// handTemp := handSorting(v.HandCard)
		input.HandCardList[k].sortTwoCards()
	}
	// 初始化玩家
	players := make([]Players, input.PlayerNumber)

	// 初始化随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用独立的随机数生成器

	// 随机分配娱乐ID
	usedIDs := make(map[string]bool) // 记录已分配的ID，避免重复

	for i := 0; i < input.PlayerNumber; i++ {
		for {
			randomIndex := rng.Intn(len(entertainmentIDs)) // 使用本地生成器生成随机索引
			randomID := entertainmentIDs[randomIndex]      // 选择对应的ID
			if !usedIDs[randomID] {                        // 检查是否已经被使用
				players[i].ID = randomID // 分配给玩家
				usedIDs[randomID] = true // 标记为已使用
				break                    // 跳出循环，分配下一个玩家
			}
		}
	}
	// //统计获得胜利最多的玩家
	mostWinPlayer := make(map[string]int)
	// //统计获得胜利做的手牌
	mostWinPlayerHand := make(map[HandCard]int)
	// 模拟牌局
	for i := 0; i < input.RoundNumber; i++ {
		winners := shuffleJudgeDemo(players, input.HandCardList)
		for k, v := range winners {
			// fmt.Println("--AAA---" + strconv.Itoa(i+1) + "---AAA---")
			fmt.Println(k, v)
			//统计获得胜利最多的玩家
			mostWinPlayer[v.ID]++
			//统计获得胜利做的手牌
			mostWinPlayerHand[v.Hand]++
		}
	}
	for k, v := range players {
		fmt.Println(k, v.ID)
	}
	fmt.Println(mostWinPlayer)
	fmt.Println(mostWinPlayerHand)

	return nil
}

// ReadConfig 读取外部配置文件
func ReadConfig(dir string) (HandConfig, error) {

	// pwdPath, err := os.Getwd()
	// if err != nil {
	// 	return HandConfig{}, fmt.Errorf("get dirpath error: %v", err)
	// }
	// 使用os.Open打开文件，它返回一个文件指针和可能的错误
	file, err := os.Open(dir)
	if err != nil {
		return HandConfig{}, fmt.Errorf("open config file error: %v", err)
	}
	// 使用defer关键字确保文件最终会被关闭，避免资源泄露
	defer file.Close()

	// 读取文件内容到字节切片，这里使用io.ReadAll来替代原ioutil.ReadFile的功能
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return HandConfig{}, fmt.Errorf("read config file content error: %v", err)
	}

	var confDemo HandConfig
	err = yaml.Unmarshal(yamlFile, &confDemo)
	if err != nil {
		return HandConfig{}, fmt.Errorf("unmarshal Config Error: %v", err)
	}

	return confDemo, nil

}
