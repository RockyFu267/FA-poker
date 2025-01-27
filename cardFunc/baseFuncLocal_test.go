package cardFunc

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_ReadConfigDemo02(t *testing.T) {
	res, err := ReadConfigDemo02("/Users/fuao/Desktop/开发/github/FA-poker/Hand02.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	// 将结构体转换为JSON字节切片
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		return
	}
	fmt.Println(string(jsonData))

}

func Test_ReadConfig(t *testing.T) {
	res, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	// 将结构体转换为JSON字节切片
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		return
	}
	fmt.Println(string(jsonData))

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

func Test_HandWinRateSimulationWeb01(t *testing.T) {
	handConfig, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	res, err := HandWinRateSimulationWeb01(handConfig)
	if err != nil {
		log.Println(err)
		return
	}
	// 将结构体转换为JSON字节切片
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func Test_HandWinRateSimulationDemo02(t *testing.T) {
	handConfig, err := ReadConfig("/Users/fuao/Desktop/开发/github/FA-poker/Hand02.yaml")
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
