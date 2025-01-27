package wechat

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Group struct{}

// CreateGroupApi 创建群组
func (g Group) CreateGroupApi(ctx *gin.Context) {
	var group model.GroupEntity
	if err := ctx.ShouldBind(&group); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	if err := global.AquilaDb.Create(&group).Error; err != nil {
		utils.Fail(ctx, "群组创建失败")
		return
	}
	utils.Success(ctx, nil)
}

// GetGroupApi 获取群组信息
func (g Group) GetGroupApi(ctx *gin.Context) {
	id := ctx.Query("id")
	var group model.GroupEntity
	if err := global.AquilaDb.Preload("Types").First(&group, id).Error; err != nil {
		utils.Fail(ctx, "群组不存在")
		return
	}
	utils.Success(ctx, group)
}

// GetGroupPageApi 分页获取群组列表
func (g Group) GetGroupPageApi(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	typeID := ctx.Query("typeId")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var groups []model.GroupEntity
	var total int64

	query := global.AquilaDb.Model(&model.GroupEntity{})

	if typeID != "" {
		subQuery := global.AquilaDb.Table("group_type_relation").
			Select("group_id").
			Where("group_type_id = ?", typeID)
		
		query = query.Where("id IN (?)", subQuery)
	}

	// 先获取总数
	query.Count(&total)

	// 然后查询数据并预加载类型
	err := query.Preload("Types").
		Order("order_num asc").
		Scopes(utils.Paginate(pageNumInt, pageSizeInt)).
		Find(&groups).Error

	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  groups,
	})
}

// UpdateGroupApi 更新群组信息
func (g Group) UpdateGroupApi(ctx *gin.Context) {
	var group model.GroupEntity
	if err := ctx.ShouldBind(&group); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	if err := global.AquilaDb.Save(&group).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}
	utils.Success(ctx, nil)
}

// DeleteGroupApi 删除群组
func (g Group) DeleteGroupApi(ctx *gin.Context) {
	id := ctx.Query("id")
	if err := global.AquilaDb.Delete(&model.GroupEntity{}, id).Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}
	utils.Success(ctx, nil)
}

// SetGroupTypesApi 设置群组类型
func (g Group) SetGroupTypesApi(ctx *gin.Context) {
	type Param struct {
		GroupID      uint   `json:"groupId" binding:"required"`
		GroupTypeIDs []uint `json:"groupTypeIds" binding:"required"`
	}

	var param Param
	if err := ctx.ShouldBind(&param); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	// 开始事务
	tx := global.AquilaDb.Begin()

	// 删除原有关联
	if err := tx.Where("group_id = ?", param.GroupID).Delete(&model.GroupTypeRelation{}).Error; err != nil {
		tx.Rollback()
		utils.Fail(ctx, "设置失败")
		return
	}

	// 创建新关联
	for _, typeID := range param.GroupTypeIDs {
		relation := model.GroupTypeRelation{
			GroupID:     param.GroupID,
			GroupTypeID: typeID,
		}
		if err := tx.Create(&relation).Error; err != nil {
			tx.Rollback()
			utils.Fail(ctx, "设置失败")
			return
		}
	}

	tx.Commit()
	utils.Success(ctx, nil)
}
