package wechat

import (
    "aquila/global"
    "aquila/model"
    "aquila/utils"
    "github.com/gin-gonic/gin"
    "strconv"
)

type GroupType struct{}

// CreateGroupTypeApi 创建群组类型
func (gt GroupType) CreateGroupTypeApi(ctx *gin.Context) {
    var groupType model.GroupTypeEntity
    if err := ctx.ShouldBind(&groupType); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Create(&groupType).Error; err != nil {
        utils.Fail(ctx, "群组类型创建失败")
        return
    }
    utils.Success(ctx, nil)
}

// GetGroupTypeApi 获取群组类型信息
func (gt GroupType) GetGroupTypeApi(ctx *gin.Context) {
    id := ctx.Query("id")
    var groupType model.GroupTypeEntity
    if err := global.AquilaDb.First(&groupType, id).Error; err != nil {
        utils.Fail(ctx, "群组类型不存在")
        return
    }
    utils.Success(ctx, groupType)
}

// GetGroupTypePageApi 分页获取群组类型列表
func (gt GroupType) GetGroupTypePageApi(ctx *gin.Context) {
    pageNum := ctx.DefaultQuery("pageNum", "1")
    pageSize := ctx.DefaultQuery("pageSize", "10")
    pageNumInt, _ := strconv.Atoi(pageNum)
    pageSizeInt, _ := strconv.Atoi(pageSize)

    var groupTypes []model.GroupTypeEntity
    var total int64
    
    global.AquilaDb.Model(&model.GroupTypeEntity{}).Count(&total)
    global.AquilaDb.Order("order_num asc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&groupTypes)

    utils.Success(ctx, gin.H{
        "total": total,
        "data":  groupTypes,
    })
}

// UpdateGroupTypeApi 更新群组类型信息
func (gt GroupType) UpdateGroupTypeApi(ctx *gin.Context) {
    var groupType model.GroupTypeEntity
    if err := ctx.ShouldBind(&groupType); err != nil {
        utils.Fail(ctx, "参数绑定失败"+err.Error())
        return
    }

    if err := global.AquilaDb.Save(&groupType).Error; err != nil {
        utils.Fail(ctx, "更新失败")
        return
    }
    utils.Success(ctx, nil)
}

// DeleteGroupTypeApi 删除群组类型
func (gt GroupType) DeleteGroupTypeApi(ctx *gin.Context) {
    id := ctx.Query("id")
    if err := global.AquilaDb.Delete(&model.GroupTypeEntity{}, id).Error; err != nil {
        utils.Fail(ctx, "删除失败")
        return
    }
    utils.Success(ctx, nil)
}
