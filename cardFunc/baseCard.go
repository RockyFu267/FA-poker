package cardFunc

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// CardTranslate 转换卡牌显示值
func (p Card) CardTranslate() string {
	suits := map[string]string{
		"黑桃": "♠",
		"红桃": "♥",
		"梅花": "♣",
		"方片": "♦",
	}

	ranks := map[int]string{
		14: "A", 13: "K", 12: "Q", 11: "J",
		10: "10", 9: "9", 8: "8", 7: "7",
		6: "6", 5: "5", 4: "4", 3: "3", 2: "2",
	}

	suitSymbol, suitExists := suits[p.Suit]
	rankSymbol, rankExists := ranks[p.Rank]

	if suitExists && rankExists {
		return suitSymbol + rankSymbol
	}
	return "fuck card"
}

// CardTranslate 转换卡牌rank的显示值
func (p Card) CardRankTranslate() string {

	ranks := map[int]string{
		14: "A", 13: "K", 12: "Q", 11: "J",
		10: "10", 9: "9", 8: "8", 7: "7",
		6: "6", 5: "5", 4: "4", 3: "3", 2: "2",
	}

	rankSymbol, rankExists := ranks[p.Rank]

	if rankExists {
		return rankSymbol
	}
	return "fuck card"
}

// ShuffleCard 洗牌
func ShuffleCard() (New52CardList [52]Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化52张牌
	var Card52 = [52]Card{
		{Suit: "黑桃", Rank: 14},
		{Suit: "黑桃", Rank: 2},
		{Suit: "黑桃", Rank: 3},
		{Suit: "黑桃", Rank: 4},
		{Suit: "黑桃", Rank: 5},
		{Suit: "黑桃", Rank: 6},
		{Suit: "黑桃", Rank: 7},
		{Suit: "黑桃", Rank: 8},
		{Suit: "黑桃", Rank: 9},
		{Suit: "黑桃", Rank: 10},
		{Suit: "黑桃", Rank: 11},
		{Suit: "黑桃", Rank: 12},
		{Suit: "黑桃", Rank: 13},
		{Suit: "红桃", Rank: 14},
		{Suit: "红桃", Rank: 2},
		{Suit: "红桃", Rank: 3},
		{Suit: "红桃", Rank: 4},
		{Suit: "红桃", Rank: 5},
		{Suit: "红桃", Rank: 6},
		{Suit: "红桃", Rank: 7},
		{Suit: "红桃", Rank: 8},
		{Suit: "红桃", Rank: 9},
		{Suit: "红桃", Rank: 10},
		{Suit: "红桃", Rank: 11},
		{Suit: "红桃", Rank: 12},
		{Suit: "红桃", Rank: 13},
		{Suit: "梅花", Rank: 14},
		{Suit: "梅花", Rank: 2},
		{Suit: "梅花", Rank: 3},
		{Suit: "梅花", Rank: 4},
		{Suit: "梅花", Rank: 5},
		{Suit: "梅花", Rank: 6},
		{Suit: "梅花", Rank: 7},
		{Suit: "梅花", Rank: 8},
		{Suit: "梅花", Rank: 9},
		{Suit: "梅花", Rank: 10},
		{Suit: "梅花", Rank: 11},
		{Suit: "梅花", Rank: 12},
		{Suit: "梅花", Rank: 13},
		{Suit: "方片", Rank: 14},
		{Suit: "方片", Rank: 2},
		{Suit: "方片", Rank: 3},
		{Suit: "方片", Rank: 4},
		{Suit: "方片", Rank: 5},
		{Suit: "方片", Rank: 6},
		{Suit: "方片", Rank: 7},
		{Suit: "方片", Rank: 8},
		{Suit: "方片", Rank: 9},
		{Suit: "方片", Rank: 10},
		{Suit: "方片", Rank: 11},
		{Suit: "方片", Rank: 12},
		{Suit: "方片", Rank: 13},
	}
	// //洗牌
	// var new52 [52]Card
	// 洗牌
	r.Shuffle(len(Card52), func(i, j int) {
		Card52[i], Card52[j] = Card52[j], Card52[i]
	})
	// b := 0
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// for _, i := range r.Perm(len(Card52)) {
	// 	val := Card52[i]
	// 	// fmt.Println(val)
	// 	// fmt.Println(i)
	// 	new52[b] = val
	// 	b = b + 1
	// }
	//fmt.Println(new52)
	return Card52
}

