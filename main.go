package main

import (
	fcf "FA-poker/cardFunc"
	"flag"
	"log"
)

func main() {
	var startsps *string = flag.String("sps", "", "Use -sps <Configuration File Path>")
	//获取参数
	flag.Parse()
	//检查参数合法性
	if *startsps == "" {
		log.Println("Please use the <-sps> parameter correctly and enter the specified configuration file path.")
		return
	}
	path := *startsps
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
}
