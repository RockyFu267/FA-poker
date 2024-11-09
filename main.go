package main

import (
	cf "FA-poker/cardFunc"
	"fmt"
)

func main() {
	println("fuck")
	res := cf.ShuffleCard()
	fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Println(cardView)
	}

}
