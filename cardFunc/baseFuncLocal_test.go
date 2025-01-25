package cardFunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_ReadConfig(t *testing.T) {
	res, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_HandWinRateSimulationDemo01(t *testing.T) {
	handConfig, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	err = HandWinRateSimulationDemo01(handConfig)
	if err != nil {
		log.Println(err)
		return
	}
}

func Test_HandWinRateSimulationDemo02(t *testing.T) {
	handConfig, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	err = HandWinRateSimulationDemo02(handConfig)
	if err != nil {
		log.Println(err)
		return
	}
}
