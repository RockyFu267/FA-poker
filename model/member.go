package model

import "aquila/global"

const TableNameMemberEntity = "member"

type MemberEntity struct {
	global.GModel
	BelongWxid              string `gorm:"column:belong_wxid;type:varchar(100);comment:所属微信号" json:"belongWxid"`
	AvatarMaxURL            string `gorm:"column:avatar_max_url;type:varchar(255);comment:头像大图" json:"avatarMaxUrl"`
	AvatarMinURL            string `gorm:"column:avatar_min_url;type:varchar(255);comment:头像小图" json:"avatarMinUrl"`
	City                    string `gorm:"column:city;type:varchar(50);comment:城市" json:"city"`
	Country                 string `gorm:"column:country;type:varchar(50);comment:国家" json:"country"`
	EnBrief                 string `gorm:"column:en_brief;type:varchar(100);comment:英文简称" json:"enBrief"`
	EnWhole                 string `gorm:"column:en_whole;type:varchar(255);comment:英文全称" json:"enWhole"`
	MomentsBackgroundImgURL string `gorm:"column:moments_background_img_url;type:varchar(255);comment:朋友圈背景图" json:"momentsBackgroudImgUrl"`
	Nick                    string `gorm:"column:nick;type:varchar(100);comment:昵称" json:"nick"`
	NickBrief               string `gorm:"column:nick_brief;type:varchar(100);comment:昵称简拼" json:"nickBrief"`
	NickWhole               string `gorm:"column:nick_whole;type:varchar(255);comment:昵称全拼" json:"nickWhole"`
	Province                string `gorm:"column:province;type:varchar(50);comment:省份" json:"province"`
	Remark                  string `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	RemarkBrief             string `gorm:"column:remark_brief;type:varchar(100);comment:备注简拼" json:"remarkBrief"`
	RemarkWhole             string `gorm:"column:remark_whole;type:varchar(255);comment:备注全拼" json:"remarkWhole"`
	Sex                     string `gorm:"column:sex;type:varchar(10);comment:性别,0=未知,1=男,2=女" json:"sex"`
	Sign                    string `gorm:"column:sign;type:varchar(255);comment:个性签名" json:"sign"`
	V3                      string `gorm:"column:v3;type:varchar(100);comment:V3数据" json:"v3"`
	Wxid                    string `gorm:"column:wxid;type:varchar(100);unique;comment:微信ID" json:"wxid"`
	WxNum                   string `gorm:"column:wx_num;type:varchar(100);comment:微信号" json:"wxNum"`
}

// TableName MemberEntity's table name
func (*MemberEntity) TableName() string {
	return TableNameMemberEntity
}