// GetTopTwoCards 从洗好的牌中取前两张牌，并按指定规则排序后赋值给 HandCard 类型的变量
func GetTopTwoCards(deck [52]Card) HandCard {
	// 取前两张牌
	topTwoCards := [2]Card{deck[0], deck[1]}

	// 排序逻辑
	if topTwoCards[0].Rank < topTwoCards[1].Rank ||
		(topTwoCards[0].Rank == topTwoCards[1].Rank && compareSuits(topTwoCards[0].Suit, topTwoCards[1].Suit)) {
		// 交换两张牌
		topTwoCards[0], topTwoCards[1] = topTwoCards[1], topTwoCards[0]
	}

	// 将排序后的牌赋值给 HandCard
	hand := HandCard{
		HandCard: topTwoCards,
	}

	return hand
}

// compareSuits 比较两张牌的花色，黑桃 > 红桃 > 梅花 > 方片
func compareSuits(suit1, suit2 string) bool {
	order := map[string]int{
		"黑桃": 1,
		"红桃": 2,
		"梅花": 3,
		"方片": 4,
	}
	return order[suit1] > order[suit2]
}

// ShuffleAndRecord 执行洗牌次，将所有组合以及频率放在一个map里并写入文件中。同时要将所有组合
func ShuffleAndRecord(iterations int, filename string) {
	results := make(map[HandCard]int)

	for i := 0; i < iterations; i++ {
		deck := ShuffleCard()
		hand := GetTopTwoCards(deck)
		results[hand]++
	}

	// 将结果按组合频率升序排序  并统计指定组合
	//口袋对出现的次数
	pairCount := 0
	//同花出现次数
	suitedCount := 0
	//同花连张出现次数
	suitedConnectorCount := 0
	//Axs 同色且至少包含一张A的出现的次数
	suitedAceCount := 0
	//Axo 同色且至少包含一张A的出现的次数
	aceCount := 0
	//两张牌都大于等于10出现的次数
	highCardCount := 0

	type kv struct {
		Key   HandCard
		Value int
	}
	var sortedResults []kv
	for k, v := range results {
		sortedResults = append(sortedResults, kv{k, v})
		if k.HandCard[0].Rank == k.HandCard[1].Rank {
			pairCount += v
		}
		if k.HandCard[0].Suit == k.HandCard[1].Suit {
			suitedCount += v
		}
		//要考虑A和2的关系,A的rank是14，实际A和2也是连长
		if k.HandCard[0].Suit == k.HandCard[1].Suit && k.HandCard[0].Rank == k.HandCard[1].Rank+1 || (k.HandCard[0].Rank == 14 && k.HandCard[1].Rank == 2 && k.HandCard[0].Suit == k.HandCard[1].Suit) {
			suitedConnectorCount += v
		}
		if k.HandCard[0].Suit == k.HandCard[1].Suit && (k.HandCard[0].Rank == 14 || k.HandCard[1].Rank == 14) {
			suitedAceCount += v
		}
		if k.HandCard[0].Rank == 14 || k.HandCard[1].Rank == 14 {
			aceCount += v
		}
		if k.HandCard[0].Rank >= 10 && k.HandCard[1].Rank >= 10 {
			highCardCount += v
		}
	}

	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].Value < sortedResults[j].Value
	})

	//统计并打印所有两张牌大小一样的组合，忽略花色
	pairMap := make(map[string]int)
	//统计并打印所有两张牌suit相同的组合
	suitMap := make(map[string]int)
	//统计并打印所有两张牌suit不相同的组合，且rank不相同的组合
	offsuitMap := make(map[string]int)
	for k, _ := range results {
		if k.HandCard[0].Rank == k.HandCard[1].Rank {
			pairMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()] = pairMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()] + 1
		}
		if k.HandCard[0].Suit == k.HandCard[1].Suit {
			suitMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()+"s"] = suitMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()+"s"] + 1
		}
		if k.HandCard[0].Suit != k.HandCard[1].Suit && k.HandCard[0].Rank != k.HandCard[1].Rank {
			offsuitMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()+"o"] = offsuitMap[k.HandCard[0].CardRankTranslate()+k.HandCard[1].CardRankTranslate()+"o"] + 1
		}

	}
	fmt.Println("Two cards of the same list:", len(pairMap), "----\n", pairMap)
	fmt.Println("Two cards of the same suit:", len(suitMap), "----\n", suitMap)
	fmt.Println("Two cards of different suit:", len(offsuitMap), "----\n", offsuitMap)
	// 将结果写入文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(sortedResults); err != nil {
		fmt.Println("Error encoding results to file:", err)
		return
	}

	// 总结
	fmt.Printf("Total combinations: %d\n", len(results))
	fmt.Println("Top 10 most frequent combinations:")
	//最多次数前十
	for i := len(sortedResults) - 1; i >= len(sortedResults)-10 && i >= 0; i-- {
		fmt.Printf("%s: %d\n", sortedResults[i].Key.HandCard[0].CardTranslate()+sortedResults[i].Key.HandCard[1].CardTranslate(), sortedResults[i].Value)
	}
	//最少次数前十
	fmt.Println("Top 10 least frequent combinations:")
	for i := 0; i < 10 && i < len(sortedResults); i++ {
		fmt.Printf("%s: %d\n", sortedResults[i].Key.HandCard[0].CardTranslate()+sortedResults[i].Key.HandCard[1].CardTranslate(), sortedResults[i].Value)
	}

	//口袋对出现的次数
	fmt.Printf("Pair count: %d\n", pairCount)
	pairCountPercent := fmt.Sprintf("%.4f%%", float64(pairCount)/float64(iterations)*100)
	fmt.Println("Pair count:", pairCountPercent)
	//同花出现次数
	fmt.Printf("Suited  count: %d\n", suitedCount)
	suitedCountPercent := fmt.Sprintf("%.4f%%", float64(suitedCount)/float64(iterations)*100)
	fmt.Println("Suited count:", suitedCountPercent)
	//同花连张出现次数
	fmt.Printf("Suited connectors count: %d\n", suitedConnectorCount)
	suitedConnectorCountPercent := fmt.Sprintf("%.4f%%", float64(suitedConnectorCount)/float64(iterations)*100)
	fmt.Println("Suited connectors count:", suitedConnectorCountPercent)
	//AXs出现的次数
	fmt.Printf("Suited Ace count: %d\n", suitedAceCount)
	suitedAceCountPercent := fmt.Sprintf("%.4f%%", float64(suitedAceCount)/float64(iterations)*100)
	fmt.Println("Suited Ace count:", suitedAceCountPercent)
	//AXo出现的次数
	fmt.Printf("Ace high count: %d\n", aceCount)
	aceCountPercent := fmt.Sprintf("%.4f%%", float64(aceCount)/float64(iterations)*100)
	fmt.Println("Suited Ace count:", aceCountPercent)
	//两张牌都大于等于10出现的次数
	fmt.Printf("High cards count: %d\n", highCardCount)
	highCardCountPercent := fmt.Sprintf("%.4f%%", float64(highCardCount)/float64(iterations)*100)
	fmt.Println("High cards count:", highCardCountPercent)
}

