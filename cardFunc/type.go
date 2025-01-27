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
	Advance  Advance `json:"advance,omitempty"`
	PlayerID string  `json:"playerid,omitempty"`
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
	ChipSum      int64    `json:"chipsum,omitempty"`
	ChipBackHand int64    `json:"chipbackhand,omitempty"`
	BankRollSum  int64    `json:"bankrollsum,omitempty"`
	Card7        [7]Card  `json:"card7,omitempty"`
	Card5        [5]Card  `json:"card5,omitempty"`
	CardAll      []Card   `json:"cardAll,omitempty"` //目前默认长度是7
	Grade        int      `json:"grade"`
	TableNum     int      `json:"tablenum,omitempty"`
	Sitnum       int      `json:"sitnum,omitempty"`
	IsActive     bool     `json:"isactive,omitempty"`
	IsFold       bool     `json:"isfold,omitempty"`
	IsAllIn      bool     `json:"isallin,omitempty"`
	WinCount     int      `json:"winCount"`           //单人训练统计次数用的
	WinRate      float64  `json:"winRate"`            //单人训练统计胜率用的
	Vpip         float64  `json:"vpip,omitempty"`     //Voluntarily Put In Pot主动入池率
	PFR          float64  `json:"pfr,omitempty"`      //Pre-Flop Raise 翻牌前加注概率
	FRR          float64  `json:"frr,omitempty"`      //Flop Raise 翻牌后加注概率
	TRR          float64  `json:"trr,omitempty"`      //Turn Raise 转牌后加注概率
	RRR          float64  `json:"rrr,omitempty"`      //River Raise 河牌后加注概率
	ReRaise      float64  `json:"rerraise,omitempty"` //Re-Raise 加注后加注概率
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
	PlayerNumber int        `json:"playernumber,omitempty"`
	HandCardList []HandCard `json:"handcardlist,omitempty"`
	RoundNumber  int        `json:"roundnumber,omitempty"`
	DebugSwitch  bool       `json:"debugswitch,omitempty"`
}

// type PracticeList struct {
// 	AllHandList map[HandCard]int `json:"allhandlist"`
// 	WinnerList  Players          `json:"winerlist"`
// }

type HandConfigDemo02 struct {
	PlayerNumber int        `json:"playernumber,omitempty"`
	HandCardList []HandCard `json:"handcardlist,omitempty"`
	PublicCard   PublicCard `json:"publiccard,omitempty"`
	RoundNumber  int        `json:"roundnumber,omitempty"`
	DebugSwitch  bool       `json:"debugswitch,omitempty"`
}

// Advance 高级配置，目前只有范围  单机使用
type Advance struct {
	Range float64 `json:"range,omitempty"`
}
type PublicCard struct {
	Flop  [3]Card `json:"flop,omitempty"`
	Turn  Card    `json:"turn,omitempty"`
	River Card    `json:"river,omitempty"`
}

type PracticeResDemo02 struct {
	PlayerWinCount []PlayersRes   `json:"playerwincount"`       //统计玩家的获胜次数  按座次排序
	WinGradeList   []WinGradeList `json:"wingradelist"`         //获胜的成牌牌力分布统计  按出现次数排序
	DrawCount      int            `json:"drawcount"`            //平局次数
	So169ComboList []So169Combo   `json:"so169combo,omitempty"` //169组so组合的胜率统计  按胜率排序
	RoundNumber    int            `json:"roundnumber,omitempty"`
}

type PlayersRes struct {
	PlayerID string  `json:"playerid"`
	WinCount int     `json:"wincount"`
	WinRate  float64 `json:"winrate"`
}

type So169Combo struct {
	WinRateRank int     `json:"winraterank"`
	So169       string  `json:"so169"`
	WinRate     float64 `json:"winrate"`
	ExistCount  int     `json:"existcount"`
	WinCount    int     `json:"wincount"`
}

type WinGradeList struct {
	Grade     int    `json:"grade"`
	GradeName string `json:"gradename"`
	WinCount  int    `json:"wincount"`
}
