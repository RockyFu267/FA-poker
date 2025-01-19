package cardFunc

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
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
	"铁头娃",
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
func HandWinRateSimulationDemo01(input HandConfig) error {
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
	if len(input.HandCardList) > 0 {
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
	// 统计获得胜利最多的玩家
	mostWinPlayer := make(map[string]int)
	type winnerKV struct {
		Key   string
		Value int
	}
	var mostWinPlayerSlice []winnerKV
	//统计获得胜利做的手牌
	mostWinHand := make(map[HandCard]int)
	type hadnKV struct {
		Key   HandCard
		Value int
	}
	//统计真实胜率
	allWinRealRate := make(map[string]int)
	type realWinRateKV struct {
		Key   string
		Value float64
	}

	var mostWinrHandSlice []hadnKV
	var allHandRateSlice []realWinRateKV
	//统计平局的次数
	var tieCount int

	//所有被发出来的手牌统计
	allHandListOrigin := make(map[string]int)

	// 模拟牌局
	for i := 0; i < input.RoundNumber; i++ {
		winners, temphandlist := shuffleJudgeDemo01(players, input.HandCardList)
		for k, v := range temphandlist {
			allHandListOrigin[k] = allHandListOrigin[k] + v
		}
		if len(winners) > 1 {
			// fmt.Println("出现了多个玩家同时获得胜利的情况") //debug
			tieCount++
			continue
		}

		for _, v := range winners {
			// fmt.Println("--AAA---" + strconv.Itoa(i+1) + "---AAA---")
			// fmt.Println(k, v) //debug
			//统计获得胜利最多的玩家
			mostWinPlayer[v.ID]++
			//统计获得胜利做的手牌
			mostWinHand[v.Hand]++
			if v.Hand.HandCard[0].Suit == v.Hand.HandCard[1].Suit {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()+"s"]++
				continue
			}
			if v.Hand.HandCard[0].Suit != v.Hand.HandCard[1].Suit && v.Hand.HandCard[0].Rank != v.Hand.HandCard[1].Rank {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()+"o"]++
				continue
			}
			if v.Hand.HandCard[0].Rank == v.Hand.HandCard[1].Rank {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()]++
				continue
			}
		}
	}

	for k, v := range mostWinPlayer {
		mostWinPlayerSlice = append(mostWinPlayerSlice, winnerKV{k, v})
	}
	for k, v := range mostWinHand {
		mostWinrHandSlice = append(mostWinrHandSlice, hadnKV{k, v})
	}
	for k, v := range allWinRealRate {
		tempRate := float64(v) / float64(allHandListOrigin[k])
		allHandRateSlice = append(allHandRateSlice, realWinRateKV{k, tempRate})
	}

	//输出排序结果
	sort.Slice(mostWinPlayerSlice, func(i, j int) bool {
		return mostWinPlayerSlice[i].Value > mostWinPlayerSlice[j].Value
	})
	sort.Slice(mostWinrHandSlice, func(i, j int) bool { //所有获得胜利的具体手牌
		return mostWinrHandSlice[i].Value > mostWinrHandSlice[j].Value
	})
	sort.Slice(allHandRateSlice, func(i, j int) bool { //所有手牌的胜率
		return allHandRateSlice[i].Value > allHandRateSlice[j].Value
	})
	// 输出结果
	fmt.Println("玩家ID和对应ID:") //debug
	for k, v := range players {
		fmt.Println(k, v.ID)
	}
	// fmt.Println(mostWinPlayer) //debug
	// fmt.Println(mostWinHand) //debug
	fmt.Println("玩家胜利次数排序如下：") //debug
	for i := 0; i < len(mostWinPlayerSlice); i++ {
		fmt.Println(mostWinPlayerSlice[i].Key, mostWinPlayerSlice[i].Value)
	}
	if len(input.HandCardList) > 0 {
		n := len(mostWinrHandSlice)
		if n > 10 {
			n = 10
		}
		fmt.Println("胜率位于前列的手牌组合:")
		for i := 0; i < n; i++ { //输出具体的卡牌
			fmt.Println(mostWinrHandSlice[i].Key.HandCard[0].CardTranslate(), mostWinrHandSlice[i].Key.HandCard[1].CardTranslate(), mostWinrHandSlice[i].Value)
		}
	} else {
		fmt.Println("169组so组合的胜率排序如下：")
		for i := 0; i < len(allHandRateSlice); i++ { //输出所有手牌的胜率
			fmt.Println(allHandRateSlice[i].Key, allHandRateSlice[i].Value)
		}
	}
	fmt.Println("平局次数：", tieCount)
	return nil
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
	// 统计获得胜利最多的玩家
	mostWinPlayer := make(map[string]int)
	type winnerKV struct {
		Key   string
		Value int
	}
	var mostWinPlayerSlice []winnerKV
	//统计获得胜利做的手牌
	mostWinHand := make(map[HandCard]int)
	type hadnKV struct {
		Key   HandCard
		Value int
	}
	var mostWinrHandSlice []hadnKV
	//统计平局的次数
	var tieCount int

	// 模拟牌局
	for i := 0; i < input.RoundNumber; i++ {
		winners := shuffleJudgeDemo(players, input.HandCardList)
		if len(winners) > 1 {
			fmt.Println("出现了多个玩家同时获得胜利的情况")
			tieCount++
			continue
		}
		for _, v := range winners {
			// fmt.Println("--AAA---" + strconv.Itoa(i+1) + "---AAA---")
			// fmt.Println(k, v) //debug
			//统计获得胜利最多的玩家
			mostWinPlayer[v.ID]++
			//统计获得胜利做的手牌
			mostWinHand[v.Hand]++
		}
	}

	for k, v := range mostWinPlayer {
		mostWinPlayerSlice = append(mostWinPlayerSlice, winnerKV{k, v})
	}
	for k, v := range mostWinHand {
		mostWinrHandSlice = append(mostWinrHandSlice, hadnKV{k, v})
	}
	//输出排序结果
	sort.Slice(mostWinPlayerSlice, func(i, j int) bool {
		return mostWinPlayerSlice[i].Value > mostWinPlayerSlice[j].Value
	})
	sort.Slice(mostWinrHandSlice, func(i, j int) bool {
		return mostWinrHandSlice[i].Value > mostWinrHandSlice[j].Value
	})
	for k, v := range players {
		fmt.Println(k, v.ID)
	}
	// fmt.Println(mostWinPlayer) //debug
	// fmt.Println(mostWinHand) //debug
	for i := 0; i < input.PlayerNumber; i++ {
		fmt.Println(mostWinPlayerSlice[i].Key, mostWinPlayerSlice[i].Value)
	}
	for i := 0; i < len(mostWinrHandSlice); i++ {
		fmt.Println(mostWinrHandSlice[i].Key.HandCard[0].CardTranslate(), mostWinrHandSlice[i].Key.HandCard[1].CardTranslate(), mostWinrHandSlice[i].Value)
	}
	fmt.Println("平局次数：", tieCount)

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