// DealCards 发牌，返回玩家手牌和公共牌
func DealCards(New52CardList [52]Card, playersNumber int) (resHandCard []HandCard, resPublicCard []Card) {
	// 初始化玩家手牌
	resHandCard = make([]HandCard, playersNumber)
	resPublicCard = make([]Card, 52-(playersNumber*2))

	// 每个玩家分两张牌
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[0] = New52CardList[j-1]
	}
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[1] = New52CardList[playersNumber-1+j]
	}
	j := 0
	for i := 1; i <= 52-(playersNumber*2); i++ {
		resPublicCard[j] = New52CardList[2*playersNumber-1+i]
		j++
	}

	for i, k := range resHandCard {
		k.sortTwoCards()
		resHandCard[i] = k
	}
	return resHandCard, resPublicCard
}

// CombineCardsDemo 合成公共牌和所有玩家手牌的组合
func CombineCardsDemo(playersHandCard []HandCard, publicCard []Card) (resHandCard [][]Card) {
	n := len(playersHandCard)
	// //debug
	// fmt.Println("publicCard:", n)
	// //---
	// 初始化 res 切片，每个元素也是一个切片
	resHandCard = make([][]Card, n)
	for i := 0; i < n; i++ {
		// 添加玩家的手牌
		for j := 0; j < 2; j++ {
			resHandCard[i] = append(resHandCard[i], playersHandCard[i].HandCard[j])
		}
		// 添加公共牌 这里默认5张
		resHandCard[i] = append(resHandCard[i], publicCard[:5]...)
	}
	// 这7张再排序，从大到小，花色顺序按照 黑桃、红桃、梅花、方片的顺序
	for i := 0; i < n; i++ {
		for j := 0; j < 7; j++ {
			for k := j + 1; k < 7; k++ {
				if resHandCard[i][j].Rank < resHandCard[i][k].Rank {
					temp := resHandCard[i][j]
					resHandCard[i][j] = resHandCard[i][k]
					resHandCard[i][k] = temp
				}
			}
		}
	}
	return resHandCard
}

