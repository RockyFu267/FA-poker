package main

import (
	fcf "FA-poker/cardFunc"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	// var startsps *string = flag.String("sps", "", "Use -sps <Configuration File Path>")
	// //获取参数
	// flag.Parse()
	// //检查参数合法性
	// if *startsps == "" {
	// 	log.Println("Please use the <-sps> parameter correctly and enter the specified configuration file path.")
	// 	return
	// }
	// path := *startsps
	if runtime.GOOS == "windows" {
		fmt.Println("当前系统是 Windows")
	} else {
		fmt.Println("当前系统不是 Windows")
	}

	fmt.Println("如有bug，请反馈至rocky267@gmail.com")
	fmt.Println("请输入 1 开始范围胜率模拟统计： 记得修改配置文件Hand.yaml")
	fmt.Println("请输入 2 开始单人AOF模拟对局：(还没发版)")
	fmt.Println("请输入 3 开始单人SNG模拟对局：计划中")
	fmt.Println("请输入 4 开始单人CASH模拟对局：计划中")
	fmt.Println("请输入 5 开始单人SIt&Go模拟对局：计划中")
	fmt.Println("请输入 6 开始单人MTT模拟对局：(还没发版)")
	fmt.Println("请输入 7 开始单人级别挑战：还没发版 通关者联系rocky领取奖励")
	fmt.Println("请输入 8 开始多人自定义对局：计划中")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch input.Text() {
		case "1":
			path := "./Hand.yaml"
			configTmp, err := fcf.ReadConfig(path)
			if err != nil {
				log.Println(err)
				return
			}
			err = fcf.HandWinRateSimulationDemo01(configTmp)
			if err != nil {
				log.Println(err)
				return
			}
			if runtime.GOOS == "windows" {
				time.Sleep(10000 * time.Second)
			}
			return
		case "2":
			fmt.Println("预计春节前出，催更请联系rocky267@gmail.com")
			if runtime.GOOS == "windows" {
				time.Sleep(10 * time.Second)
			}
			return
		case "3":
			fmt.Println("预计春节后，催更请联系rocky267@gmail.com")
			if runtime.GOOS == "windows" {
				time.Sleep(10 * time.Second)
			}
			return
		case "4":
			fmt.Println("预计春节后，催更请联系rocky267@gmail.com")
			if runtime.GOOS == "windows" {
				time.Sleep(10 * time.Second)
			}
			return
		case "7":
			fmt.Println("预计春节后，催更请联系rocky267@gmail.com")
			if runtime.GOOS == "windows" {
				time.Sleep(10 * time.Second)
			}
			return
		default:
			fmt.Println("有空就写,现在没空，催更请联系rocky267@gmail.com")
			if runtime.GOOS == "windows" {
				time.Sleep(10 * time.Second)
			}
			return
		}

	}
}
