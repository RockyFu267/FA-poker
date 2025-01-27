package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 获取微信列表（X0000）
func GetWechatList() {

	url := "http://127.0.0.1:7777/DaenWxHook/httpapi/"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `"type": "X0000",` + "" + `"data": {}` + "" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 微信状态检测（Q0000）
func WechatStatus() {

	url := "http://127.0.0.1:7777/DaenWxHook/httpapi/?wxid=wxid_3sq4tklb6c3121"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `
    "type": "Q0000",` + "" + `"data": {` + "" + `}` + "" + `
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
