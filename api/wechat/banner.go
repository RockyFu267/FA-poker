package wechat

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Banner struct{}

// CreateBannerApi 创建Banner
func (b Banner) CreateBannerApi(ctx *gin.Context) {
	var banner model.BannerEntity
	if err := ctx.ShouldBind(&banner); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	if err := global.AquilaDb.Create(&banner).Error; err != nil {
		utils.Fail(ctx, "Banner创建失败")
		return
	}
	utils.Success(ctx, nil)
}

// GetBannerApi 获取Banner信息
func (b Banner) GetBannerApi(ctx *gin.Context) {
	id := ctx.Query("id")
	var banner model.BannerEntity
	if err := global.AquilaDb.First(&banner, id).Error; err != nil {
		utils.Fail(ctx, "Banner不存在")
		return
	}
	utils.Success(ctx, banner)
}

// GetBannerPageApi 分页获取Banner列表
func (b Banner) GetBannerPageApi(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var banners []model.BannerEntity
	var total int64

	global.AquilaDb.Model(&model.BannerEntity{}).Where("is_hidden = '0'").Count(&total)
	global.AquilaDb.Where("is_hidden = '0'").Order("order_num asc").
		Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&banners)

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  banners,
	})
}

// UpdateBannerApi 更新Banner信息
func (b Banner) UpdateBannerApi(ctx *gin.Context) {
	var banner model.BannerEntity
	if err := ctx.ShouldBind(&banner); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	if err := global.AquilaDb.Save(&banner).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}
	utils.Success(ctx, nil)
}

// DeleteBannerApi 删除Banner(软删除)
func (b Banner) DeleteBannerApi(ctx *gin.Context) {
	id := ctx.Query("id")
	if err := global.AquilaDb.Model(&model.BannerEntity{}).
		Where("id = ?", id).
		Update("del_flag", "1").Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}
	utils.Success(ctx, nil)
}