// Judge5From7 7选五的21种牌型的牌力，高牌的牌力为0，对子的牌力为1，两对的牌力为2，三条的牌力为3，顺子的牌力为4，同花的牌力为5，葫芦的牌力为6，四条的牌力为7，同花顺的牌力为8
func Judge5From7(playersAllCard [7]Card) (Grade int, MaxCard5 [5]Card) {
	//输入的7张牌，大小已经是按从大到小排列
	//定义四个花色的map，用来统计花色出现的次数
	suitMap := make(map[string]int)
	//记录最长等差数列的长度
	maxLen := 1
	currentLen := 1
	//记录最多大小相同的牌的长度
	sameMap := make(map[int]int)
	//要先判断是不是同花，至少有5张是相同花色的
	//先排除A2345的可能
	for i := 0; i < 7; i++ {
		//记录花色
		suitMap[playersAllCard[i].Suit] = suitMap[playersAllCard[i].Suit] + 1
		//记录大小相同的牌
		sameMap[playersAllCard[i].Rank] = sameMap[playersAllCard[i].Rank] + 1
	}
	//如果sameMap存在4张相同的牌，说明是四条
	for k, v := range sameMap {
		if v == 4 {
			//debug 四条统计
			fmt.Println("四条：", k, v)
			//---

			var MaxCard Card
			//max截取信号
			sign := false
			//maxcard的下标
			n := 0
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank != k && !sign { //不是4条的值就是最大值
					MaxCard = playersAllCard[i]
					sign = true
				}
				if playersAllCard[i].Rank == k {
					MaxCard5[n] = playersAllCard[i]
					n++
				}
			}
			//4条外的最后一张牌
			MaxCard5[4] = MaxCard
			return 7, MaxCard5
		}
	}
	//如果sameMap存在2种3张相同的牌，或者存在3张相同的牌的同时还至少存在一个对子，说明是葫芦
	count3 := 0
	for _, freq := range sameMap {
		if freq == 3 {
			count3++
			if count3 == 2 {
				return 6, MaxCard5
			}
		}

	}
	//判断是否有顺子，至少有5张是连续的
	for k, v := range suitMap {
		if v >= 5 {
			//debug 同花统计
			fmt.Println(k, v)
			//
			//判断是否是同花顺
			for i := 0; i < 7; i++ {
				if i >= 1 {
					if playersAllCard[i-1].Rank-playersAllCard[i].Rank == 1 {
						currentLen++
					} else {
						if currentLen > maxLen {
							maxLen = currentLen
						}
						currentLen = 1
					}
				}
			} //说明至少是个顺子
			if maxLen >= 5 {
				//debug 顺子统计
				fmt.Println(maxLen)
				//
				return 8, MaxCard5
			} else { //说明不是同花顺，不是顺子，返回同花
				return 5, MaxCard5
			}
			//debug 花色统计
			fmt.Println(suitMap)
			//
			return 0, MaxCard5
		}
	}
	return 0, MaxCard5
}

// sortTwoCards 对两张手牌进行排序
func (p *HandCard) sortTwoCards() {
	if p.HandCard[0].Rank < p.HandCard[1].Rank ||
		(p.HandCard[0].Rank == p.HandCard[1].Rank && compareSuits(p.HandCard[0].Suit, p.HandCard[1].Suit)) {
		// 交换两张牌
		p.HandCard[0], p.HandCard[1] = p.HandCard[1], p.HandCard[0]
	}
}

// max 返回较大的数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
