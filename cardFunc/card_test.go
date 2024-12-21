package cardFunc

import (
	"fmt"
	"testing"
)

func Test_CardTranslate(t *testing.T) {
	var p, p1, p2, p3 Card
	p.Rank = 14
	p.Suit = "黑桃"

	p1.Rank = 13
	p1.Suit = "红桃"

	p2.Rank = 12
	p2.Suit = "梅花"

	p3.Rank = 11
	p3.Suit = "方片"

	fmt.Println(p, p1, p2, p3)
}

func Test_ShuffleCard(t *testing.T) {

	res := ShuffleCard()
	fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Println(cardView)
	}
}

func Test_compareSuits(t *testing.T) {
	a := "黑桃"
	b := "红桃"

	fmt.Println(compareSuits(a, b))
}

func Test_GetTopTwoCards(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := ShuffleCard()
		// fmt.Println(a)
		b := GetTopTwoCards(a)
		fmt.Println(b.HandCard[0].CardTranslate(), b.HandCard[1].CardTranslate())
	}

}

func Test_ShuffleAndRecord(t *testing.T) {
	ShuffleAndRecord(100000, "res.txt")

}

func Test_DealCards(t *testing.T) {

	res := ShuffleCard()
	// fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Print(cardView, " ")
	}

	res01, res02 := DealCards(res, 10)
	fmt.Println(len(res01))
	for k, v := range res01 {
		fmt.Println(k, v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate())
	}
	fmt.Println("--------")
	for _, v := range res02 {
		fmt.Print(v.CardTranslate(), " ")
	}
}

func Test_CombineCardsDemo(t *testing.T) {
	// 1. 生成52张牌
	res := ShuffleCard()
	// fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Print(cardView, " ")
	}

	res01, res02 := DealCards(res, 10)
	combinations := CombineCardsDemo(res01, res02)
	fmt.Println("\n所有组合的数量：", len(combinations))
	for k, v := range combinations {
		fmt.Print("\n组合", k, ":")
		for _, card := range v {
			fmt.Print(card.CardTranslate(), " ")
		}
	}

}

func Test_Judge5From7(t *testing.T) {
	inputTest := [7]Card{
		// //带A的同花顺
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// //4条用例
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 8, Suit: "梅花"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// //4条用例
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// //4条用例1 长度3 4+ 2+1或1+2
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 2, Suit: "黑桃"},
		//4条用例2 长度3 1+4+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		//4条用例3 长度3 2+4+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 7, Suit: "黑桃"},
		//4条用例4 长度3 2+1或者1+2 +4
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 7, Suit: "黑桃"},

		// // //4条用例5 4+3  长度 2
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 7, Suit: "黑桃"},
		// {Rank: 7, Suit: "红桃"},
		// {Rank: 7, Suit: "梅花"},
		// // //4条用例6 4+3  长度 2
		// {Rank: 7, Suit: "黑桃"},
		// {Rank: 7, Suit: "红桃"},
		// {Rank: 7, Suit: "梅花"},
		// {Rank: 6, Suit: "黑桃"},
		// {Rank: 6, Suit: "红桃"},
		// {Rank: 6, Suit: "梅花"},
		// {Rank: 6, Suit: "方片"},
		// // 葫芦33 用例1 长度 3 3+3+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "梅花"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 2, Suit: "黑桃"},
		// // 葫芦33 用例2 长度3   1+3+3
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "梅花"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 12, Suit: "梅花"},
		// 葫芦33 用例3 3+3+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "梅花"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 12, Suit: "梅花"},
		// // 葫芦32 用例1 长度3  3+2+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "梅花"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 12, Suit: "梅花"},
		// 葫芦32 用例2 长度3  2+3+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 12, Suit: "梅花"},
		// 葫芦32 用例2 长度3  2+2+3
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "梅花"},
	}

	grade, card5 := Judge5From7(inputTest)
	fmt.Println(grade)
	fmt.Println(card5)
}
