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

// ShuffleAndRecord 执行洗牌次，将所有组合以及频率放在一个map里并写入文件
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

func DealCards(New52CardList [52]Card, playersNumber int) (resHandCard []HandCard, resPublicCard []HandCard) {
	// 初始化玩家手牌
	resHandCard = make([]HandCard, playersNumber)

	// 每个玩家分两张牌
	for i := 0; i < playersNumber; i++ {
		resHandCard[i] = HandCard{
			HandCard: [2]Card{
				New52CardList[i*2],
				New52CardList[i*2+1],
			},
		}
	}
	return resHandCard, resHandCard
}

// Judge5From7 7选五的21种牌型的牌力
func Judge5From7(CardArry7_5 [7]Card, PlayName string) (CardRankSlice []CardRank) {
	var CardRanktest CardRank
	var ii, j, k, l, m int64
	var iii, jjj int64
	var MaxArr [5]Card
	var temp Card
	for ii = 0; ii <= 2; ii++ {
		for j = ii + 1; j <= 3; j++ {
			for k = j + 1; k <= 4; k++ {
				for l = k + 1; l <= 5; l++ {
					for m = l + 1; m <= 6; m++ {
						// fmt.Println(CardArry7_5[ii], CardArry7_5[j], CardArry7_5[k], CardArry7_5[l], CardArry7_5[m])
						MaxArr[0] = CardArry7_5[ii]
						MaxArr[1] = CardArry7_5[j]
						MaxArr[2] = CardArry7_5[k]
						MaxArr[3] = CardArry7_5[l]
						MaxArr[4] = CardArry7_5[m]
						for iii = 0; iii < 4; iii++ {
							for jjj = 0; jjj < 4; jjj++ {
								if MaxArr[jjj].Rank < MaxArr[jjj+1].Rank {
									temp = MaxArr[jjj+1]
									MaxArr[jjj+1] = MaxArr[jjj]
									MaxArr[jjj] = temp
								}
							}
						}
						// --debug//fmt.Println(MaxArr)
						//判断是否是同花
						if MaxArr[0].Suit == MaxArr[1].Suit &&
							MaxArr[0].Suit == MaxArr[2].Suit &&
							MaxArr[0].Suit == MaxArr[3].Suit &&
							MaxArr[0].Suit == MaxArr[4].Suit {
							//判断是否是同花顺
							if (MaxArr[0].Rank-MaxArr[1].Rank == 1 &&
								MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
								MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
								MaxArr[3].Rank-MaxArr[4].Rank == 1) ||
								(MaxArr[0].Rank-MaxArr[1].Rank == 9 &&
									MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
									MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
									MaxArr[3].Rank-MaxArr[4].Rank == 1) {
								//如果是同花顺
								CardRanktest.Grade = 9
								CardRanktest.Value = MaxArr
								CardRanktest.PlayName = PlayName
								CardRankSlice = append(CardRankSlice, CardRanktest)
								// --debug//fmt.Println("straight flush")
							} else {
								//如果是同花
								CardRanktest.Grade = 6
								CardRanktest.Value = MaxArr
								CardRanktest.PlayName = PlayName
								CardRankSlice = append(CardRankSlice, CardRanktest)
								// --debug//fmt.Println("flush")
							}
						} else {
							//判断是否是金刚
							if (MaxArr[0].Rank == MaxArr[3].Rank) ||
								(MaxArr[1].Rank == MaxArr[4].Rank) {
								//如果是金刚
								//判断重组金刚数组的顺序 保证前四个为相等的值，最后的值为单值
								if MaxArr[1].Rank == MaxArr[4].Rank {
									CardRanktest.Grade = 8
									CardRanktest.Value[4] = MaxArr[0]
									CardRanktest.Value[0] = MaxArr[1]
									CardRanktest.Value[1] = MaxArr[2]
									CardRanktest.Value[2] = MaxArr[3]
									CardRanktest.Value[3] = MaxArr[4]
									CardRanktest.PlayName = PlayName
									CardRankSlice = append(CardRankSlice, CardRanktest)
								} else {
									CardRanktest.Grade = 8
									CardRanktest.Value = MaxArr
									CardRanktest.PlayName = PlayName
									CardRankSlice = append(CardRankSlice, CardRanktest)
									// --debug//fmt.Println("KINGKONG")
								}
							} else {
								//判断是否是葫芦
								if (MaxArr[0].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank) ||
									(MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[4].Rank) {
									//如果是葫芦
									//判断三条数组的顺序 保证前三个值相等，后两个值为相等得对子
									if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[4].Rank {
										CardRanktest.Grade = 7
										CardRanktest.Value[0] = MaxArr[2]
										CardRanktest.Value[1] = MaxArr[3]
										CardRanktest.Value[2] = MaxArr[4]
										CardRanktest.Value[3] = MaxArr[0]
										CardRanktest.Value[4] = MaxArr[1]
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
									} else {
										CardRanktest.Grade = 7
										CardRanktest.Value = MaxArr
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
										// --debug//fmt.Println("full house")
									}
								} else {
									//判断是否是顺子
									if (MaxArr[0].Rank-MaxArr[1].Rank == 1 &&
										MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
										MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
										MaxArr[3].Rank-MaxArr[4].Rank == 1) ||
										(MaxArr[0].Rank-MaxArr[1].Rank == 9 &&
											MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
											MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
											MaxArr[3].Rank-MaxArr[4].Rank == 1) {
										//如果是顺子
										CardRanktest.Grade = 5
										CardRanktest.Value = MaxArr
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
										// --debug//fmt.Println("straight")
									} else {
										//判断是否是三条
										if (MaxArr[0].Rank == MaxArr[2].Rank) ||
											(MaxArr[2].Rank == MaxArr[4].Rank) ||
											(MaxArr[1].Rank == MaxArr[3].Rank) {
											//如果是三条
											//判断三条数组的顺序 保证前三个值相等，后两个值为大小顺序的单值
											if MaxArr[2].Rank == MaxArr[4].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value[0] = MaxArr[2]
												CardRanktest.Value[1] = MaxArr[3]
												CardRanktest.Value[2] = MaxArr[4]
												CardRanktest.Value[3] = MaxArr[0]
												CardRanktest.Value[4] = MaxArr[1]
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
											}
											if MaxArr[1].Rank == MaxArr[3].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value[0] = MaxArr[1]
												CardRanktest.Value[1] = MaxArr[2]
												CardRanktest.Value[2] = MaxArr[3]
												CardRanktest.Value[3] = MaxArr[0]
												CardRanktest.Value[4] = MaxArr[4]
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
											}
											if MaxArr[0].Rank == MaxArr[2].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value = MaxArr
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
												// --debug//fmt.Println("three of a kind")
											}
										} else {
											//判断是否是两对
											if (MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[3].Rank) ||
												(MaxArr[1].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank) ||
												(MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[3].Rank == MaxArr[4].Rank) {
												//如果是两对
												//判断两对数组的顺序 保证前四个值为从大到小的对子，最后一个值为单值
												if MaxArr[1].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value[0] = MaxArr[1]
													CardRanktest.Value[1] = MaxArr[2]
													CardRanktest.Value[2] = MaxArr[3]
													CardRanktest.Value[3] = MaxArr[4]
													CardRanktest.Value[4] = MaxArr[0]
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
												}
												if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[3].Rank == MaxArr[4].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value[0] = MaxArr[0]
													CardRanktest.Value[1] = MaxArr[1]
													CardRanktest.Value[2] = MaxArr[3]
													CardRanktest.Value[3] = MaxArr[4]
													CardRanktest.Value[4] = MaxArr[2]
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
												}
												if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[3].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value = MaxArr
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
													// --debug//fmt.Println("two pair")
												}
											} else {
												//判断是否是对子
												if MaxArr[0].Rank == MaxArr[1].Rank || MaxArr[1].Rank == MaxArr[2].Rank || MaxArr[2].Rank == MaxArr[3].Rank || MaxArr[3].Rank == MaxArr[4].Rank {
													//如果是对子
													//判断对子数组的顺序 保证前2个值为对子，最后三个值为从大到小单值
													if MaxArr[1].Rank == MaxArr[2].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[1]
														CardRanktest.Value[1] = MaxArr[2]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[3]
														CardRanktest.Value[4] = MaxArr[4]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[2].Rank == MaxArr[3].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[2]
														CardRanktest.Value[1] = MaxArr[3]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[1]
														CardRanktest.Value[4] = MaxArr[4]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[3].Rank == MaxArr[4].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[3]
														CardRanktest.Value[1] = MaxArr[4]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[1]
														CardRanktest.Value[4] = MaxArr[2]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[0].Rank == MaxArr[1].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value = MaxArr
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
														// --debug//fmt.Println("one pair")
													}
												} else {
													//为高张
													CardRanktest.Grade = 1
													CardRanktest.Value = MaxArr
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
													// --debug//fmt.Println("high card")
												}
											}
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
	return CardRankSlice
}
