package model

import "aquila/global"

const TableNameGroupEntity = "group"

type GroupEntity struct {
    global.GModel
    Name         string  `gorm:"column:name;type:varchar(255);comment:群名称" json:"name"`                  // 群名称
    Announcement string  `gorm:"column:announcement;type:text;comment:群公告" json:"announcement"`           // 群公告
    Tags         string  `gorm:"column:tags;type:text;comment:群标签" json:"tags"`                         // 群标签
    Avatar       string  `gorm:"column:avatar;comment:群头像" json:"avatar"`                               // 群头像
    Longitude    float64 `gorm:"column:longitude;comment:经度" json:"longitude"`                          // 经度
    Latitude     float64 `gorm:"column:latitude;comment:纬度" json:"latitude"`                           // 纬度
    OrderNum     int64   `gorm:"column:order_num;comment:排序号" json:"orderNum"`                         // 排序号
    Qrcode       string  `gorm:"column:qrcode;comment:群二维码" json:"qrcode"`                            // 群二维码
    Keyword      string  `gorm:"column:keyword;type:varchar(255);comment:关键词" json:"keyword"`           // 关键词
    Wxid         string  `gorm:"column:wxid;type:varchar(255);comment:微信群ID" json:"wxid"`              // 微信群ID
    Types        []GroupTypeEntity `gorm:"many2many:group_type_relation;foreignKey:ID;joinForeignKey:GroupID;References:ID;joinReferences:GroupTypeID" json:"types"`                   // 添加多对多关系
}

func (*GroupEntity) TableName() string {
    return TableNameGroupEntity
}
