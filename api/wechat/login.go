package wechat

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct{}

type WxLoginReq struct {
	Code         string     `json:"code"`
	UserInfo     WxUserInfo `json:"userInfo"`
	From         string     `json:"from"`
	InviteOpenid string     `json:"inviteOpenid"`
}

type WxUserInfo struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
}

type WxLoginResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type UpdateUserReq struct {
	Openid   string     `json:"openid"`
	UserInfo WxUserInfo `json:"userInfo"`
}

func (l Login) WxLoginApi(c *gin.Context) {
	var req WxLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}

	fmt.Printf("Received request: %+v\n", req) // 修复格式化字符串

	wxResp, err := l.requestWxServer(req.Code)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	wxUser, err := l.createOrUpdateUser(wxResp, req)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	utils.Success(c, wxUser)
}

func (l Login) requestWxServer(code string) (*WxLoginResp, error) {
	appID := global.AquilaConfig.Wechat.Appid
	secret := global.AquilaConfig.Wechat.Secret
	wxLoginURL := global.AquilaConfig.Wechat.Url

	resp, err := http.Get(fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		wxLoginURL, appID, secret, code))
	if err != nil {
		return nil, fmt.Errorf("微信服务器请求失败")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败")
	}

	var wxResp WxLoginResp
	if err := json.Unmarshal(body, &wxResp); err != nil {
		return nil, fmt.Errorf("解析响应失败")
	}

	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf(wxResp.ErrMsg)
	}

	return &wxResp, nil
}

func (l Login) createOrUpdateUser(wxResp *WxLoginResp, req WxLoginReq) (*model.WxUserEntity, error) {
	var wxUser model.WxUserEntity
	result := global.AquilaDb.Where("openid = ?", wxResp.OpenID).First(&wxUser)

	if result.Error != nil {
		// 用户不存在，创建新用户
		wxUser = model.WxUserEntity{
			Openid:       wxResp.OpenID,
			SessionKey:   wxResp.SessionKey,
			NickName:     req.UserInfo.NickName,
			AvatarUrl:    req.UserInfo.AvatarUrl,
			Gender:       req.UserInfo.Gender,
			City:         req.UserInfo.City,
			Province:     req.UserInfo.Province,
			Country:      req.UserInfo.Country,
			From:         req.From,
			InviteOpenid: req.InviteOpenid,
			IsWeixin:     true,
			IsAdmin:      false,
		}
		if err := global.AquilaDb.Create(&wxUser).Error; err != nil {
			return nil, fmt.Errorf("创建用户失败: %v", err)
		}
	}

	return &wxUser, nil
}

func (l Login) UpdateUserApi(c *gin.Context) {
	var req UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}

	var wxUser model.WxUserEntity
	result := global.AquilaDb.Where("openid = ?", req.Openid).First(&wxUser)
	if result.Error != nil {
		utils.Fail(c, "用户不存在")
		return
	}
	fmt.Println(req.UserInfo)
	// 更新用户信息
	wxUser.NickName = req.UserInfo.NickName
	wxUser.AvatarUrl = req.UserInfo.AvatarUrl
	wxUser.Gender = req.UserInfo.Gender
	wxUser.City = req.UserInfo.City
	wxUser.Province = req.UserInfo.Province
	wxUser.Country = req.UserInfo.Country

	if err := global.AquilaDb.Save(&wxUser).Error; err != nil {
		utils.Fail(c, fmt.Sprintf("更新用户信息失败: %v", err))
		return
	}

	utils.Success(c, wxUser)
}
