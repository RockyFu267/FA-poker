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
	ShuffleAndRecord(1000000, "res.txt")

}

func Test_DealCards(t *testing.T) {

	res := ShuffleCard()
	fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Println(cardView)
	}

	res01, res02 := DealCards(res, 2)
	fmt.Print(res01)
	fmt.Print("--------")
	fmt.Print(res02)
}
