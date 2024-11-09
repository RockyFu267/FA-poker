package cardFunc

// Card 牌的结构体
type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

// CardRank 翻拍后玩家的最大牌行同级数组
type CardRank struct {
	Grade    int64   `json:"grade"`
	Value    [5]Card `json:"value"`
	PlayName string  `json:"playName"`
	Value7   [7]Card `json:"value7"`
}

// HanCard 手牌
type HandCard struct {
	HandCard [2]Card `json:"handCard"`
}

//
