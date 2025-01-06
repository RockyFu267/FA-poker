package cardFunc

import (
	"fmt"
	"reflect"
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

func Test_shortOfShuffleCard(t *testing.T) {
	inputTest := []Card{
		{Rank: 4, Suit: "方片"},
		{Rank: 4, Suit: "梅花"},
		{Rank: 4, Suit: "黑桃"},
		{Rank: 8, Suit: "方片"},
		{Rank: 7, Suit: "梅花"},
		{Rank: 4, Suit: "红桃"},
		{Rank: 9, Suit: "方片"},
	}
	res := shortOfShuffleCard(inputTest)
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

	fmt.Println("xxxxxxxxxxxxxxxxxxx")

	inputTest := []Card{
		{Rank: 4, Suit: "方片"},
		{Rank: 4, Suit: "梅花"},
		{Rank: 4, Suit: "黑桃"},
		{Rank: 8, Suit: "方片"},
		{Rank: 7, Suit: "梅花"},
		{Rank: 4, Suit: "红桃"},
		{Rank: 9, Suit: "方片"},
	}
	resA := shortOfShuffleCard(inputTest)
	for _, i := range resA {
		cardView := i.CardTranslate()
		fmt.Print(cardView, " ")
	}

	resA01, resA02 := DealCards(resA, 5)
	fmt.Println(len(resA01))
	for k, v := range resA01 {
		fmt.Println(k, v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate())
	}
	fmt.Println(len(resA), len(resA01), len(resA02))
	fmt.Println("xxxxxxxxxxxxxxxxxxx")
	for _, v := range resA02 {
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
		////带A的同花顺
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
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
		// //4条用例7 长度4 4+1+1+1
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 8, Suit: "梅花"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 7, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// //4条用例8 长度4 1+4+1+1
		// {Rank: 10, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//4条用例9 长度4 1+1+4+1
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// //4条用例10 长度4 1+1+1+4
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "梅花"},
		// {Rank: 9, Suit: "方片"},

		// // 葫芦33 用例1 长度3 3+3+1
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
		// 葫芦33 用例3 长度3 3+3+1
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
		// 葫芦32 用例3 长度4  3+2+1+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// 葫芦32 用例4 长度4  3+1+2+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// 葫芦32 用例5 长度4  3+1+1+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// 葫芦32 用例7 长度4  1+3+1+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// 葫芦32 用例8 长度4  1+1+3+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// 葫芦32 用例9 长度4  2+1+3+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// 葫芦32 用例10 长度4  1+2+3+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// 葫芦32 用例11 长度4  1+1+2+3
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// 葫芦32 用例12 长度4  1+2+1+3
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// 葫芦32 用例13 长度4  2+1+1+3
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// 2对 用例1 长度4  2+2+2+1
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// 2对 用例2 长度4  2+2+1+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// 2对 用例3 长度4  2+1+2+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// 2对 用例4 长度4  1+2+2+2
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// 同花顺 用例1 长度5  A+B+C+D+EEE
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 10, Suit: "方片"},
		// 同花顺 用例2 长度5 A+B+C+DDD+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 10, Suit: "黑桃"},
		//同花顺 用例3 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//同花顺 用例4 长度5 A+BBB+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//同花顺 用例5 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//----------------------同花顺5432A
		// 同花顺 用例6 长度5  A+B+C+D+EEE
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// {Rank: 2, Suit: "红桃"},
		// {Rank: 2, Suit: "方片"},
		// 同花顺 用例7 长度5 A+B+C+DDD+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 3, Suit: "方片"},
		// {Rank: 2, Suit: "黑桃"},
		//同花顺 用例8 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 4, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//同花顺 用例9 长度5 A+BBB+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 5, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//同花顺 用例10 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// 同花顺 用例1 长度6  A+B+C+D+E+FF
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "方片"},
		// 同花顺 用例2 长度6 A+B+C+DD+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 7, Suit: "红桃"},
		//同花顺 用例3 长度6 AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		//----------------------同花顺5432A
		// 同花顺 用例4 长度6  AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "红桃"},
		// 同花顺 用例7 长度6 A+B+CC+D+E+F
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "红桃"},
		//同花顺 用例8 长度6 A+B+CCC+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "黑桃"},
		// {Rank: 2, Suit: "红桃"},
		//同花 用例1 长度5 A+B+C+D+EEE
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 9, Suit: "方片"},
		// 同花 用例2 长度5 A+B+C+DDD+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		//同花 用例3 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		//同花 用例4 长度5 A+BBB+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		//同花 用例5 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		//同花 用例6 长度6 AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 8, Suit: "方片"},
		// 同花 用例7 长度6 A+B+CC+D+E+F
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 8, Suit: "红桃"},
		//同花 用例7 长度6 A+B+C+D+E+FF
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		//顺子 用例1 长度5 A+B+C+D+EEE
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 10, Suit: "方片"},
		// 顺子 用例2 长度5 A+B+C+DDD+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "红桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 10, Suit: "黑桃"},
		//顺子 用例3 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 12, Suit: "方片"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//顺子 用例4 长度5 A+BBB+C+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 13, Suit: "方片"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//顺子 用例5 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		//顺子 用例6 长度6 AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		// 顺子 用例7 长度6 A+B+CC+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		//顺子 用例8 长度6 A+B+C+D+E+FF
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 10, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 8, Suit: "红桃"},
		//----------------------顺5432A
		// A顺子 用例1 长度5  A+B+C+D+EEE
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// {Rank: 2, Suit: "红桃"},
		// {Rank: 2, Suit: "方片"},
		// A顺子 用例2 长度5 A+B+C+DDD+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 3, Suit: "方片"},
		// {Rank: 2, Suit: "黑桃"},
		//A顺子 用例3 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 4, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//A顺子 用例4 长度5 A+BBB+C+D+E
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 5, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//A顺子 用例5 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 5, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// A顺子 用例6 长度6  AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "红桃"},
		// A顺子 用例7 长度6 A+B+C+DDD+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 5, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "红桃"},
		//A顺子 用例8 长度6 A+B+CCC+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 8, Suit: "黑桃"},
		// {Rank: 5, Suit: "红桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 3, Suit: "红桃"},
		// {Rank: 2, Suit: "黑桃"},
		// {Rank: 2, Suit: "红桃"},
		//三条 用例1 长度5 AAA+B+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 14, Suit: "方片"},
		// {Rank: 6, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//三条 用例2 长度5 A+B+CCC+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 6, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 4, Suit: "红桃"},
		// {Rank: 4, Suit: "方片"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		//三条 用例3 长度5 A+B+C+D+EEE
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 6, Suit: "方片"},
		// {Rank: 4, Suit: "黑桃"},
		// {Rank: 3, Suit: "黑桃"},
		// {Rank: 2, Suit: "黑桃"},
		// {Rank: 2, Suit: "红桃"},
		// {Rank: 2, Suit: "方片"},
		//两对 用例1 长度5 AA+BB+C+D+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		//两对 用例2 长度5 AA+B+C+D+EE
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 9, Suit: "红桃"},
		//两对 用例3 长度5 A+BB+C+DD+E
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 11, Suit: "黑桃"},
		// {Rank: 11, Suit: "方片"},
		// {Rank: 9, Suit: "红桃"},
		//对子 用例1 长度6 AA+B+C+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 14, Suit: "红桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 6, Suit: "红桃"},
		//对子 用例2 长度6 A+B+CC+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "黑桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 6, Suit: "红桃"},
		//对子 用例3 长度6 A+B+CC+D+E+F
		// {Rank: 14, Suit: "黑桃"},
		// {Rank: 13, Suit: "红桃"},
		// {Rank: 12, Suit: "红桃"},
		// {Rank: 9, Suit: "黑桃"},
		// {Rank: 8, Suit: "方片"},
		// {Rank: 6, Suit: "黑桃"},
		// {Rank: 6, Suit: "红桃"},
	}

	grade, card5 := Judge5From7(inputTest)
	fmt.Println(grade)
	fmt.Println(card5)
}

