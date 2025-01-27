package model

import (
	"aquila/global"
	"time"
)

const TableNameWxUser = "wx_user"

type WxUserEntity struct {
	global.GModel
	NickName     string    `gorm:"column:nick_name;type:varchar(50)" json:"nickName"`
	AvatarUrl    string    `gorm:"column:avatar_url;type:varchar(250);default:''" json:"avatarUrl"`
	City         string    `gorm:"column:city;type:varchar(50);default:''" json:"city"`
	Country      string    `gorm:"column:country;type:varchar(50);default:''" json:"country"`
	Province     string    `gorm:"column:province;type:varchar(50);default:''" json:"province"`
	Openid       string    `gorm:"column:openid;type:varchar(100);uniqueIndex" json:"openid"`
	SessionKey   string    `gorm:"column:session_key;type:varchar(100)" json:"sessionKey"`
	Language     string    `gorm:"column:language;type:varchar(50);default:''" json:"language"`
	From         string    `gorm:"column:from;type:varchar(50);default:''" json:"from"`
	IsDemote     bool      `gorm:"column:is_demote;default:false" json:"isDemote"`
	Gender       int       `gorm:"column:gender;default:0" json:"gender"`
	IsWeixin     bool      `gorm:"column:is_weixin;default:true" json:"isWeixin"`
	IsAdmin      bool      `gorm:"column:is_admin;default:false" json:"isAdmin"`
	Email        string    `gorm:"column:email;type:varchar(50);default:''" json:"email"`
	Password     string    `gorm:"column:password;type:varchar(100);default:''" json:"password"`
	Phonenumber  string    `gorm:"column:phonenumber;type:varchar(11);default:''" json:"phonenumber"`
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"`
	DelFlag      string    `gorm:"column:del_flag;type:char(1);default:'0'" json:"delFlag"`
	Remark       string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	InviteOpenid string    `gorm:"column:invite_openid;type:varchar(100)" json:"inviteOpenid"`
}

func (*WxUserEntity) TableName() string {
	return TableNameWxUser
}
