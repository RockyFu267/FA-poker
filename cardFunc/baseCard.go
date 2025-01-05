package cardFunc

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// 指定人数 洗牌，发牌，并比较谁的牌最大,并且可以选择指定手牌
func shuffleJudgeDemo(playlist []Players, appointHandCardList []HandCard) (winner []Players) {
	playerNum := len(playlist)
	var Card52 []Card
	var resHandList []HandCard
	var pubHandList []Card

	if appointHandCardList != nil { //填写指定发牌的逻辑
	} else {
		Card52 = ShuffleCard() //洗牌
		resHandList, pubHandList = DealCards(Card52, playerNum)
	}
	fmt.Println("手牌组：", len(resHandList), resHandList)        //debug
	fmt.Println("公共手牌长度以及手牌 ", len(pubHandList), pubHandList) //debug

	maxGrade := 0
	maxCard5 := [5]int{0, 0, 0, 0, 0}
	// 假装这里已经处理了座次
	for i := 0; i < len(resHandList); i++ {
		var tempCard7 [7]Card
		playlist[i].Hand = resHandList[i]
		tempCard7[0] = playlist[i].Hand.HandCard[0]
		tempCard7[1] = playlist[i].Hand.HandCard[1]
		tempCard7[2] = pubHandList[0]
		tempCard7[3] = pubHandList[1]
		tempCard7[4] = pubHandList[2]
		tempCard7[5] = pubHandList[3]
		tempCard7[6] = pubHandList[4]
		tempCard7 = sortCards(tempCard7)
		playlist[i].Card7 = tempCard7
		playlist[i].Grade, playlist[i].Card5 = Judge5From7(playlist[i].Card7)
		fmt.Println(playlist[i].ID, playlist[i].Hand)              //debug
		fmt.Println(playlist[i].Grade, "-max-", playlist[i].Card5) //debug
		if maxGrade == playlist[i].Grade {
			for j := 0; j < 5; j++ {
				if playlist[i].Card5[j].Rank > maxCard5[j] {
					maxCard5[0] = playlist[i].Card5[0].Rank
					maxCard5[1] = playlist[i].Card5[1].Rank
					maxCard5[2] = playlist[i].Card5[2].Rank
					maxCard5[3] = playlist[i].Card5[3].Rank
					maxCard5[4] = playlist[i].Card5[4].Rank
					break
				}
				if playlist[i].Card5[j].Rank == maxCard5[j] {
					continue
				}
				if playlist[i].Card5[j].Rank < maxCard5[j] {
					break
				}
			}
			continue
		}
		if maxGrade < playlist[i].Grade {
			maxGrade = playlist[i].Grade
			for j := 0; j < 5; j++ {
				maxCard5[j] = playlist[i].Card5[j].Rank
			}
			continue
		}
	}
	fmt.Println("len ", len(winner))   //debug
	fmt.Println("maxGrade ", maxGrade) //debug
	fmt.Println("maxCard5 ", maxCard5) //debug

	for i := 0; i < len(playlist); i++ {
		if playlist[i].Grade == maxGrade {
			fmt.Println("最大的ID ", playlist[i].ID) //debug
			sign := true
			for j := 0; j < 5; j++ {
				if playlist[i].Card5[j].Rank == maxCard5[j] {
					continue
				} else {
					sign = false
				}
			}
			if sign {
				winner = append(winner, playlist[i])
			}
		}
	}
	fmt.Println("len2 ", len(winner)) //debug

	return winner
}

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
func ShuffleCard() (New52CardList []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化52张牌
	var Card52 = []Card{
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

// ShuffleCard 少牌洗牌
func shortOfShuffleCard(input []Card) (New52CardList []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化52张牌
	var Card52 = []Card{
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
	var result []Card
	existMap := make(map[Card]bool)
	for _, v := range input {
		key := v
		existMap[key] = true
	}
	for _, v := range Card52 {
		key := v
		if !existMap[key] {
			result = append(result, v)
		}
	}
	Card52 = result

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
func GetTopTwoCards(deck []Card) HandCard {
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
func DealCards(New52CardList []Card, playersNumber int) (resHandCard []HandCard, resPublicCard []Card) {
	// 初始化玩家手牌
	resHandCard = make([]HandCard, playersNumber)
	resPublicCard = make([]Card, len(New52CardList)-(playersNumber*2))

	// 每个玩家分两张牌
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[0] = New52CardList[j-1]
	}
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[1] = New52CardList[playersNumber-1+j]
	}
	j := 0
	for i := 1; i <= len(New52CardList)-(playersNumber*2); i++ {
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

// udge5From7 7选五的21种牌型的牌力，高牌的牌力为0，对子的牌力为1，两对的牌力为2，三条的牌力为3，顺子的牌力为4，同花的牌力为5，葫芦的牌力为6，四条的牌力为7，同花顺的牌力为8
func Judge5From7(playersAllCard [7]Card) (Grade int, MaxCard5 [5]Card) { // Judge5From7 7选五的21种牌型的牌力，高牌的牌力为0，对子的牌力为1，两对的牌力为2，三条的牌力为3，顺子的牌力为4，同花的牌力为5，葫芦的牌力为6，四条的牌力为7，同花顺的牌力为8
	//输入的7张牌，大小已经是按从大到小排列
	suitMap := make(map[string]int)       //定义四个花色的map，用来统计花色出现的次数
	suitListMap := make(map[string][]int) //定义四个花色的map，用来统计相同花色的rank
	sameMap := make(map[int]int)          //记录最多大小相同的牌的长度
	for i := 0; i < 7; i++ {
		suitMap[playersAllCard[i].Suit] = suitMap[playersAllCard[i].Suit] + 1                                     //记录花色
		sameMap[playersAllCard[i].Rank] = sameMap[playersAllCard[i].Rank] + 1                                     //记录大小相同的牌
		suitListMap[playersAllCard[i].Suit] = append(suitListMap[playersAllCard[i].Suit], playersAllCard[i].Rank) //记录相同花色的牌
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
			acesign := false
			fiveSign := false
			fourSign := false
			threeSign := false
			twoSign := false
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == 14 && !acesign {
					MaxCard5[4] = playersAllCard[i]
					acesign = true
					continue
				}
				if playersAllCard[i].Rank == 5 && !fiveSign {
					MaxCard5[0] = playersAllCard[i]
					fiveSign = true
					continue
				}
				if playersAllCard[i].Rank == 4 && !fourSign {
					MaxCard5[1] = playersAllCard[i]
					fourSign = true
					continue
				}
				if playersAllCard[i].Rank == 3 && !threeSign {
					MaxCard5[2] = playersAllCard[i]
					threeSign = true
					continue
				}
				if playersAllCard[i].Rank == 2 && !twoSign {
					MaxCard5[3] = playersAllCard[i]
					twoSign = true
					continue
				}
			}
			return Grade, MaxCard5
		} else { //只能是两对或者三条
			pariRank1 := 0
			pariRank2 := 0
			cont3value := 0
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
					cont3value = k
					break
				}
			}
			if cont3value != 0 { //只能是三条
				Grade = 3
				max1 := 0
				max2 := 0
				for i := 0; i < len(playersAllCard); i++ {
					if playersAllCard[i].Rank == cont3value {
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						i = i + 2
						continue
					}
					if max1 == 0 {
						max1 = playersAllCard[i].Rank
						MaxCard5[3] = playersAllCard[i]
						continue
					}
					if max2 == 0 {
						max2 = playersAllCard[i].Rank
						MaxCard5[4] = playersAllCard[i]
						continue
					}
				}

				return Grade, MaxCard5

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
		straighACEtoFive := false
		straighACEtoFive = containsStraightKeys(sameMap)
		pairRank1 := 0
		var templist []int
		var tempList []int
		for k, v := range sameMap {
			templist = append(templist, k)
			if v == 2 {
				pairRank1 = k
			}
		}
		sortDescending(templist)
		tempList = templist
		if templist[0]-templist[4] == 4 {
			tempList = templist[:5]
		}
		if templist[1]-templist[5] == 4 && templist[0]-templist[4] != 4 {
			tempList = templist[1:]
		}

		for k, v := range suitListMap { //判断是否有同花，可能是同花、同花顺
			if len(v) == 6 { //有同花
				if suitListMap[k][0]-suitListMap[k][4] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][1]-suitListMap[k][5] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][5] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][0] == 14 && suitListMap[k][2] == 5 && suitListMap[k][3] == 4 && suitListMap[k][4] == 3 && suitListMap[k][5] == 2 { //同花顺 指定牌型 5432A的牌型

					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[4] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][5] {
							MaxCard5[3] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				} else { //只是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
			if len(v) == 5 { //
				if suitListMap[k][0]-suitListMap[k][4] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][0] == 14 && suitListMap[k][1] == 5 && suitListMap[k][2] == 4 && suitListMap[k][3] == 3 && suitListMap[k][4] == 2 { //同花顺 指定牌型 5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[4] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[3] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				} else { //只是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}

		}
		if len(tempList) == 5 { //没有同花的可能，只能是顺子
			Grade = 4
			j := 0 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if tempList[j] == playersAllCard[i].Rank { //长度可能超过5 所以还得加上判断
					MaxCard5[j] = playersAllCard[i]
					if j >= 4 {
						break
					}
					j++
					continue
				}
			}
			return Grade, MaxCard5
		}
		if straighACEtoFive { //ace顺子 5432A
			Grade = 4
			acesign := false
			fiveSign := false
			fourSign := false
			threeSign := false
			twoSign := false
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == 14 && !acesign {
					MaxCard5[4] = playersAllCard[i]
					acesign = true
					continue
				}
				if playersAllCard[i].Rank == 5 && !fiveSign {
					MaxCard5[0] = playersAllCard[i]
					fiveSign = true
					continue
				}
				if playersAllCard[i].Rank == 4 && !fourSign {
					MaxCard5[1] = playersAllCard[i]
					fourSign = true
					continue
				}
				if playersAllCard[i].Rank == 3 && !threeSign {
					MaxCard5[2] = playersAllCard[i]
					threeSign = true
					continue
				}
				if playersAllCard[i].Rank == 2 && !twoSign {
					MaxCard5[3] = playersAllCard[i]
					twoSign = true
					continue
				}
			}
			return Grade, MaxCard5
		} else { //只能是对子
			Grade = 1
			j := 2 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == pairRank1 { //赋值对子
					MaxCard5[0] = playersAllCard[i]
					MaxCard5[1] = playersAllCard[i+1]
					i = i + 1
					continue
				} else { //赋值后三位
					if j < 5 {
						MaxCard5[j] = playersAllCard[i]
						j++
					}
				}
			}
			return Grade, MaxCard5
		}
	case 7: //可能是同花顺、同花、顺子、高牌
		for k, v := range suitMap { //判断是否是同花顺、顺子
			if v >= 5 { //至少是同花
				if playersAllCard[0].Rank-playersAllCard[4].Rank == 4 && playersAllCard[0].Suit == k && playersAllCard[1].Suit == k && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[0]
					MaxCard5[1] = playersAllCard[1]
					MaxCard5[2] = playersAllCard[2]
					MaxCard5[3] = playersAllCard[3]
					MaxCard5[4] = playersAllCard[4]
					return Grade, MaxCard5
				}
				if playersAllCard[1].Rank-playersAllCard[5].Rank == 4 && playersAllCard[1].Suit == k && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[1]
					MaxCard5[1] = playersAllCard[2]
					MaxCard5[2] = playersAllCard[3]
					MaxCard5[3] = playersAllCard[4]
					MaxCard5[4] = playersAllCard[5]
					return Grade, MaxCard5
				}
				if playersAllCard[2].Rank-playersAllCard[6].Rank == 4 && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k && playersAllCard[6].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[2]
					MaxCard5[1] = playersAllCard[3]
					MaxCard5[2] = playersAllCard[4]
					MaxCard5[3] = playersAllCard[5]
					MaxCard5[4] = playersAllCard[6]
					return Grade, MaxCard5
				}
				if playersAllCard[0].Rank == 14 && playersAllCard[3].Rank == 5 && playersAllCard[4].Rank == 4 && playersAllCard[5].Rank == 3 && playersAllCard[6].Rank == 2 && playersAllCard[0].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k && playersAllCard[6].Suit == k { //同花顺 5432A
					Grade = 8
					MaxCard5[4] = playersAllCard[0]
					MaxCard5[0] = playersAllCard[3]
					MaxCard5[1] = playersAllCard[4]
					MaxCard5[2] = playersAllCard[5]
					MaxCard5[3] = playersAllCard[6]
					return Grade, MaxCard5
				} else { // 只能是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
		}
		if playersAllCard[0].Rank-playersAllCard[4].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
			return Grade, MaxCard5
		}
		if playersAllCard[1].Rank-playersAllCard[5].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[1]
			MaxCard5[1] = playersAllCard[2]
			MaxCard5[2] = playersAllCard[3]
			MaxCard5[3] = playersAllCard[4]
			MaxCard5[4] = playersAllCard[5]
			return Grade, MaxCard5
		}
		if playersAllCard[2].Rank-playersAllCard[6].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[2]
			MaxCard5[1] = playersAllCard[3]
			MaxCard5[2] = playersAllCard[4]
			MaxCard5[3] = playersAllCard[5]
			MaxCard5[4] = playersAllCard[6]
			return Grade, MaxCard5
		}
		if playersAllCard[0].Rank == 14 && playersAllCard[3].Rank == 5 && playersAllCard[4].Rank == 4 && playersAllCard[5].Rank == 3 && playersAllCard[6].Rank == 2 { //顺子 5432A
			Grade = 4
			MaxCard5[4] = playersAllCard[0]
			MaxCard5[0] = playersAllCard[3]
			MaxCard5[1] = playersAllCard[4]
			MaxCard5[2] = playersAllCard[5]
			MaxCard5[3] = playersAllCard[6]
			return Grade, MaxCard5
		} else { //只能是高牌
			Grade = 0
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
			return Grade, MaxCard5
		}
	default:
		Grade = 0
		return Grade, MaxCard5
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

// sortDescending 对切片进行降序排序
func sortDescending(arr []int) {
	// 使用 sort.Slice 进行降序排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j] // 比较函数，定义降序排序
	})
}

// handSorting 手牌排序
func handSorting(resHand []Card) []Card { //手牌排序
	n := len(resHand)
	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if resHand[j].Rank < resHand[k].Rank {
				temp := resHand[j]
				resHand[j] = resHand[k]
				resHand[k] = temp
			}
		}
	}
	return resHand
}

// sortCards对[7]Card数组进行排序
func sortCards(cards [7]Card) [7]Card {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			// 先比较牌面数字大小
			if cards[j].Rank < cards[j+1].Rank {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			} else if cards[j].Rank == cards[j+1].Rank {
				// 牌面数字相同，比较花色权重
				if suitWeight[cards[j].Suit] < suitWeight[cards[j+1].Suit] {
					cards[j], cards[j+1] = cards[j+1], cards[j]
				}
			}
		}
	}
	return cards
}

// 定义花色对应的权重，用于比较相同牌面数字时的大小关系
var suitWeight = map[string]int{
	"黑桃": 4,
	"红桃": 3,
	"方片": 2,
	"梅花": 1,
}
