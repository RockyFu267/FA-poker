package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

var suits = []string{"黑桃", "红桃", "梅花", "方片"}                  // 四种花色
var ranks = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14} // 点数2-14，J(11), Q(12), K(13), A(14)

func createDeck() []string {
	var deck []string
	// 创建52张牌
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, fmt.Sprintf("%s%d", suit, rank))
		}
	}
	return deck
}

// 检查牌组中是否包含某张牌
func contains(cards []string, card string) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}

// 洗牌并抽取两张牌
func getRandomHand(deck []string) string {
	// 打乱牌组
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// 取前两张作为对手手牌
	card1 := deck[0]
	card2 := deck[1]

	// 保证顺序：大牌排前，小牌排后
	return sortHand(card1 + " " + card2)
}

// 对两张牌进行排序，大小牌排前，花色优先级排序
func sortHand(hand string) string {
	cards := strings.Fields(hand)

	// 获取牌的花色和点数
	card1Suit := cards[0][:len(cards[0])-1] // 花色
	card1Rank := cards[0][len(cards[0])-1:] // 点数
	card2Suit := cards[1][:len(cards[1])-1] // 花色
	card2Rank := cards[1][len(cards[1])-1:] // 点数

	// 将牌面值转为数字
	card1Value := stringToRank(card1Rank)
	card2Value := stringToRank(card2Rank)

	// 如果两张牌大小不相同，按大小排序
	if card1Value > card2Value {
		return cards[0] + " " + cards[1]
	} else if card1Value < card2Value {
		return cards[1] + " " + cards[0]
	}

	// 如果大小相同，按花色优先级排序
	if suitPriority(card1Suit) < suitPriority(card2Suit) {
		return cards[0] + " " + cards[1]
	} else {
		return cards[1] + " " + cards[0]
	}
}

// 将字符串转换为点数
func stringToRank(rank string) int {
	switch rank {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		return atoi(rank)
	}
}

// 字符串转数字
func atoi(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		result = result*10 + int(s[i]-'0')
	}
	return result
}

// 花色优先级：黑桃 > 红桃 > 梅花 > 方片
func suitPriority(suit string) int {
	switch suit {
	case "黑桃":
		return 1
	case "红桃":
		return 2
	case "梅花":
		return 3
	case "方片":
		return 4
	default:
		return 5
	}
}

// 打印结果到文件
func printResultsToFile(handStats map[string]int) {
	// 打开文件
	file, err := os.Create("hand_statistics.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 将map转换为slice，并按频率排序
	var sortedStats []struct {
		Hand  string
		Count int
	}

	for hand, count := range handStats {
		sortedStats = append(sortedStats, struct {
			Hand  string
			Count int
		}{Hand: hand, Count: count})
	}

	// 按照频率排序，升序排列
	sort.Slice(sortedStats, func(i, j int) bool {
		return sortedStats[i].Count < sortedStats[j].Count
	})

	// 打印结果到文件
	for _, stat := range sortedStats {
		file.WriteString(fmt.Sprintf("%s: %d 次\n", stat.Hand, stat.Count))
	}

	// 输出总结信息
	file.WriteString("\n总结：\n")
	file.WriteString(fmt.Sprintf("对手手牌一共有 %d 种组合\n", len(handStats)))

	// 输出最高出现次数前5的手牌组合
	file.WriteString("\n最高出现次数前5的手牌组合：\n")
	for i := len(sortedStats) - 5; i < len(sortedStats); i++ {
		file.WriteString(fmt.Sprintf("%s: %d 次\n", sortedStats[i].Hand, sortedStats[i].Count))
	}

	// 输出最低出现次数前5的手牌组合
	file.WriteString("\n最低出现次数前5的手牌组合：\n")
	for i := 0; i < 5; i++ {
		file.WriteString(fmt.Sprintf("%s: %d 次\n", sortedStats[i].Hand, sortedStats[i].Count))
	}
}

func main() {
	// 从标准输入读取玩家输入的手牌
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入你的两张牌（例如：方片7 红桃A）：")
	playerInput, _ := reader.ReadString('\n')
	playerInput = strings.TrimSpace(playerInput)
	playerHand := strings.Fields(playerInput)

	// 创建一副扑克牌
	deck := createDeck()

	// 去掉玩家的两张牌
	var remainingDeck []string
	for _, card := range deck {
		if !contains(playerHand, card) {
			remainingDeck = append(remainingDeck, card)
		}
	}

	// 统计手牌组合频率
	handStats := make(map[string]int)

	// 模拟10万次洗牌和发牌
	numSimulations := 1000000
	for i := 0; i < numSimulations; i++ {
		// 获取对手手牌组合
		hand := getRandomHand(remainingDeck)
		handStats[hand]++
	}

	// 输出统计结果到文件
	printResultsToFile(handStats)

	fmt.Println("统计结果已输出到文件 hand_statistics.txt 中")
}
