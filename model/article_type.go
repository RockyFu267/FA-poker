package model

import "aquila/global"

const TableNameArticleTypeEntity = "article_type"

type ArticleTypeEntity struct {
    global.GModel
    Name     string `gorm:"column:name;type:varchar(255);comment:类型名称" json:"name"`
    OrderNum int64  `gorm:"column:order_num;default:0;comment:排序号" json:"orderNum"`
}

func (*ArticleTypeEntity) TableName() string {
    return TableNameArticleTypeEntity
}
