package model

import "aquila/global"

const TableNameArticleEntity = "article"

type ArticleEntity struct {
	global.GModel
	Title         string            `gorm:"column:title;type:varchar(255);comment:标题" json:"title"`
	SubTitle      string            `gorm:"column:sub_title;type:varchar(255);comment:副标题" json:"subTitle"`
	Summary       string            `gorm:"column:summary;type:varchar(255);comment:摘要" json:"summary"`
	Author        string            `gorm:"column:author;type:varchar(50);comment:作者" json:"author"`
	Content       string            `gorm:"column:content;type:text;comment:内容" json:"content"`
	Cover         string            `gorm:"column:cover;type:varchar(255);comment:封面图" json:"cover"`
	Status        int64             `gorm:"column:status;default:0;comment:状态" json:"status"` // 0: 有效 1: 过期
	ArticleTypeID uint              `gorm:"column:article_type_id;comment:文章类型ID" json:"articleTypeId"`
	ArticleType   ArticleTypeEntity `gorm:"foreignKey:ArticleTypeID" json:"articleType"`
}

func (*ArticleEntity) TableName() string {
	return TableNameArticleEntity
}
