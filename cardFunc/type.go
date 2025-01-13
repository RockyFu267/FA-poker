package cardFunc

// Card 牌的结构体
type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

// CardRank 翻拍后玩家的最大牌行同级数组
type CardRank struct {
	Grade    int     `json:"grade"`
	Value    [5]Card `json:"value"`
	PlayName string  `json:"playname"`
	Value7   [7]Card `json:"value7"`
}

// HanCard 手牌
type HandCard struct {
	HandCard [2]Card `json:"handCard"`
}

// TableSingle 牌桌
type TableSingle struct {
	ID               string         `json:"id"`
	SitMax           int            `json:"sitmax"`
	SitOn            int            `json:"siton"`
	BigBlindSitNum   int            `json:"bigblindsitnum"`
	SmallBlindSitNum int            `json:"smallblindsitnum"`
	BigBlind         int            `json:"bigblind"`
	SmallBlind       int            `json:"smallblind"`
	DealingOrder     []int          `json:"dealingorder"`
	Players          []Players      `json:"players"`
	RoundHistory     []RoundHistory `json:"roundhistory"`
	ChipSumMaxLimt   int            `json:"chipsummaxlimt"`
	ChipSumMinLimt   int            `json:"chipsumminlimt"`
}

// Players 玩家
type Players struct {
	ID           string   `json:"id"`
	Hand         HandCard `json:"hand"`
	ChipSum      int64    `json:"chipsum"`
	ChipBackHand int64    `json:"chipbackhand"`
	BankRollSum  int64    `json:"bankrollsum"`
	Card7        [7]Card  `json:"card7"`
	Card5        [5]Card  `json:"card5"`
	CardAll      []Card   `json:"cardAll"` //目前默认长度是7
	Grade        int      `json:"grade"`
	TableNum     int      `json:"tablenum"`
	Sitnum       int      `json:"sitnum"`
	IsActive     bool     `json:"isactive"`
	IsFold       bool     `json:"isfold"`
	IsAllIn      bool     `json:"isallin"`
}

// RoundHistory	牌局历史
type RoundHistory struct {
	ID         string `json:"id"`
	TableID    string `json:"tableid"`
	Round      int    `json:"round"`
	BigBlind   int    `json:"bigblind"`
	SmallBlind int    `json:"smallblind"`
	Dealer     int    `json:"dealer"`
	Players    []Players
}

type HandConfig struct {
	PlayerNumber int        `json:"playernumber"`
	HandCardList []HandCard `json:"handcardlist"`
	RoundNumber  int        `json:"roundnumber"`
}