// 牌型的测试用例
type testCard struct {
	Card7 [7]Card
	Card5 [5]Card
	grade int
}

func Test_Judge5From7Auto(t *testing.T) {
	var inputTest []testCard
	inputTest = append(inputTest, testCard{
		//-----同花顺---------- 567
		Card7: [7]Card{ //同花顺 长度6 1
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度6 2
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "方片"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 1
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 9, Suit: "黑桃"},
			{Rank: 8, Suit: "方片"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "红桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 2
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "方片"},
		},
		Card5: [5]Card{
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 3
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 12, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 4
			{Rank: 14, Suit: "黑桃"},
			{Rank: 14, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
			{Rank: 14, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 5
			{Rank: 14, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 6, Suit: "红桃"},
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 6, Suit: "红桃"},
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		grade: 8,
	}, testCard{
		Card7: [7]Card{ //同花顺 长度7 6
			{Rank: 14, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "黑桃"},
			{Rank: 2, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 5, Suit: "红桃"},
			{Rank: 4, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
			{Rank: 14, Suit: "红桃"},
		},
		grade: 8,
		//-----四条----------234
		//-----葫芦----------34
		//-----同花----------567
	}, testCard{
		Card7: [7]Card{ //同花 长度7 1
			{Rank: 14, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 6, Suit: "红桃"},
			{Rank: 4, Suit: "黑桃"},
			{Rank: 3, Suit: "黑桃"},
			{Rank: 2, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "红桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 6, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		grade: 5,
	}, testCard{
		Card7: [7]Card{ //同花 长度7 2
			{Rank: 14, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 6, Suit: "红桃"},
			{Rank: 4, Suit: "黑桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 6, Suit: "红桃"},
			{Rank: 3, Suit: "红桃"},
			{Rank: 2, Suit: "红桃"},
		},
		grade: 5,
		//-----顺子----------567
	}, testCard{
		Card7: [7]Card{ //顺子 长度7 1
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 12, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
		},
		grade: 4,
	}, testCard{
		Card7: [7]Card{ //顺子 长度7 2
			{Rank: 14, Suit: "黑桃"},
			{Rank: 12, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 7, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 12, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
		},
		grade: 4,
	}, testCard{
		Card7: [7]Card{ //顺子 长度7 3
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 7, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 7, Suit: "红桃"},
		},
		grade: 4,
		//-----三条----------5
		//-----两对---------- 456
	}, testCard{
		Card7: [7]Card{ //两对 长度5 1
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 14, Suit: "黑桃"},
		},
		grade: 2,
	}, testCard{
		Card7: [7]Card{ //两对 长度5 2
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 9, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 14, Suit: "黑桃"},
		},
		grade: 2,
		//-----对子----------6
	}, testCard{
		Card7: [7]Card{ //对子 长度6 1
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 8, Suit: "黑桃"},
			{Rank: 8, Suit: "红桃"},
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
		},
		grade: 1,
	}, testCard{
		Card7: [7]Card{ //对子 长度6 2
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
		},
		Card5: [5]Card{
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
			{Rank: 14, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
		},
		grade: 1,
		//-----高牌----------7
	}, testCard{
		Card7: [7]Card{ //高牌 长度7 1
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
			{Rank: 8, Suit: "黑桃"},
			{Rank: 6, Suit: "红桃"},
		},
		Card5: [5]Card{
			{Rank: 14, Suit: "黑桃"},
			{Rank: 13, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
			{Rank: 10, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
		},
		grade: 0,
	})
	for k, v := range inputTest {
		grade, card5 := Judge5From7(v.Card7)
		if grade != v.grade || !reflect.DeepEqual(card5, v.Card5) {
			t.Errorf("Test_Judge5From7Auto 第%d个测试用例失败", k)
			fmt.Println(v.Card7)
		}
		fmt.Print(v.Card5[0].CardTranslate(), " ", v.Card5[1].CardTranslate(), " ", v.Card5[2].CardTranslate(), " ", v.Card5[3].CardTranslate(), " ", v.Card5[4].CardTranslate(), "\n")
	}
	fmt.Println("Test_Judge5From7Auto 一共通过", len(inputTest), "个测试用例")
}

func Test_containsStraightKeys(t *testing.T) {
	// 测试包含所有必需键的情况
	cards1 := map[int]int{
		14: 1,
		2:  2,
		3:  3,
		4:  4,
		5:  5,
		6:  6,
		7:  7,
		13: 1,
	}

	// 测试缺少一个必需键的情况
	cards2 := map[int]int{
		14: 1,
		2:  2,
		3:  3,
		4:  4,
	}

	// 测试空地图的情况
	cards3 := map[int]int{}

	res01 := containsStraightKeys(cards1)
	res02 := containsStraightKeys(cards2)
	res03 := containsStraightKeys(cards3)

	fmt.Println(res01)
	fmt.Println(res02)
	fmt.Println(res03)
}

func Test_shuffleJudgeDemo(t *testing.T) {
	var p1, p2, p3, p4, p5, p6 Players
	p1.ID = "1"
	p2.ID = "2"
	p3.ID = "3"
	p4.ID = "4"
	p5.ID = "5"
	p6.ID = "6"
	var players = []Players{p1, p2, p3, p4, p5, p6}
	var appointHandCardList []HandCard
	appointHandCardList = append(appointHandCardList, HandCard{
		HandCard: [2]Card{
			{Rank: 14, Suit: "黑桃"},
			{Rank: 14, Suit: "红桃"},
		},
	}, HandCard{
		HandCard: [2]Card{
			{Rank: 13, Suit: "黑桃"},
			{Rank: 13, Suit: "红桃"},
		},
	}, HandCard{
		HandCard: [2]Card{
			{Rank: 12, Suit: "黑桃"},
			{Rank: 12, Suit: "红桃"},
		},
	}, HandCard{
		HandCard: [2]Card{
			{Rank: 11, Suit: "黑桃"},
			{Rank: 11, Suit: "红桃"},
		},
	}, HandCard{
		HandCard: [2]Card{
			{Rank: 10, Suit: "黑桃"},
			{Rank: 10, Suit: "红桃"},
		},
	}, HandCard{
		HandCard: [2]Card{
			{Rank: 9, Suit: "黑桃"},
			{Rank: 9, Suit: "红桃"},
		},
	})
	winners := shuffleJudgeDemo(players, appointHandCardList)
	for k, v := range winners {
		fmt.Println("--AAA---")
		fmt.Println(k, v)
	}
}

func Test_sortCards(t *testing.T) {
	inputTest := [7]Card{
		{Rank: 4, Suit: "方片"},
		{Rank: 4, Suit: "梅花"},
		{Rank: 4, Suit: "黑桃"},
		{Rank: 8, Suit: "方片"},
		{Rank: 7, Suit: "梅花"},
		{Rank: 4, Suit: "红桃"},
		{Rank: 9, Suit: "方片"},
	}
	res := sortCards(inputTest)
	fmt.Println(res)
}
