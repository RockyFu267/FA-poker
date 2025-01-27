package model

import "aquila/global"

const TableNameBannerEntity = "banner"

type BannerEntity struct {
    global.GModel
    Title        string `gorm:"column:title;type:varchar(255);comment:标题" json:"title"`                      // 标题
    ImageUrl     string `gorm:"column:image_url;type:varchar(255);comment:图片地址" json:"imageUrl"`            // 图片地址
    RedirectType string `gorm:"column:redirect_type;type:char(1);default:'';comment:跳转类型" json:"redirectType"` // 跳转类型: 0群 1文章 2外链 3小程序
    RedirectInfo string `gorm:"column:redirect_info;type:varchar(255);default:'';comment:跳转信息" json:"redirectInfo"` // 跳转信息
    OrderNum     int64  `gorm:"column:order_num;default:0;comment:排序号" json:"orderNum"`                     // 排序号
		IsHidden     int64  `gorm:"column:is_hidden;default:0;comment:是否隐藏" json:"isHidden"`                  // 是否隐藏
}

func (*BannerEntity) TableName() string {
    return TableNameBannerEntity
}
