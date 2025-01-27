package model

import "aquila/global"

const TableNameGroupTypeEntity = "group_type"

type GroupTypeEntity struct {
    global.GModel
    Name       string `gorm:"column:name;type:varchar(255);comment:类型名称" json:"name"`             // 类型名称
    IconActive string `gorm:"column:icon_active;type:varchar(255);comment:激活图标" json:"iconActive"` // 激活状态图标
    IconNormal string `gorm:"column:icon_normal;type:varchar(255);comment:普通图标" json:"iconNormal"` // 普通状态图标
    Color      string `gorm:"column:color;type:varchar(50);comment:颜色" json:"color"`              // 颜色
    OrderNum   int64  `gorm:"column:order_num;default:0;comment:排序号" json:"orderNum"`
}

func (*GroupTypeEntity) TableName() string {
    return TableNameGroupTypeEntity
}
