package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Menu struct{}

func (m *Menu) CreateMenuApi(ctx *gin.Context) {
	var req MenuDto
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	fmt.Printf("Received request: %+v\n", req) // 打印接收到的请求内容
	var menu model.MenuEntity
	err = global.AquilaDb.Where("name = ?", req.Name).First(&menu).Error
	if err != nil {
		// 创建新菜单
		menu = model.MenuEntity{
			Name:      req.Name,
			ParentId:  req.ParentId,
			OrderNum:  req.OrderNum,
			Path:      req.Path,
			Component: req.Component,
			Query:     req.Query,
			IsFrame:   req.IsFrame,
			MenuType:  req.MenuType,
			IsCatch:   req.IsCatch,
			IsHidden:  req.IsHidden,
			Perms:     req.Perms,
			Icon:      req.Icon,
			Status:    req.Status,
			Remark:    req.Remark,
		}
		err = global.AquilaDb.Create(&menu).Error
		if err != nil {
			utils.Fail(ctx, "菜单创建失败")
			return
		}
		utils.Success(ctx, "菜单创建成功")
		return
	}
	utils.Fail(ctx, "菜单已存在")
}

func (m *Menu) GetMenuApi(ctx *gin.Context) {
	var req MenuDto
	// get请求参数绑定
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	fmt.Printf("Received request: %+v\n", req) // 打印接收到的请求内容
	var menu model.MenuEntity
	err = global.AquilaDb.Where("name = ?", req.Name).First(&menu).Error
	if err != nil {
		utils.Fail(ctx, "菜单不存在")
		return
	}
	utils.Success(ctx, menu)
}

func (m *Menu) GetMenuAllApi(ctx *gin.Context) {
	var req MenuPageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var menus []model.MenuEntity
	// 获取所有的menu
	err := global.AquilaDb.Find(&menus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	// menus 根据id和parentId，组装成树形结构，children: []MenuEntity
	var menuTree []UserMenuTreeDto
	menuTree = getMenuTree(0, menus)

	req.Data = menuTree
	req.Total = int64(len(menuTree))
	utils.Success(ctx, req)
}

// 删除菜单
func (m *Menu) DeleteMenuApi(ctx *gin.Context) {
	var req MenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, err.Error())
		return
	}

	if err := global.AquilaDb.Where("id = ?", req.Id).Delete(&model.MenuEntity{}).Error; err != nil {
		utils.Fail(ctx, "删除菜单失败")
		return
	}
	utils.Success(ctx, "删除成功")
}

// 修改菜单
func (m *Menu) UpdateMenuApi(ctx *gin.Context) {
	var req MenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, err.Error())
		return
	}

	updates := map[string]interface{}{
		"name":      req.Name,
		"parent_id": req.ParentId,
		"order_num": req.OrderNum,
		"path":      req.Path,
		"component": req.Component,
		"query":     req.Query,
		"is_frame":  req.IsFrame,
		"menu_type": req.MenuType,
		"is_catch":  req.IsCatch,
		"is_hidden": req.IsHidden,
		"perms":     req.Perms,
		"icon":      req.Icon,
		"status":    req.Status,
		"remark":    req.Remark,
	}

	if err := global.AquilaDb.Model(&model.MenuEntity{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		utils.Fail(ctx, "更新菜单失败")
		return
	}
	utils.Success(ctx, "更新成功")
}

// 分页查询菜单
func (m *Menu) GetMenuPageApi(ctx *gin.Context) {
	var req MenuPageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, err.Error())
		return
	}

	var menus []model.MenuEntity
	var total int64

	db := global.AquilaDb.Model(&model.MenuEntity{})

	// // 条件查询
	// if req.Name != "" {
	// 	db = db.Where("name LIKE ?", "%"+req.Name+"%")
	// }
	// if req.Status != nil {
	// 	db = db.Where("status = ?", req.Status)
	// }

	// 获取总数
	db.Count(&total)

	// 分页查询
	offset := (req.PageNum - 1) * req.PageSize
	err := db.Offset(int(offset)).Limit(int(req.PageSize)).Find(&menus).Error
	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	req.Total = total
	req.Data = menus
	utils.Success(ctx, req)
}
