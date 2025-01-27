package system

import (
	"aquila/global"
	"aquila/utils"

	"github.com/gin-gonic/gin"
)

type InitDb struct{}

type InitDto struct {
	Force bool `json:"force"` // 是否强制初始化
}

const (
	userSQL = `INSERT INTO user (id, created_at, updated_at, deleted_at, username, password, user_type, mobile, sort, status, last_login_ip, last_login_nation, last_login_province, last_login_city, last_login_date, salt, email, avatar, nickname) VALUES (1, '2024-09-04 21:22:07', '2024-09-12 13:58:03', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 0, '', 1, 0, '', '', '', '', '2024-09-04 21:22:07', '', '', '', '111');`

	roleSQL = `INSERT INTO role (id, created_at, updated_at, deleted_at, name, remark, status) 
		VALUES (1, '2024-09-10 23:20:44', '2024-09-10 23:20:44', NULL, 'admin', '', 0),
               (2, '2024-09-10 23:21:45', '2024-09-10 23:21:45', NULL, 'test1', '', 0),
               (3, '2024-09-10 23:21:48', '2024-09-10 23:21:48', NULL, 'test2', '', 0),
               (4, '2024-09-11 10:52:50', '2024-09-11 10:52:50', NULL, 'test3', '', 0),
               (5, '2024-09-12 14:37:09', '2024-09-12 14:38:58', NULL, 'role1', '', 0);`

	menuSQL = `INSERT INTO menu (id, created_at, updated_at, deleted_at, name, parent_id, order_num, path, component, query, is_frame, menu_type, is_catch, is_hidden, perms, icon, status, remark) 
		VALUES (1, '2024-09-11 10:03:48', '2024-09-11 10:03:48', NULL, '系统管理', 0, 0, '/system', 'Layout', '', 1, 'C', 0, 0, '', 'system', 0, ''),
               (2, '2024-09-11 10:12:11', '2024-09-11 10:12:11', NULL, '用户管理', 1, 1, 'user', 'system/user/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (3, '2024-09-11 10:23:24', '2024-09-11 10:23:24', NULL, '菜单管理', 1, 3, 'menu', 'system/menu/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (4, '2024-09-11 10:23:41', '2024-09-11 10:23:41', NULL, '角色管理', 1, 2, 'role', 'system/role/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (5, '2024-09-12 16:27:44', '2024-09-12 16:27:44', NULL, '1', 0, 0, '', '', '', 0, '', 0, 0, '', '', 0, '');`

	roleMenuSQL = `INSERT INTO role_menu (id, created_at, updated_at, deleted_at, role_id, menu_id) 
		VALUES (22, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 1),
               (23, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 2),
               (24, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 3),
               (25, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 4);`

	userRoleSQL = `INSERT INTO user_role (id, created_at, updated_at, deleted_at, user_id, role_id) 
		VALUES (5, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 1),
               (6, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 2),
               (7, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 3),
               (8, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 4);`
)

func (i InitDb) InitializeDBApi(ctx *gin.Context) {
	var req InitDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	// 如果不是强制初始化,检查是否已经初始化过
	if !req.Force {
		var count int64
		global.AquilaDb.Table("user").Count(&count)
		if count > 0 {
			utils.Fail(ctx, "数据库已初始化")
			return
		}
	}

	// 开启事务
	tx := global.AquilaDb.Begin()

	// 清空已有数据
	if req.Force {
		tables := []string{"user_role", "role_menu", "menu", "role", "user"}
		for _, table := range tables {
			if err := tx.Exec("TRUNCATE TABLE " + table).Error; err != nil {
				tx.Rollback()
				utils.Fail(ctx, "清空表数据失败:"+err.Error())
				return
			}
		}
	}

	// 按顺序执行SQL语句
	sqlStatements := []string{
		userSQL,     // 先创建用户
		roleSQL,     // 再创建角色
		menuSQL,     // 再创建菜单
		roleMenuSQL, // 再创建角色-菜单关系
		userRoleSQL, // 最后创建用户-角色关系
	}

	for _, sql := range sqlStatements {
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			utils.Fail(ctx, "执行SQL失败:"+err.Error())
			return
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		utils.Fail(ctx, "提交事务失败:"+err.Error())
		return
	}

	utils.Success(ctx, nil)
}
