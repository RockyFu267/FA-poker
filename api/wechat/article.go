package wechat

import (
    "aquila/global"
    "aquila/model"
    "aquila/utils"
    "github.com/gin-gonic/gin"
    "strconv"
)

type Article struct{}

// CreateArticleApi 创建文章
func (a Article) CreateArticleApi(ctx *gin.Context) {
    var article model.ArticleEntity
    if err := ctx.ShouldBind(&article); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Create(&article).Error; err != nil {
        utils.Fail(ctx, "文章创建失败")
        return
    }
    utils.Success(ctx, nil)
}

// GetArticleApi 获取文章信息
func (a Article) GetArticleApi(ctx *gin.Context) {
    id := ctx.Query("id")
    var article model.ArticleEntity
    if err := global.AquilaDb.Preload("ArticleType").First(&article, id).Error; err != nil {
        utils.Fail(ctx, "文章不存在")
        return
    }
    utils.Success(ctx, article)
}

// GetArticlePageApi 分页获取文章列表
func (a Article) GetArticlePageApi(ctx *gin.Context) {
    pageNum := ctx.DefaultQuery("pageNum", "1")
    pageSize := ctx.DefaultQuery("pageSize", "10")
    typeID := ctx.Query("typeId")
    pageNumInt, _ := strconv.Atoi(pageNum)
    pageSizeInt, _ := strconv.Atoi(pageSize)

    var articles []model.ArticleEntity
    var total int64

    query := global.AquilaDb.Model(&model.ArticleEntity{}).Preload("ArticleType")

    if typeID != "" {
        query = query.Where("article_type_id = ?", typeID)
    }

    query.Count(&total)
    query.Order("created_at desc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&articles)

    utils.Success(ctx, gin.H{
        "total": total,
        "data":  articles,
    })
}

// UpdateArticleApi 更新文章信息
func (a Article) UpdateArticleApi(ctx *gin.Context) {
    var article model.ArticleEntity
    if err := ctx.ShouldBind(&article); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Save(&article).Error; err != nil {
        utils.Fail(ctx, "更新失败")
        return
    }
    utils.Success(ctx, nil)
}

// DeleteArticleApi 删除文章
func (a Article) DeleteArticleApi(ctx *gin.Context) {
    id := ctx.Query("id")
    if err := global.AquilaDb.Delete(&model.ArticleEntity{}, id).Error; err != nil {
        utils.Fail(ctx, "删除失败")
        return
    }
    utils.Success(ctx, nil)
}
