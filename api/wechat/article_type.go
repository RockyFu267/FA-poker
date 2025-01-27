package wechat

import (
    "aquila/global"
    "aquila/model"
    "aquila/utils"
    "github.com/gin-gonic/gin"
)

type ArticleType struct{}

// CreateArticleTypeApi 创建文章类型
func (at ArticleType) CreateArticleTypeApi(ctx *gin.Context) {
    var articleType model.ArticleTypeEntity
    if err := ctx.ShouldBind(&articleType); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Create(&articleType).Error; err != nil {
        utils.Fail(ctx, "文章类型创建失败")
        return
    }
    utils.Success(ctx, nil)
}

// GetArticleTypeApi 获取文章类型信息
func (at ArticleType) GetArticleTypeApi(ctx *gin.Context) {
    id := ctx.Query("id")
    var articleType model.ArticleTypeEntity
    if err := global.AquilaDb.First(&articleType, id).Error; err != nil {
        utils.Fail(ctx, "文章类型不存在")
        return
    }
    utils.Success(ctx, articleType)
}

// GetArticleTypeListApi 获取文章类型列表
func (at ArticleType) GetArticleTypeListApi(ctx *gin.Context) {
    var articleTypes []model.ArticleTypeEntity
    if err := global.AquilaDb.Order("order_num asc").Find(&articleTypes).Error; err != nil {
        utils.Fail(ctx, "查询失败")
        return
    }
    utils.Success(ctx, articleTypes)
}

// UpdateArticleTypeApi 更新文章类型信息
func (at ArticleType) UpdateArticleTypeApi(ctx *gin.Context) {
    var articleType model.ArticleTypeEntity
    if err := ctx.ShouldBind(&articleType); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Save(&articleType).Error; err != nil {
        utils.Fail(ctx, "更新失败")
        return
    }
    utils.Success(ctx, nil)
}

// DeleteArticleTypeApi 删除文章类型
func (at ArticleType) DeleteArticleTypeApi(ctx *gin.Context) {
    id := ctx.Query("id")
    if err := global.AquilaDb.Delete(&model.ArticleTypeEntity{}, id).Error; err != nil {
        utils.Fail(ctx, "删除失败")
        return
    }
    utils.Success(ctx, nil)
}
