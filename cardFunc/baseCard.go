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
	suitMap := make(map[string]int) //定义四个花色的map，用来统计花色出现的次数
	sameMap := make(map[int]int)    //记录最多大小相同的牌的长度
	for i := 0; i < 7; i++ {
		suitMap[playersAllCard[i].Suit] = suitMap[playersAllCard[i].Suit] + 1 //记录花色
		sameMap[playersAllCard[i].Rank] = sameMap[playersAllCard[i].Rank] + 1 //记录大小相同的牌
	}

	//根据map长度来判断大小
	switch len(sameMap) { //这种写法不适合后期做娱乐技能判定，标准先这样
	case 2: //只有可能是金刚
		Grade = 7
		Value4 := 0 //不需要1
		for k, v := range sameMap {
			if v == 4 {
				Value4 = k
			}
		}
		if playersAllCard[0].Rank == Value4 { //不是43就是34
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
		} else {
			MaxCard5[0] = playersAllCard[3]
			MaxCard5[1] = playersAllCard[4]
			MaxCard5[2] = playersAllCard[5]
			MaxCard5[3] = playersAllCard[6]
			MaxCard5[4] = playersAllCard[0]
		}
		return Grade, MaxCard5
	case 3: //可能是金刚也可能是葫芦  这种写法不适合后期做娱乐技能判定，标准先这样
		count3 := 0 //有3个相同的牌出现，那一定不是金刚了，值大于2就一定是葫芦（3+3+1）
		count2 := 0 //如果对子出现两次，那一定是葫芦（3+2+2）
		for _, v := range sameMap {
			if v == 3 {
				count3 = count3 + 1
				if count3 == 2 { //只有以下组合（3+3+1）
					Grade = 6
					if playersAllCard[0].Rank == playersAllCard[2].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank { //3+3+1
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[0].Rank == playersAllCard[2].Rank && playersAllCard[4].Rank == playersAllCard[6].Rank { //3+1+3
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[5]
						return Grade, MaxCard5
					} else { //1+3+3
						MaxCard5[0] = playersAllCard[1]
						MaxCard5[1] = playersAllCard[2]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[5]
						return Grade, MaxCard5
					}
				}
			}
			if v == 2 {
				count2 = count2 + 1
				if count2 == 2 { //只有以下组合（3+2+2）
					Grade = 6
					if playersAllCard[0].Rank == playersAllCard[2].Rank { //322
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[2].Rank == playersAllCard[4].Rank { //232
						MaxCard5[0] = playersAllCard[2]
						MaxCard5[1] = playersAllCard[3]
						MaxCard5[2] = playersAllCard[4]
						MaxCard5[3] = playersAllCard[0]
						MaxCard5[4] = playersAllCard[1]
						return Grade, MaxCard5
					} else { //223  playersAllCard[4].Rank == playersAllCard[6].Rank
						MaxCard5[0] = playersAllCard[4]
						MaxCard5[1] = playersAllCard[5]
						MaxCard5[2] = playersAllCard[6]
						MaxCard5[3] = playersAllCard[0]
						MaxCard5[4] = playersAllCard[1]
						return Grade, MaxCard5
					}
				}

			}
			if v == 4 { //如果有4个相同的牌，那一定是金刚 （4+2+1）
				Grade = 7
				if playersAllCard[0].Rank == playersAllCard[3].Rank { //4+ 2+1或1+2
					MaxCard5[0] = playersAllCard[0]
					MaxCard5[1] = playersAllCard[1]
					MaxCard5[2] = playersAllCard[2]
					MaxCard5[3] = playersAllCard[3]
					MaxCard5[4] = playersAllCard[4]
					return Grade, MaxCard5
				}
				if playersAllCard[1].Rank == playersAllCard[4].Rank { //1+4+2
					MaxCard5[0] = playersAllCard[1]
					MaxCard5[1] = playersAllCard[2]
					MaxCard5[2] = playersAllCard[3]
					MaxCard5[3] = playersAllCard[4]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				}
				if playersAllCard[2].Rank == playersAllCard[5].Rank { //2+4+1
					MaxCard5[0] = playersAllCard[2]
					MaxCard5[1] = playersAllCard[3]
					MaxCard5[2] = playersAllCard[4]
					MaxCard5[3] = playersAllCard[5]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				} else { //1+2或2+1 +4
					MaxCard5[0] = playersAllCard[3]
					MaxCard5[1] = playersAllCard[4]
					MaxCard5[2] = playersAllCard[5]
					MaxCard5[3] = playersAllCard[6]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				}

			}

		}
	case 4: //可能是金刚也可能是葫芦 也可能是两对  这种写法不适合后期做娱乐技能判定，标准先这样
		count2 := 0      //如果对子出现3次，那一定是2对 2+2+2+1
		count2Value := 0 //用于记录葫芦组合的对子的值
		for k, v := range sameMap {
			if v == 4 { // 如果有4个相同的牌，那一定是金刚 （4+1+1+1）
				Grade = 7
				count4Value := 0 //用于记录4个相同的牌  为了提前结束循环
				maxValue := 0    //用于记录最大的牌 	为了提前结束循环
				for i := 0; i < 7; i++ {
					if playersAllCard[i].Rank != k && maxValue == 0 { //不是4条的值就是最大值
						MaxCard5[4] = playersAllCard[i]
						maxValue = playersAllCard[i].Rank
						if count4Value != 0 { //所有值都已确认
							return Grade, MaxCard5
						}
						continue

					}
					if playersAllCard[i].Rank == k { //找到确认金刚的值
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						MaxCard5[3] = playersAllCard[i+3]
						i = i + 3
						count4Value = k
						if maxValue != 0 { //所有值都已确认
							return Grade, MaxCard5
						} else {
							continue
						}
					}
				}
				return Grade, MaxCard5
			}
			if v == 2 {
				count2Value = k //用于记录葫芦中的唯一对子

				count2 = count2 + 1
				if count2 == 3 { //只有以下组合（2+2+2+1）
					Grade = 2
					if playersAllCard[0].Rank == playersAllCard[1].Rank && playersAllCard[2].Rank == playersAllCard[3].Rank { //2+2+2+1
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[0].Rank == playersAllCard[1].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank { //2+1+2+2
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[2]
						return Grade, MaxCard5
					} else { //1+2+2+2  playersAllCard[1].Rank == playersAllCard[2].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank
						MaxCard5[0] = playersAllCard[1]
						MaxCard5[1] = playersAllCard[2]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[0]
						return Grade, MaxCard5
					}
				}
			}
			if v == 3 { //只有可能是 3+2+ 1+ 1的葫芦组合
				Grade = 6
				for i := 0; i < len(playersAllCard); i++ {
					if playersAllCard[i].Rank == k { //葫芦中的3已确认
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						break
					}
				}
			}
		}
		for i := 0; i < len(playersAllCard); i++ { //葫芦中三条已在上面的循环中赋值了，还差葫芦中的对子
			if playersAllCard[i].Rank == count2Value {
				MaxCard5[3] = playersAllCard[i]
				MaxCard5[4] = playersAllCard[i+1]
				break
			}
		}
		return Grade, MaxCard5
	case 5: //可能是同花顺、同花、顺子、三条、两对
		straighACEtoFive := false
		straighACEtoFive = containsStraightKeys(sameMap)
		for k, v := range suitMap { //判断是否有同花，可能是同花、同花顺
			if v == 5 { //有同花
				if playersAllCard[0].Rank-playersAllCard[6].Rank == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
				if straighACEtoFive { //同花顺 指定牌型 5432A的牌型
					Grade = 8
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i] //顺序还要调整 现在是A5432
							j++
						}
					}
					MaxCard5[0], MaxCard5[1], MaxCard5[2], MaxCard5[3], MaxCard5[4] = MaxCard5[1], MaxCard5[2], MaxCard5[3], MaxCard5[4], MaxCard5[0] //调整顺序
					return Grade, MaxCard5
				} else { //没有同花顺的可能，就是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
		}
		if playersAllCard[0].Rank-playersAllCard[6].Rank == 4 { //没有同花的可能，可能是顺子，不包含5432A的牌型
			Grade = 4
			j := 0 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if i == 0 { //第一个直接赋值 仅适用same长度为5的牌型
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				}
				if MaxCard5[j-1].Rank-playersAllCard[i].Rank == 1 {
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				}
			}
			return Grade, MaxCard5
		}
		if straighACEtoFive { //ace顺子 5432A
			Grade = 4
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == 14 {
					MaxCard5[4] = playersAllCard[i]
					continue
				}
				if playersAllCard[i].Rank == 5 {
					MaxCard5[0] = playersAllCard[i]
					continue
				}
				if playersAllCard[i].Rank == 4 {
					MaxCard5[1] = playersAllCard[i]
					continue
				}
				if playersAllCard[i].Rank == 3 {
					MaxCard5[2] = playersAllCard[i]
					continue
				}
				if playersAllCard[i].Rank == 2 {
					MaxCard5[3] = playersAllCard[i]
					continue
				}
			}
			return Grade, MaxCard5
		} else { //只能是两对或者三条
			pariRank1 := 0
			pariRank2 := 0
			for k, v := range sameMap {
				if v == 2 { //只可能是两对
					Grade = 2
					if pariRank1 == 0 {
						pariRank1 = k
						continue
					}
					pariRank2 = k
					break
				}
				if v == 3 { //只可能是三条
					Grade = 6
					break
				}
			}
			j := 0 //maxCard5的下标
			maxCardSign := false
			for i := 0; i < len(playersAllCard); i++ {
				if playersAllCard[i].Rank == pariRank1 || playersAllCard[i].Rank == pariRank2 {
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				} else {
					if !maxCardSign {
						MaxCard5[4] = playersAllCard[i]
						maxCardSign = true
					}
					continue
				}
			}
			return Grade, MaxCard5
		}
	case 6: //可能是同花顺、同花、顺子、两对
	case 7:
	default:
	}

	return 0, MaxCard5
}

// containsStraightKeys 判断map中是否同时包含14, 2, 3, 4, 5的key
func containsStraightKeys(cards map[int]int) bool {
	requiredKeys := []int{14, 2, 3, 4, 5}

	for _, key := range requiredKeys {
		if _, exists := cards[key]; !exists {
			return false
		}
	}
	return true
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
