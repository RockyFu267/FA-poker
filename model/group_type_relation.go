package model

import "aquila/global"

const TableNameGroupTypeRelation = "group_type_relation"

type GroupTypeRelation struct {
    global.GModel
    GroupID     uint `gorm:"column:group_id;comment:群组ID" json:"groupId"`
    GroupTypeID uint `gorm:"column:group_type_id;comment:群组类型ID" json:"groupTypeId"`
}

func (*GroupTypeRelation) TableName() string {
    return TableNameGroupTypeRelation
}
