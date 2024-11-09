package cardFunc

import (
	"fmt"
	"testing"
)

func Test_ShuffleCard(t *testing.T) {

	res := ShuffleCard()
	fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Println(cardView)
	}
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
