package cardFunc

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
