package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	BaseURL = "http://127.0.0.1:7777/DaenWxHook/httpapi"
	SendTypeURL             = "Q0012"
	SendTypeImage           = "Q0010"
	SendTypeInvite          = "Q0021"
	SendTypeText            = "Q0001"
	SendTypeFriendList      = "Q00005"
	SendTypePass            = "Q0017"
	SendTypeGroupList       = "Q0006"
	SendTypeGroupMemberList = "Q0008"
)

// WechatMessage 微信消息结构体
type WechatMessage struct {
	Event int         `json:"event"`
	Wxid  string      `json:"wxid"`
	Data  MessageData `json:"data"`
}

// MessageData 嵌套的消息数据结构体
type MessageData struct {
	Type      string      `json:"type"`
	Des       string      `json:"des"`
	Data      InnerData   `json:"data"`
	Timestamp string      `json:"timestamp"`
	Wxid      string      `json:"wxid"`
	Port      int         `json:"port"`
	Pid       int         `json:"pid"`
	Flag      string      `json:"flag"`
}

// InnerData 更深层次的嵌套数据结构体
type InnerData struct {
	TimeStamp     string   `json:"timeStamp"`
	FromType      int      `json:"fromType"`
	MsgType       int      `json:"msgType"`
	MsgSource     int      `json:"msgSource"`
	FromWxid      string   `json:"fromWxid"`
	FinalFromWxid string   `json:"finalFromWxid"`
	AtWxidList    []string `json:"atWxidList"`
	Silence       int      `json:"silence"`
	MemberCount   int      `json:"membercount"`
	Signature     string   `json:"signature"`
	Msg           string   `json:"msg"`
	MsgId         string   `json:"msgId"`
}

type Qianxun struct{}

func (b *Qianxun) HandleWechatMsgApi(ctx *gin.Context) {
	// 打印基本请求信息
	fmt.Printf("请求方式: %s\n", ctx.Request.Method)
	fmt.Printf("请求路由: %s\n", ctx.Request.URL.Path)
	fmt.Printf("状态码: %d\n", ctx.Writer.Status())
	fmt.Printf("请求IP: %s\n", ctx.ClientIP())

	// 读取和解析请求体
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "读取请求体失败"})
		return
	}

	// 打印请求体
	fmt.Printf("请求体: %s\n", string(body))

	// 尝试解析JSON
	var msg WechatMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		ctx.JSON(400, gin.H{"error": "JSON解析失败"})
		return
	}

	// 打印解析后的消息
	fmt.Printf("事件: %d\n", msg.Event)
	fmt.Printf("微信ID: %s\n", msg.Wxid)
	fmt.Printf("消息类型: %s\n", msg.Data.Type)
	fmt.Printf("消息描述: %s\n", msg.Data.Des)
	fmt.Printf("消息内容: %+v\n", msg.Data.Data.Msg)

	// 根据事件类型处理不同的业务逻辑
	if msg.Event == 10008 || msg.Event == 10010 {
		fmt.Println("群聊消息")
		// fromWxid := msg.Data.Data.FromWxid
		message := msg.Data.Data.Msg
		msgType := msg.Data.Data.MsgType
		// finalFromWxid := msg.Data.Data.FinalFromWxid

		if msgType == 1 {
			fmt.Println("文本消息")
		} else if msgType == 10000 {
			if strings.Contains(message, "邀请") {
				// name := strings.Split(message, "\"")[3]
				// data := map[string]string{
				// 	"wxid":    fromWxid,
				// 	"title":   fmt.Sprintf("欢迎%s进群", name),
				// 	"content": "和大家打个招呼吧",
				// 	"jumpUrl": "http://127.0.0.1/img/logo.jpg",
				// 	"path":    "http://127.0.0.1/img/logo.jpg",
				// }
				// b.sendMsg(msg.Wxid, SendTypeURL, data)
			}
		}
	} else if msg.Event == 10009 {
		fmt.Println("私聊消息")
		// fromWxid := msg.Data.Data.FromWxid
		// message := msg.Data.Data.Msg
		url := fmt.Sprintf("%s/?wxid=%s", BaseURL, msg.Wxid)
		fmt.Println(url, msg.Data)

		data := map[string]interface{}{
			"type": msg.Data.Type,
			"data": msg.Data,
		}

		resp, err := http.Post(url, "application/json", strings.NewReader(fmt.Sprintf("%v", data)))
		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			ctx.JSON(500, gin.H{"error": "请求失败"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("发送成功")
			ctx.JSON(200, gin.H{"status": "success"})
		} else {
			fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
			ctx.JSON(500, gin.H{"error": "请求失败"})
		}
	} else {
		ctx.JSON(200, gin.H{"status": "success"})
	}
}